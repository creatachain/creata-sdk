package client_test

import (
	sdk "github.com/creatachain/creata-sdk/types"
	distributiontypes "github.com/creatachain/creata-sdk/x/distribution/types"
	govtypes "github.com/creatachain/creata-sdk/x/gov/types"
	client "github.com/creatachain/creata-sdk/x/icp/core/02-client"
	clienttypes "github.com/creatachain/creata-sdk/x/icp/core/02-client/types"
	"github.com/creatachain/creata-sdk/x/icp/core/exported"
	icptmtypes "github.com/creatachain/creata-sdk/x/icp/light-clients/07-augusteum/types"
	icptesting "github.com/creatachain/creata-sdk/x/icp/testing"
)

func (suite *ClientTestSuite) TestNewClientUpdateProposalHandler() {
	var (
		content govtypes.Content
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
			"nil proposal", func() {
				content = nil
			}, false,
		},
		{
			"unsupported proposal type", func() {
				content = distributiontypes.NewCommunityPoolSpendProposal(icptesting.Title, icptesting.Description, suite.chainA.SenderAccount.GetAddress(), sdk.NewCoins(sdk.NewCoin("communityfunds", sdk.NewInt(10))))
			}, false,
		},
	}

	for _, tc := range testCases {
		tc := tc

		suite.Run(tc.name, func() {
			suite.SetupTest() // reset

			tc.malleate()

			proposalHandler := client.NewClientUpdateProposalHandler(suite.chainA.App.ICPKeeper.ClientKeeper)

			err = proposalHandler(suite.chainA.GetContext(), content)

			if tc.expPass {
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(err)
			}
		})
	}

}
