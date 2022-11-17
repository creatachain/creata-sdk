package keeper

import (
	"encoding/binary"

	"github.com/creatachain/creata-sdk/codec"

	"github.com/creatachain/creata-sdk/x/upgrade/types"

	msm "github.com/creatachain/augusteum/msm/types"

	sdk "github.com/creatachain/creata-sdk/types"
	sdkerrors "github.com/creatachain/creata-sdk/types/errors"
)

// NewQuerier creates a querier for upgrade cli and REST endpoints
func NewQuerier(k Keeper, legacyQuerierCdc *codec.LegacyAmino) sdk.Querier {
	return func(ctx sdk.Context, path []string, req msm.RequestQuery) ([]byte, error) {
		switch path[0] {

		case types.QueryCurrent:
			return queryCurrent(ctx, req, k, legacyQuerierCdc)

		case types.QueryApplied:
			return queryApplied(ctx, req, k, legacyQuerierCdc)

		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unknown %s query endpoint: %s", types.ModuleName, path[0])
		}
	}
}

func queryCurrent(ctx sdk.Context, _ msm.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	plan, has := k.GetUpgradePlan(ctx)
	if !has {
		return nil, nil
	}

	res, err := legacyQuerierCdc.MarshalJSON(&plan)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func queryApplied(ctx sdk.Context, req msm.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryAppliedPlanRequest

	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	applied := k.GetDoneHeight(ctx, params.Name)
	if applied == 0 {
		return nil, nil
	}

	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, uint64(applied))

	return bz, nil
}
