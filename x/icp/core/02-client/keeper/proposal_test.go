package keeper_test

import (
	"github.com/creatachain/creata-sdk/x/icp/core/02-client/types"
	clienttypes "github.com/creatachain/creata-sdk/x/icp/core/02-client/types"
	"github.com/creatachain/creata-sdk/x/icp/core/exported"
	icptmtypes "github.com/creatachain/creata-sdk/x/icp/light-clients/07-augusteum/types"
	icptesting "github.com/creatachain/creata-sdk/x/icp/testing"
)

func (suite *KeeperTestSuite) TestClientUpdateProposal() {
	var (
		content *types.ClientUpdateProposal
		err     error
	)

	testCases := []struct {
		name     string
		malleate func()
		expPass  bool
	}{
		{
			"valid update client proposal", func() {
				clientA, _ := suite.coordinator.SetupClients(suite.chainA, suite.chainB, exported.Augusteum)
				clientState := suite.chainA.GetClientState(clientA)

				tmClientState, ok := clientState.(*icptmtypes.ClientState)
				suite.Require().True(ok)
				tmClientState.AllowUpdateAfterMisbehaviour = true
				tmClientState.FrozenHeight = tmClientState.LatestHeight
				suite.chainA.App.ICPKeeper.ClientKeeper.SetClientState(suite.chainA.GetContext(), clientA, tmClientState)

				// use next header for chainB to update the client on chainA
				header, err := suite.chainA.ConstructUpdateTMClientHeader(suite.chainB, clientA)
				suite.Require().NoError(err)

				content, err = clienttypes.NewClientUpdateProposal(icptesting.Title, icptesting.Description, clientA, header)
				suite.Require().NoError(err)
			}, true,
		},
		{
			"client type does not exist", func() {
				content, err = clienttypes.NewClientUpdateProposal(icptesting.Title, icptesting.Description, icptesting.InvalidID, &icptmtypes.Header{})
				suite.Require().NoError(err)
			}, false,
		},
		{
			"cannot update localhost", func() {
				content, err = clienttypes.NewClientUpdateProposal(icptesting.Title, icptesting.Description, exported.Localhost, &icptmtypes.Header{})
				suite.Require().NoError(err)
			}, false,
		},
		{
			"client does not exist", func() {
				content, err = clienttypes.NewClientUpdateProposal(icptesting.Title, icptesting.Description, icptesting.InvalidID, &icptmtypes.Header{})
				suite.Require().NoError(err)
			}, false,
		},
		{
			"cannot unpack header, header is nil", func() {
				clientA, _ := suite.coordinator.SetupClients(suite.chainA, suite.chainB, exported.Augusteum)
				content = &clienttypes.ClientUpdateProposal{icptesting.Title, icptesting.Description, clientA, nil}
			}, false,
		},
		{
			"update fails", func() {
				header := &icptmtypes.Header{}
				clientA, _ := suite.coordinator.SetupClients(suite.chainA, suite.chainB, exported.Augusteum)
				content, err = clienttypes.NewClientUpdateProposal(icptesting.Title, icptesting.Description, clientA, header)
				suite.Require().NoError(err)
			}, false,
		},
	}

	for _, tc := range testCases {
		tc := tc

		suite.Run(tc.name, func() {
			suite.SetupTest() // reset

			tc.malleate()

			err = suite.chainA.App.ICPKeeper.ClientKeeper.ClientUpdateProposal(suite.chainA.GetContext(), content)

			if tc.expPass {
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(err)
			}
		})
	}

}
