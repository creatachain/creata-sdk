package keeper_test

import (
	"fmt"
	"testing"

	"github.com/creatachain/creata-sdk/codec"

	"github.com/stretchr/testify/require"

	msm "github.com/creatachain/augusteum/msm/types"

	"github.com/creatachain/creata-sdk/testutil/testdata"
	keep "github.com/creatachain/creata-sdk/x/auth/keeper"
	"github.com/creatachain/creata-sdk/x/auth/types"
)

func TestQueryAccount(t *testing.T) {
	app, ctx := createTestApp(true)
	legacyQuerierCdc := codec.NewAminoCodec(app.LegacyAmino())

	req := msm.RequestQuery{
		Path: "",
		Data: []byte{},
	}

	path := []string{types.QueryAccount}
	querier := keep.NewQuerier(app.AccountKeeper, legacyQuerierCdc.LegacyAmino)

	bz, err := querier(ctx, []string{"other"}, req)
	require.Error(t, err)
	require.Nil(t, bz)

	req = msm.RequestQuery{
		Path: fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryAccount),
		Data: []byte{},
	}
	res, err := querier(ctx, path, req)
	require.Error(t, err)
	require.Nil(t, res)

	req.Data = legacyQuerierCdc.MustMarshalJSON(&types.QueryAccountRequest{Address: ""})
	res, err = querier(ctx, path, req)
	require.Error(t, err)
	require.Nil(t, res)

	_, _, addr := testdata.KeyTestPubAddr()
	req.Data = legacyQuerierCdc.MustMarshalJSON(&types.QueryAccountRequest{Address: addr.String()})
	res, err = querier(ctx, path, req)
	require.Error(t, err)
	require.Nil(t, res)

	app.AccountKeeper.SetAccount(ctx, app.AccountKeeper.NewAccountWithAddress(ctx, addr))
	res, err = querier(ctx, path, req)
	require.NoError(t, err)
	require.NotNil(t, res)

	res, err = querier(ctx, path, req)
	require.NoError(t, err)
	require.NotNil(t, res)

	var account types.AccountI
	err2 := legacyQuerierCdc.LegacyAmino.UnmarshalJSON(res, &account)
	require.Nil(t, err2)
}