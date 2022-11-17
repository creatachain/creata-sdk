package keeper_test

import (
	"testing"

	"github.com/creatachain/augusteum/crypto"
	"github.com/stretchr/testify/suite"

	"github.com/creatachain/creata-sdk/baseapp"
	sdk "github.com/creatachain/creata-sdk/types"
	"github.com/creatachain/creata-sdk/x/icp/applications/transfer/types"
	icptesting "github.com/creatachain/creata-sdk/x/icp/testing"
)

type KeeperTestSuite struct {
	suite.Suite

	coordinator *icptesting.Coordinator

	// testing chains used for convenience and readability
	chainA *icptesting.TestChain
	chainB *icptesting.TestChain
	chainC *icptesting.TestChain

	queryClient types.QueryClient
}

func (suite *KeeperTestSuite) SetupTest() {
	suite.coordinator = icptesting.NewCoordinator(suite.T(), 3)
	suite.chainA = suite.coordinator.GetChain(icptesting.GetChainID(0))
	suite.chainB = suite.coordinator.GetChain(icptesting.GetChainID(1))
	suite.chainC = suite.coordinator.GetChain(icptesting.GetChainID(2))

	queryHelper := baseapp.NewQueryServerTestHelper(suite.chainA.GetContext(), suite.chainA.App.InterfaceRegistry())
	types.RegisterQueryServer(queryHelper, suite.chainA.App.TransferKeeper)
	suite.queryClient = types.NewQueryClient(queryHelper)
}

func (suite *KeeperTestSuite) TestGetTransferAccount() {
	expectedMaccAddr := sdk.AccAddress(crypto.AddressHash([]byte(types.ModuleName)))

	macc := suite.chainA.App.TransferKeeper.GetTransferAccount(suite.chainA.GetContext())

	suite.Require().NotNil(macc)
	suite.Require().Equal(types.ModuleName, macc.GetName())
	suite.Require().Equal(expectedMaccAddr, macc.GetAddress())
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
