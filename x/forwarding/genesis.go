package forwarding

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/noble-assets/forwarding/x/forwarding/keeper"
	"github.com/noble-assets/forwarding/x/forwarding/types"
)

func InitGenesis(ctx context.Context, k *keeper.Keeper, genesis types.GenesisState) {
	for channel, count := range genesis.NumOfAccounts {
		_ = k.NumOfAccounts.Set(ctx, channel, count)
	}

	for channel, count := range genesis.NumOfForwards {
		_ = k.NumOfForwards.Set(ctx, channel, count)
	}

	for channel, total := range genesis.TotalForwarded {
		_ = k.TotalForwarded.Set(ctx, channel, total)
	}
}

func ExportGenesis(ctx sdk.Context, k *keeper.Keeper) *types.GenesisState {
	return &types.GenesisState{
		NumOfAccounts:  k.GetAllNumOfAccounts(ctx),
		NumOfForwards:  k.GetAllNumOfForwards(ctx),
		TotalForwarded: k.GetAllTotalForwarded(ctx),
	}
}
