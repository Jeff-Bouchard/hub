package querier

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/sentinel-official/hub/x/dvpn/subscription/keeper"
	"github.com/sentinel-official/hub/x/dvpn/subscription/types"
)

func queryPlan(ctx sdk.Context, req abci.RequestQuery, k keeper.Keeper) ([]byte, sdk.Error) {
	var params types.QueryPlanParams
	if err := types.ModuleCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, types.ErrorUnmarshal()
	}

	plan, found := k.GetPlan(ctx, params.Address, params.ID)
	if !found {
		return nil, nil
	}

	res, err := types.ModuleCdc.MarshalJSON(plan)
	if err != nil {
		return nil, types.ErrorMarshal()
	}

	return res, nil
}

func queryPlans(ctx sdk.Context, _ abci.RequestQuery, k keeper.Keeper) ([]byte, sdk.Error) {
	res, err := types.ModuleCdc.MarshalJSON(k.GetPlans(ctx))
	if err != nil {
		return nil, types.ErrorMarshal()
	}

	return res, nil
}

func queryPlansOfProvider(ctx sdk.Context, req abci.RequestQuery, k keeper.Keeper) ([]byte, sdk.Error) {
	var params types.QueryPlansOfProviderParams
	if err := types.ModuleCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, types.ErrorUnmarshal()
	}

	res, err := types.ModuleCdc.MarshalJSON(k.GetPlansOfProvider(ctx, params.Address))
	if err != nil {
		return nil, types.ErrorMarshal()
	}

	return res, nil
}
