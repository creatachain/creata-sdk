package keeper_test

import (
	"github.com/creatachain/creata-sdk/x/icp/core/02-client/types"
)

func (suite *KeeperTestSuite) TestParams() {
	expParams := types.DefaultParams()

	params := suite.chainA.App.ICPKeeper.ClientKeeper.GetParams(suite.chainA.GetContext())
	suite.Require().Equal(expParams, params)

	expParams.AllowedClients = []string{}
	suite.chainA.App.ICPKeeper.ClientKeeper.SetParams(suite.chainA.GetContext(), expParams)
	params = suite.chainA.App.ICPKeeper.ClientKeeper.GetParams(suite.chainA.GetContext())
	suite.Require().Empty(expParams.AllowedClients)
}
