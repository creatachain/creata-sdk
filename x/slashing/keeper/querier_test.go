package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	msm "github.com/creatachain/augusteum/msm/types"
	tmproto "github.com/creatachain/augusteum/proto/augusteum/types"

	"github.com/creatachain/creata-sdk/codec"
	"github.com/creatachain/creata-sdk/creataapp"
	"github.com/creatachain/creata-sdk/x/slashing/keeper"
	"github.com/creatachain/creata-sdk/x/slashing/testslashing"
	"github.com/creatachain/creata-sdk/x/slashing/types"
)

func TestNewQuerier(t *testing.T) {
	app := creataapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})
	app.SlashingKeeper.SetParams(ctx, testslashing.TestParams())
	legacyQuerierCdc := codec.NewAminoCodec(app.LegacyAmino())
	querier := keeper.NewQuerier(app.SlashingKeeper, legacyQuerierCdc.LegacyAmino)

	query := msm.RequestQuery{
		Path: "",
		Data: []byte{},
	}

	_, err := querier(ctx, []string{types.QueryParameters}, query)
	require.NoError(t, err)
}

func TestQueryParams(t *testing.T) {
	cdc := codec.NewLegacyAmino()
	legacyQuerierCdc := codec.NewAminoCodec(cdc)
	app := creataapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})
	app.SlashingKeeper.SetParams(ctx, testslashing.TestParams())

	querier := keeper.NewQuerier(app.SlashingKeeper, legacyQuerierCdc.LegacyAmino)

	query := msm.RequestQuery{
		Path: "",
		Data: []byte{},
	}

	var params types.Params

	res, err := querier(ctx, []string{types.QueryParameters}, query)
	require.NoError(t, err)

	err = cdc.UnmarshalJSON(res, &params)
	require.NoError(t, err)
	require.Equal(t, app.SlashingKeeper.GetParams(ctx), params)
}
