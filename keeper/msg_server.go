// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2025, NASD Inc. All rights reserved.
// Use of this software is governed by the Business Source License included
// in the LICENSE file of this repository and at www.mariadb.com/bsl11.
//
// ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
// TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
// VERSIONS OF THE LICENSED WORK.
//
// THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
// LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
// LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
//
// TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
// AN "AS IS" BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
// EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
// TITLE.

package keeper

import (
	"context"
	"errors"
	"fmt"

	"cosmossdk.io/collections"
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	transfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
	channeltypes "github.com/cosmos/ibc-go/v8/modules/core/04-channel/types"
	"github.com/noble-assets/forwarding/v2/types"
)

var _ types.MsgServer = &Keeper{}

func (k *Keeper) RegisterAccount(ctx context.Context, msg *types.MsgRegisterAccount) (*types.MsgRegisterAccountResponse, error) {
	if !channeltypes.IsValidChannelID(msg.Channel) {
		return nil, errors.New("invalid channel")
	}

	if len(msg.Recipient) > transfertypes.MaximumReceiverLength {
		return nil, fmt.Errorf("recipient address must not exceed %d bytes", transfertypes.MaximumReceiverLength)
	}

	if msg.Fallback != "" {
		if _, err := k.accountKeeper.AddressCodec().StringToBytes(msg.Fallback); err != nil {
			return nil, errors.New("invalid fallback address")
		}
	}
	address := types.GenerateAddress(msg.Channel, msg.Recipient, msg.Fallback)

	channel, found := k.channelKeeper.GetChannel(sdk.UnwrapSDKContext(ctx), transfertypes.PortID, msg.Channel)
	if !found {
		return nil, fmt.Errorf("channel does not exist: %s", msg.Channel)
	}
	if channel.State != channeltypes.OPEN {
		return nil, fmt.Errorf("channel is not open: %s, %s", msg.Channel, channel.State)
	}

	if k.accountKeeper.HasAccount(ctx, address) {
		rawAccount := k.accountKeeper.GetAccount(ctx, address)

		if err := ValidateAccountFields(rawAccount, address); err != nil {
			return nil, err
		}

		switch account := rawAccount.(type) {
		case *authtypes.BaseAccount:
			rawAccount = &types.ForwardingAccount{
				BaseAccount: account,
				Channel:     msg.Channel,
				Recipient:   msg.Recipient,
				CreatedAt:   k.headerService.GetHeaderInfo(ctx).Height,
				Fallback:    msg.Fallback,
			}
			k.accountKeeper.SetAccount(ctx, rawAccount)

			k.IncrementNumOfAccounts(ctx, msg.Channel)
		case *types.ForwardingAccount:
			return nil, errors.New("account has already been registered")
		default:
			return nil, fmt.Errorf("unsupported account type: %T", rawAccount)
		}

		for _, denom := range k.GetAllowedDenoms(ctx) {
			balance := k.bankKeeper.GetBalance(ctx, address, denom)
			if !balance.IsZero() {
				account, ok := rawAccount.(*types.ForwardingAccount)
				if ok {
					k.SetPendingForward(ctx, account)
					break
				}
			}
		}

		if err := k.setInitialMemos(ctx, address, msg.Memos); err != nil {
			return nil, err
		}

		return &types.MsgRegisterAccountResponse{Address: address.String()}, k.eventService.EventManager(ctx).Emit(ctx, &types.AccountRegistered{
			Address:   address.String(),
			Channel:   msg.Channel,
			Recipient: msg.Recipient,
			Fallback:  msg.Fallback,
		})
	}

	base := k.accountKeeper.NewAccountWithAddress(ctx, address)
	account := types.ForwardingAccount{
		BaseAccount: authtypes.NewBaseAccount(base.GetAddress(), base.GetPubKey(), base.GetAccountNumber(), base.GetSequence()),
		Channel:     msg.Channel,
		Recipient:   msg.Recipient,
		CreatedAt:   k.headerService.GetHeaderInfo(ctx).Height,
		Fallback:    msg.Fallback,
	}

	k.accountKeeper.SetAccount(ctx, &account)
	k.IncrementNumOfAccounts(ctx, msg.Channel)

	if err := k.setInitialMemos(ctx, address, msg.Memos); err != nil {
		return nil, err
	}

	return &types.MsgRegisterAccountResponse{Address: address.String()}, k.eventService.EventManager(ctx).Emit(ctx, &types.AccountRegistered{
		Address:   address.String(),
		Channel:   account.Channel,
		Recipient: account.Recipient,
		Fallback:  account.Fallback,
	})
}

