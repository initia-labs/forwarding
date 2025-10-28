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
	"context"
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/noble-assets/forwarding/v2/simapp"
	"github.com/noble-assets/forwarding/v2/types"
)

func TestQueryGetMemoNotFound(t *testing.T) {
	app, sdkCtx := setupForwardingKeeper(t)
	ensureOpenChannel(t, app, sdkCtx, "channel-0")

	addr := registerAccountWithMemos(t, app, sdkCtx, "channel-0", "iaa1recipient", nil)

	_, err := app.ForwardingKeeper.GetMemo(sdkCtx, &types.QueryMemo{
		Address: addr,
		Denom:   "uusdc",
	})
	require.Error(t, err)
	require.ErrorContains(t, err, "failed to get memo from state")
}

func TestQueryGetMemosFiltersByAddress(t *testing.T) {
	app, sdkCtx := setupForwardingKeeper(t)
	ensureOpenChannel(t, app, sdkCtx, "channel-0")

	account1Memos := []types.MemoEntry{
		{Denom: "uusdc", Memo: "memo-usdc"},
		{Denom: "uatom", Memo: "memo-atom"},
	}
	account2Memos := []types.MemoEntry{
		{Denom: "uosmo", Memo: "memo-osmo"},
	}

	addr1 := registerAccountWithMemos(t, app, sdkCtx, "channel-0", "iaa1recipientone", account1Memos)
	registerAccountWithMemos(t, app, sdkCtx, "channel-0", "iaa1recipienttwo", account2Memos)

	resp, err := app.ForwardingKeeper.GetMemos(sdkCtx, &types.QueryMemos{
		Address: addr1,
	})
	require.NoError(t, err)
	require.Len(t, resp.Memos, len(account1Memos))

	expected := make(map[string]string, len(account1Memos))
	for _, entry := range account1Memos {
		expected[entry.Denom] = entry.Memo
	}

	for _, memo := range resp.Memos {
		want, ok := expected[memo.Denom]
		require.True(t, ok, fmt.Sprintf("unexpected denom %s", memo.Denom))
		require.Equal(t, want, memo.Memo)
	}
}

func registerAccountWithMemos(t *testing.T, appCtx *simapp.SimApp, ctx context.Context, channel, recipient string, memos []types.MemoEntry) string {
	t.Helper()

	signer := sdk.AccAddress(secp256k1.GenPrivKey().PubKey().Address()).String()
	msg := &types.MsgRegisterAccount{
		Signer:    signer,
		Recipient: recipient,
		Channel:   channel,
		Memos:     memos,
	}

	res, err := appCtx.ForwardingKeeper.RegisterAccount(ctx, msg)
	require.NoError(t, err)

	return res.Address
}
