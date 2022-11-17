package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	tmproto "github.com/creatachain/augusteum/proto/augusteum/types"

	"github.com/creatachain/creata-sdk/baseapp"
	"github.com/creatachain/creata-sdk/creataapp"
	sdk "github.com/creatachain/creata-sdk/types"
	"github.com/creatachain/creata-sdk/x/staking/keeper"
	"github.com/creatachain/creata-sdk/x/staking/types"
)

type KeeperTestSuite struct {
	suite.Suite

	app         *creataapp.CreataApp
	ctx         sdk.Context
	addrs       []sdk.AccAddress
	vals        []types.Validator
	queryClient types.QueryClient
}

func (suite *KeeperTestSuite) SetupTest() {
	app := creataapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	querier := keeper.Querier{Keeper: app.StakingKeeper}

	queryHelper := baseapp.NewQueryServerTestHelper(ctx, app.InterfaceRegistry())
	types.RegisterQueryServer(queryHelper, querier)
	queryClient := types.NewQueryClient(queryHelper)

	addrs, _, validators := createValidators(suite.T(), ctx, app, []int64{9, 8, 7})
	header := tmproto.Header{
		ChainID: "HelloChain",
		Height:  5,
	}

	// sort a copy of the validators, so that original validators does not
	// have its order changed
	sortedVals := make([]types.Validator, len(validators))
	copy(sortedVals, validators)
	hi := types.NewHistoricalInfo(header, sortedVals)
	app.StakingKeeper.SetHistoricalInfo(ctx, 5, &hi)

	suite.app, suite.ctx, suite.queryClient, suite.addrs, suite.vals = app, ctx, queryClient, addrs, validators
}
func TestParams(t *testing.T) {
	app := creataapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	expParams := types.DefaultParams()

	//check that the empty keeper loads the default
	resParams := app.StakingKeeper.GetParams(ctx)
	require.True(t, expParams.Equal(resParams))

	//modify a params, save, and retrieve
	expParams.MaxValidators = 777
	app.StakingKeeper.SetParams(ctx, expParams)
	resParams = app.StakingKeeper.GetParams(ctx)
	require.True(t, expParams.Equal(resParams))
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
