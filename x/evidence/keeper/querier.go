package keeper

import (
	"github.com/creatachain/creata-sdk/client"
	"github.com/creatachain/creata-sdk/codec"
	sdk "github.com/creatachain/creata-sdk/types"
	sdkerrors "github.com/creatachain/creata-sdk/types/errors"
	"github.com/creatachain/creata-sdk/x/evidence/exported"
	"github.com/creatachain/creata-sdk/x/evidence/types"

	msm "github.com/creatachain/augusteum/msm/types"
)

func NewQuerier(k Keeper, legacyQuerierCdc *codec.LegacyAmino) sdk.Querier {
	return func(ctx sdk.Context, path []string, req msm.RequestQuery) ([]byte, error) {
		var (
			res []byte
			err error
		)

		switch path[0] {
		case types.QueryEvidence:
			res, err = queryEvidence(ctx, req, k, legacyQuerierCdc)

		case types.QueryAllEvidence:
			res, err = queryAllEvidence(ctx, req, k, legacyQuerierCdc)

		default:
			err = sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unknown %s query endpoint: %s", types.ModuleName, path[0])
		}

		return res, err
	}
}

func queryEvidence(ctx sdk.Context, req msm.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryEvidenceRequest

	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	evidence, ok := k.GetEvidence(ctx, params.EvidenceHash)
	if !ok {
		return nil, sdkerrors.Wrap(types.ErrNoEvidenceExists, params.EvidenceHash.String())
	}

	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, evidence)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func queryAllEvidence(ctx sdk.Context, req msm.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryAllEvidenceParams

	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	evidence := k.GetAllEvidence(ctx)

	start, end := client.Paginate(len(evidence), params.Page, params.Limit, 100)
	if start < 0 || end < 0 {
		evidence = []exported.Evidence{}
	} else {
		evidence = evidence[start:end]
	}

	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, evidence)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}
