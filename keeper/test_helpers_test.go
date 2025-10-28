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
	"sync"
	"testing"
	"time"

	"cosmossdk.io/log"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	dbm "github.com/cosmos/cosmos-db"
	sdk "github.com/cosmos/cosmos-sdk/types"
	transfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
	channeltypes "github.com/cosmos/ibc-go/v8/modules/core/04-channel/types"
	"github.com/stretchr/testify/require"

	"github.com/noble-assets/forwarding/v2/simapp"
)

const (
	bech32PrefixAccAddr  = "noble"
	bech32PrefixAccPub   = bech32PrefixAccAddr + "pub"
	bech32PrefixValAddr  = bech32PrefixAccAddr + "valoper"
	bech32PrefixValPub   = bech32PrefixAccAddr + "valoperpub"
	bech32PrefixConsAddr = bech32PrefixAccAddr + "valcons"
	bech32PrefixConsPub  = bech32PrefixAccAddr + "valconspub"
)

var configureSDKOnce sync.Once

func configureSDK() {
	configureSDKOnce.Do(func() {
		cfg := sdk.GetConfig()
		cfg.SetBech32PrefixForAccount(bech32PrefixAccAddr, bech32PrefixAccPub)
		cfg.SetBech32PrefixForValidator(bech32PrefixValAddr, bech32PrefixValPub)
		cfg.SetBech32PrefixForConsensusNode(bech32PrefixConsAddr, bech32PrefixConsPub)
		cfg.Seal()
	})
}

func setupForwardingKeeper(t *testing.T) (*simapp.SimApp, sdk.Context) {
	t.Helper()

	configureSDK()

	app, err := simapp.NewSimApp(log.NewNopLogger(), dbm.NewMemDB(), nil, true, emptyAppOptions{})
	require.NoError(t, err)

	sdkCtx := app.BaseApp.NewUncachedContext(false, tmproto.Header{
		Height: 1,
		Time:   time.Now().UTC(),
	})

	return app, sdkCtx
}

func ensureOpenChannel(t *testing.T, app *simapp.SimApp, sdkCtx sdk.Context, channelID string) {
	t.Helper()

	channel := channeltypes.NewChannel(
		channeltypes.OPEN,
		channeltypes.UNORDERED,
		channeltypes.NewCounterparty(transfertypes.PortID, "counter-0"),
		[]string{"connection-0"},
		transfertypes.Version,
	)
	app.IBCKeeper.ChannelKeeper.SetChannel(sdkCtx, transfertypes.PortID, channelID, channel)
}

type emptyAppOptions struct{}

func (emptyAppOptions) Get(string) interface{} { return nil }
