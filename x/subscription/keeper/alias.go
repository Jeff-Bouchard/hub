package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
	node "github.com/sentinel-official/hub/x/node/types"
	plan "github.com/sentinel-official/hub/x/plan/types"
)

func (k Keeper) SendCoin(ctx sdk.Context, from sdk.AccAddress, to sdk.AccAddress, coin sdk.Coin) sdk.Error {
	return k.bank.SendCoins(ctx, from, to, sdk.NewCoins(coin))
}

func (k Keeper) AddDeposit(ctx sdk.Context, address sdk.AccAddress, coin sdk.Coin) sdk.Error {
	return k.deposit.Add(ctx, address, sdk.NewCoins(coin))
}

func (k Keeper) SubtractDeposit(ctx sdk.Context, address sdk.AccAddress, coin sdk.Coin) sdk.Error {
	if coin.IsZero() {
		return nil
	}

	return k.deposit.Subtract(ctx, address, sdk.NewCoins(coin))
}

func (k Keeper) GetNode(ctx sdk.Context, address hub.NodeAddress) (node.Node, bool) {
	return k.node.GetNode(ctx, address)
}

func (k Keeper) GetPlan(ctx sdk.Context, id uint64) (plan.Plan, bool) {
	return k.plan.GetPlan(ctx, id)
}
