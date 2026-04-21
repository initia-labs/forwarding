package simapp

import (
	"cosmossdk.io/core/store"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/ibc-go/v10/modules/apps/transfer"
	transferkeeper "github.com/cosmos/ibc-go/v10/modules/apps/transfer/keeper"
	transfertypes "github.com/cosmos/ibc-go/v10/modules/apps/transfer/types"
	ibc "github.com/cosmos/ibc-go/v10/modules/core"
	clienttypes "github.com/cosmos/ibc-go/v10/modules/core/02-client/types"
	connectiontypes "github.com/cosmos/ibc-go/v10/modules/core/03-connection/types"
	porttypes "github.com/cosmos/ibc-go/v10/modules/core/05-port/types"
	ibcexported "github.com/cosmos/ibc-go/v10/modules/core/exported"
	ibckeeper "github.com/cosmos/ibc-go/v10/modules/core/keeper"
	tendermint "github.com/cosmos/ibc-go/v10/modules/light-clients/07-tendermint"
	"github.com/noble-assets/forwarding/v2"
)

func (app *SimApp) RegisterLegacyModules() error {
	if err := app.RegisterStores(
		storetypes.NewKVStoreKey(ibcexported.StoreKey),
		storetypes.NewKVStoreKey(transfertypes.StoreKey),
	); err != nil {
		return err
	}

	app.ParamsKeeper.Subspace(ibcexported.ModuleName).WithKeyTable(clienttypes.ParamKeyTable().RegisterParamSet(&connectiontypes.Params{}))
	app.ParamsKeeper.Subspace(transfertypes.ModuleName).WithKeyTable(transfertypes.ParamKeyTable())

	ibcStoreService := storeService(app, ibcexported.StoreKey)
	app.IBCKeeper = ibckeeper.NewKeeper(
		app.appCodec,
		ibcStoreService,
		app.GetSubspace(ibcexported.ModuleName),
		app.UpgradeKeeper,
		"noble1h8tqx833l3t2s45mwxjz29r85dcevy93wk63za",
	)

	transferStoreService := storeService(app, transfertypes.StoreKey)
	app.TransferKeeper = transferkeeper.NewKeeper(
		app.appCodec,
		transferStoreService,
		app.GetSubspace(transfertypes.ModuleName),
		app.IBCKeeper.ChannelKeeper,
		app.IBCKeeper.ChannelKeeper,
		app.MsgServiceRouter(),
		app.AccountKeeper,
		app.BankKeeper,
		"noble1h8tqx833l3t2s45mwxjz29r85dcevy93wk63za",
	)

	var transferStack porttypes.IBCModule
	transferStack = transfer.NewIBCModule(app.TransferKeeper)
	transferStack = forwarding.NewMiddleware(transferStack, app.AccountKeeper, app.ForwardingKeeper)

	ibcRouter := porttypes.NewRouter().AddRoute(transfertypes.ModuleName, transferStack)
	app.IBCKeeper.SetRouter(ibcRouter)

	app.ForwardingKeeper.SetIBCKeepers(app.IBCKeeper.ChannelKeeper, app.TransferKeeper)

	tmLightClientModule := tendermint.NewLightClientModule(app.appCodec, app.IBCKeeper.ClientKeeper.GetStoreProvider())

	return app.RegisterModules(
		ibc.NewAppModule(app.IBCKeeper),
		transfer.NewAppModule(app.TransferKeeper),
		tendermint.NewAppModule(tmLightClientModule),
	)
}

func storeService(app *SimApp, key string) store.KVStoreService {
	return runtime.NewKVStoreService(app.GetKey(key))
}
