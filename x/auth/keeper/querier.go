package keeper

import (
	msm "github.com/creatachain/augusteum/msm/types"

	"github.com/creatachain/creata-sdk/codec"
	sdk "github.com/creatachain/creata-sdk/types"
	sdkerrors "github.com/creatachain/creata-sdk/types/errors"
	"github.com/creatachain/creata-sdk/x/auth/types"
)

// NewQuerier creates a querier for auth REST endpoints
func NewQuerier(k AccountKeeper, legacyQuerierCdc *codec.LegacyAmino) sdk.Querier {
	return func(ctx sdk.Context, path []string, req msm.RequestQuery) ([]byte, error) {
		switch path[0] {
		case types.QueryAccount:
			return queryAccount(ctx, req, k, legacyQuerierCdc)

		case types.QueryParams:
			return queryParams(ctx, k, legacyQuerierCdc)

		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unknown query path: %s", path[0])
		}
	}
}

func queryAccount(ctx sdk.Context, req msm.RequestQuery, k AccountKeeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryAccountRequest
	if err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	addr, err := sdk.AccAddressFromBech32(params.Address)
	if err != nil {
		return nil, err
	}

	account := k.GetAccount(ctx, addr)
	if account == nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownAddress, "account %s does not exist", params.Address)
	}

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, account)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}

func queryParams(ctx sdk.Context, k AccountKeeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	params := k.GetParams(ctx)

	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}