package querier

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/plan/keeper"
	"github.com/sentinel-official/hub/x/plan/types"
)

func queryPlan(ctx sdk.Context, req abci.RequestQuery, k keeper.Keeper) ([]byte, sdk.Error) {
	var params types.QueryPlanParams
	if err := types.ModuleCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, types.ErrorUnmarshal()
	}

	plan, found := k.GetPlan(ctx, params.ID)
	if !found {
		return nil, nil
	}

	res, err := types.ModuleCdc.MarshalJSON(plan)
	if err != nil {
		return nil, types.ErrorMarshal()
	}

	return res, nil
}

func queryPlans(ctx sdk.Context, req abci.RequestQuery, k keeper.Keeper) ([]byte, sdk.Error) {
	var params types.QueryPlansParams
	if err := types.ModuleCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, types.ErrorUnmarshal()
	}

	var plans types.Plans
	if params.Status.Equal(hub.StatusActive) {
		plans = k.GetActivePlans(ctx, params.Skip, params.Limit)
	} else if params.Status.Equal(hub.StatusInactive) {
		plans = k.GetInactivePlans(ctx, params.Skip, params.Limit)
	} else {
		plans = k.GetPlans(ctx, params.Skip, params.Limit)
	}

	res, err := types.ModuleCdc.MarshalJSON(plans)
	if err != nil {
		return nil, types.ErrorMarshal()
	}

	return res, nil
}

func queryPlansForProvider(ctx sdk.Context, req abci.RequestQuery, k keeper.Keeper) ([]byte, sdk.Error) {
	var params types.QueryPlansForProviderParams
	if err := types.ModuleCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, types.ErrorUnmarshal()
	}

	var plans types.Plans
	if params.Status.Equal(hub.StatusActive) {
		plans = k.GetActivePlansForProvider(ctx, params.Address, params.Skip, params.Limit)
	} else if params.Status.Equal(hub.StatusInactive) {
		plans = k.GetInactivePlansForProvider(ctx, params.Address, params.Skip, params.Limit)
	} else {
		plans = k.GetPlansForProvider(ctx, params.Address, params.Skip, params.Limit)
	}

	res, err := types.ModuleCdc.MarshalJSON(plans)
	if err != nil {
		return nil, types.ErrorMarshal()
	}

	return res, nil
}

func queryNodesForPlan(ctx sdk.Context, req abci.RequestQuery, k keeper.Keeper) ([]byte, sdk.Error) {
	var params types.QueryNodesForPlanParams
	if err := types.ModuleCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, types.ErrorUnmarshal()
	}

	nodes := k.GetNodesForPlan(ctx, params.ID, params.Skip, params.Limit)

	res, err := types.ModuleCdc.MarshalJSON(nodes)
	if err != nil {
		return nil, types.ErrorMarshal()
	}

	return res, nil
}
