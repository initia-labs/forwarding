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

package keeper_test

import (
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/stretchr/testify/require"

	"github.com/initia-labs/forwarding/v2/keeper"
	"github.com/initia-labs/forwarding/v2/types"
)

func TestValidateAccountFields(t *testing.T) {
	key := secp256k1.GenPrivKey()
	addr := sdk.AccAddress(key.PubKey().Address())

	tests := []struct {
		name        string
		malleate    func(acc sdk.AccountI) error
		address     []byte
		errContains string
	}{
		{
			name:     "New account",
			address:  addr,
			malleate: func(acc sdk.AccountI) error { return nil },
		},
		{
			name:    "Account with nil pub key but non zero sequence",
			address: addr,
			malleate: func(acc sdk.AccountI) error {
				return acc.SetSequence(1)
			},
			errContains: "attempting to register an existing user",
		},
		{
			name:    "Account created signerlessly",
			address: addr,
			malleate: func(acc sdk.AccountI) error {
				if err := acc.SetPubKey(&types.ForwardingPubKey{Key: addr}); err != nil {
					return err
				}
				return acc.SetSequence(1)
			},
		},
		{
			name:    "Account created signerlessly but wrong address",
			address: sdk.AccAddress(secp256k1.GenPrivKey().PubKey().Address()),
			malleate: func(acc sdk.AccountI) error {
				if err := acc.SetPubKey(&types.ForwardingPubKey{Key: addr}); err != nil {
					return err
				}
				return acc.SetSequence(1)
			},
			errContains: "attempting to register an existing user",
		},
		{
			name:    "Account created with different pub key type",
			address: addr,
			malleate: func(acc sdk.AccountI) error {
				key := secp256k1.GenPrivKey()
				return acc.SetPubKey(key.PubKey())
			},
			errContains: "attempting to register an existing user",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			baseAcc := &authtypes.BaseAccount{Address: addr.String()}
			err := test.malleate(baseAcc)
			require.NoError(t, err, "expected no error configuring the account")

			err = keeper.ValidateAccountFields(baseAcc, test.address)
			if test.errContains != "" {
				require.Error(t, err)
				require.ErrorContains(t, err, test.errContains)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestRegisterAccountStoresInitialMemos(t *testing.T) {
	app, sdkCtx := setupForwardingKeeper(t)
	ensureOpenChannel(t, app, sdkCtx, "channel-0")

	memos := []types.MemoEntry{
		{Denom: "uusdc", Memo: "cross-chain-usdc"},
		{Denom: "uatom", Memo: "forward-atom"},
	}

	msg := &types.MsgRegisterAccount{
		Signer:    sdk.AccAddress(secp256k1.GenPrivKey().PubKey().Address()).String(),
		Recipient: "iaa1recipient",
		Channel:   "channel-0",
		Memos:     memos,
	}

	res, err := app.ForwardingKeeper.RegisterAccount(sdkCtx, msg)
	require.NoError(t, err)
	require.NotEmpty(t, res.Address)

	for _, entry := range memos {
		resp, err := app.ForwardingKeeper.GetMemo(sdkCtx, &types.QueryMemo{
			Address: res.Address,
			Denom:   entry.Denom,
		})
		require.NoError(t, err)
		require.Equal(t, entry.Memo, resp.Memo)
	}
}

func TestRegisterAccountRejectsTooManyMemos(t *testing.T) {
	app, sdkCtx := setupForwardingKeeper(t)
	ensureOpenChannel(t, app, sdkCtx, "channel-0")

	memos := make([]types.MemoEntry, keeper.MaxMemoEntries+1)
	for i := range memos {
		memos[i] = types.MemoEntry{
			Denom: fmt.Sprintf("denom-%d", i),
			Memo:  fmt.Sprintf("memo-%d", i),
		}
	}

	msg := &types.MsgRegisterAccount{
		Signer:    sdk.AccAddress(secp256k1.GenPrivKey().PubKey().Address()).String(),
		Recipient: "iaa1recipient",
		Channel:   "channel-0",
		Memos:     memos,
	}

	_, err := app.ForwardingKeeper.RegisterAccount(sdkCtx, msg)
	require.Error(t, err)
	require.ErrorContains(t, err, "cannot register more than")
}

func TestSetMemoRequiresOwner(t *testing.T) {
	app, sdkCtx := setupForwardingKeeper(t)
	ensureOpenChannel(t, app, sdkCtx, "channel-0")

	recipient := "iaa1recipient"
	fallback := sdk.AccAddress(secp256k1.GenPrivKey().PubKey().Address()).String()
	register := &types.MsgRegisterAccount{
		Signer:    sdk.AccAddress(secp256k1.GenPrivKey().PubKey().Address()).String(),
		Recipient: recipient,
		Fallback:  fallback,
		Channel:   "channel-0",
	}

	_, err := app.ForwardingKeeper.RegisterAccount(sdkCtx, register)
	require.NoError(t, err)

	_, err = app.ForwardingKeeper.SetMemo(sdkCtx, &types.MsgSetMemo{
		Signer:    sdk.AccAddress(secp256k1.GenPrivKey().PubKey().Address()).String(),
		Recipient: recipient,
		Fallback:  fallback,
		Channel:   "channel-0",
		Denom:     "uusdc",
		Memo:      "malicious",
	})
	require.Error(t, err)
	require.ErrorContains(t, err, "only the forwarding account's receiver or fallback account can modify memos")

	_, err = app.ForwardingKeeper.SetMemo(sdkCtx, &types.MsgSetMemo{
		Signer:    recipient,
		Recipient: recipient,
		Fallback:  fallback,
		Channel:   "channel-0",
		Denom:     "uusdc",
		Memo:      "owner-memo",
	})
	require.NoError(t, err)

	_, err = app.ForwardingKeeper.SetMemo(sdkCtx, &types.MsgSetMemo{
		Signer:    fallback,
		Recipient: recipient,
		Fallback:  fallback,
		Channel:   "channel-0",
		Denom:     "uusdc",
		Memo:      "owner-memo",
	})
	require.NoError(t, err)
}

func TestSetMemoAllowsAdditionalEntries(t *testing.T) {
	app, sdkCtx := setupForwardingKeeper(t)
	ensureOpenChannel(t, app, sdkCtx, "channel-0")

	memos := make([]types.MemoEntry, keeper.MaxMemoEntries)
	for i := range memos {
		memos[i] = types.MemoEntry{
			Denom: fmt.Sprintf("denom-%d", i),
			Memo:  fmt.Sprintf("memo-%d", i),
		}
	}

	fallback := sdk.AccAddress(secp256k1.GenPrivKey().PubKey().Address()).String()
	register := &types.MsgRegisterAccount{
		Signer:    sdk.AccAddress(secp256k1.GenPrivKey().PubKey().Address()).String(),
		Recipient: "iaa1recipient",
		Channel:   "channel-0",
		Fallback:  fallback,
		Memos:     memos,
	}

	res, err := app.ForwardingKeeper.RegisterAccount(sdkCtx, register)
	require.NoError(t, err)

	_, err = app.ForwardingKeeper.SetMemo(sdkCtx, &types.MsgSetMemo{
		Signer:    fallback,
		Recipient: register.Recipient,
		Channel:   register.Channel,
		Fallback:  fallback,
		Denom:     "extra",
		Memo:      "new-memo",
	})
	require.NoError(t, err)

	_, err = app.ForwardingKeeper.SetMemo(sdkCtx, &types.MsgSetMemo{
		Signer:    fallback,
		Recipient: register.Recipient,
		Channel:   register.Channel,
		Fallback:  fallback,
		Denom:     memos[0].Denom,
		Memo:      "updated",
	})
	require.NoError(t, err)

	_, err = app.ForwardingKeeper.SetMemo(sdkCtx, &types.MsgSetMemo{
		Signer:    fallback,
		Recipient: register.Recipient,
		Channel:   register.Channel,
		Fallback:  fallback,
		Denom:     memos[0].Denom,
		Memo:      "",
	})
	require.NoError(t, err)

	resp, err := app.ForwardingKeeper.GetMemo(sdkCtx, &types.QueryMemo{
		Address: res.Address,
		Denom:   "extra",
	})
	require.NoError(t, err)
	require.Equal(t, "new-memo", resp.Memo)

	_, err = app.ForwardingKeeper.SetMemo(sdkCtx, &types.MsgSetMemo{
		Signer:    fallback,
		Recipient: register.Recipient,
		Channel:   register.Channel,
		Fallback:  fallback,
		Denom:     "unused",
		Memo:      "",
	})
	require.NoError(t, err)
}