func (k *Keeper) ClearAccount(ctx context.Context, msg *types.MsgClearAccount) (*types.MsgClearAccountResponse, error) {
	address, err := k.accountKeeper.AddressCodec().StringToBytes(msg.Address)
	if err != nil {
		return nil, errors.New("invalid account address")
	}

	rawAccount := k.accountKeeper.GetAccount(ctx, address)
	if rawAccount == nil {
		return nil, errors.New("account does not exist")
	}
	account, ok := rawAccount.(*types.ForwardingAccount)
	if !ok {
		return nil, errors.New("account is not a forwarding account")
	}

	totalBalance := sdk.NewCoins()
	for _, denom := range k.GetAllowedDenoms(ctx) {
		balance := k.bankKeeper.GetBalance(ctx, address, denom)
		totalBalance = totalBalance.Add(balance)
	}
	if totalBalance.IsZero() {
		return nil, errors.New("account does not require clearing")
	}

	if !msg.Fallback || account.Fallback == "" {
		k.SetPendingForward(ctx, account)
		return &types.MsgClearAccountResponse{}, nil
	}

	fallback, err := k.accountKeeper.AddressCodec().StringToBytes(account.Fallback)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "failed to decode fallback address")
	}
	err = k.bankKeeper.SendCoins(ctx, address, fallback, totalBalance)
	if err != nil {
		return nil, errors.New("failed to clear balance to fallback account")
	}

	return &types.MsgClearAccountResponse{}, k.eventService.EventManager(ctx).Emit(ctx, &types.AccountCleared{
		Address:   msg.Address,
		Recipient: account.Fallback,
	})
}

func (k *Keeper) SetAllowedDenoms(ctx context.Context, msg *types.MsgSetAllowedDenoms) (*types.MsgSetAllowedDenomsResponse, error) {
	if msg.Signer != k.authority {
		return nil, sdkerrors.Wrapf(types.ErrInvalidAuthority, "expected %s, got %s", k.authority, msg.Signer)
	}

	if err := types.ValidateAllowedDenoms(msg.Denoms); err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidDenoms, err.Error())
	}

	previousDenoms := k.GetAllowedDenoms(ctx)
	if err := k.AllowedDenoms.Clear(ctx, nil); err != nil {
		return nil, errors.New("failed to clear allowed denoms from state")
	}
	for _, denom := range msg.Denoms {
		err := k.AllowedDenoms.Set(ctx, denom)
		if err != nil {
			return nil, fmt.Errorf("failed to set %s as allowed denom in state", denom)
		}
	}

	return &types.MsgSetAllowedDenomsResponse{}, k.eventService.EventManager(ctx).Emit(ctx, &types.AllowedDenomsConfigured{
		PreviousDenoms: previousDenoms,
		CurrentDenoms:  msg.Denoms,
	})
}

// ValidateAccountFields is a utility for checking if an account is eligible to be registered.
//
// A valid account must satisfy one of the following conditions.
//
// 1. Is a new account:
//   - A nil PubKey, i.e. the account never sent a transaction AND
//   - A 0 sequence.
//
// 2. Is an account registered signerlessy:
//   - A non nil PubKey with the custom type, i.e. the account has been registered sending a
//     signerless tx.
//   - Can have any sequence value.
func ValidateAccountFields(account sdk.AccountI, address sdk.AccAddress) error {
	pubKey := account.GetPubKey()

	isNewAccount := pubKey == nil && account.GetSequence() == 0
	isValidPubKey := pubKey != nil && pubKey.Equals(&types.ForwardingPubKey{Key: address})

	if !isNewAccount && !isValidPubKey {
		return fmt.Errorf("attempting to register an existing user account with address: %s", address.String())
	}
	return nil
}

// MaxMemoLength is the maximum length for a memo string.
const MaxMemoLength = 1024
const MaxMemoEntries = 10

func (k *Keeper) SetMemo(ctx context.Context, msg *types.MsgSetMemo) (*types.MsgSetMemoResponse, error) {
	if err := sdk.ValidateDenom(msg.Denom); err != nil {
		return nil, fmt.Errorf("invalid denom: %w", err)
	}
	address := types.GenerateAddress(msg.Channel, msg.Recipient, msg.Fallback)
	rawAccount := k.accountKeeper.GetAccount(ctx, address)
	if rawAccount == nil {
		return nil, errors.New("account does not exist")
	}
	_, ok := rawAccount.(*types.ForwardingAccount)
	if !ok {
		return nil, errors.New("account is not a forwarding account")
	}

	addr := address.String()
	if msg.Signer != msg.Recipient && msg.Signer != msg.Fallback && msg.Signer != k.authority {
		return nil, errors.New("only the forwarding account's receiver or fallback account can modify memos")
	}

	pair := collections.Join(addr, msg.Denom)
	if l := len(msg.Memo); l == 0 {
		err := k.Memos.Remove(ctx, pair)
		if err != nil && !errors.Is(err, collections.ErrNotFound) {
			return nil, fmt.Errorf("failed to delete memo from state: %w", err)
		}
	} else if l > MaxMemoLength {
		return nil, fmt.Errorf("memo exceeds maximum length of %d characters", MaxMemoLength)
	} else if err := k.Memos.Set(ctx, pair, msg.Memo); err != nil {
		return nil, fmt.Errorf("failed to set memo in state: %w", err)
	}

	return &types.MsgSetMemoResponse{}, k.eventService.EventManager(ctx).Emit(ctx, &types.MemoSet{
		Address: addr,
		Denom:   msg.Denom,
		Memo:    msg.Memo,
	})
}

func (k *Keeper) setInitialMemos(ctx context.Context, address sdk.AccAddress, entries []types.MemoEntry) error {
	if len(entries) == 0 {
		return nil
	}

	if err := validateMemoEntries(entries); err != nil {
		return err
	}

	addr := address.String()
	for _, entry := range entries {
		if err := k.Memos.Set(ctx, collections.Join(addr, entry.Denom), entry.Memo); err != nil {
			return fmt.Errorf("failed to set memo in state: %w", err)
		}
	}
	return nil
}

func validateMemoEntries(entries []types.MemoEntry) error {
	if len(entries) > MaxMemoEntries {
		return fmt.Errorf("cannot register more than %d memos", MaxMemoEntries)
	}

	seen := make(map[string]struct{}, len(entries))
	for _, entry := range entries {
		if entry.Denom == "" {
			return errors.New("memo denom cannot be empty")
		}
		if _, ok := seen[entry.Denom]; ok {
			return fmt.Errorf("duplicate memo denom: %s", entry.Denom)
		}
		if len(entry.Memo) == 0 {
			return fmt.Errorf("memo for denom %s cannot be empty", entry.Denom)
		}
		if len(entry.Memo) > MaxMemoLength {
			return fmt.Errorf("memo for denom %s exceeds maximum length of %d characters", entry.Denom, MaxMemoLength)
		}
		if err := sdk.ValidateDenom(entry.Denom); err != nil {
			return fmt.Errorf("invalid denom %s: %w", entry.Denom, err)
		}
		seen[entry.Denom] = struct{}{}
	}
	return nil
}
