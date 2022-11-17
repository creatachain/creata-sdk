package types_test

import (
	clienttypes "github.com/creatachain/creata-sdk/x/icp/core/02-client/types"
	channeltypes "github.com/creatachain/creata-sdk/x/icp/core/04-channel/types"
	host "github.com/creatachain/creata-sdk/x/icp/core/24-host"
	"github.com/creatachain/creata-sdk/x/icp/core/exported"
	solomachinetypes "github.com/creatachain/creata-sdk/x/icp/light-clients/06-solomachine/types"
	"github.com/creatachain/creata-sdk/x/icp/light-clients/07-augusteum/types"
	icptesting "github.com/creatachain/creata-sdk/x/icp/testing"
)

func (suite *AugusteumTestSuite) TestGetConsensusState() {
	var (
		height  exported.Height
		clientA string
	)

	testCases := []struct {
		name     string
		malleate func()
		expPass  bool
	}{
		{
			"success", func() {}, true,
		},
		{
			"consensus state not found", func() {
				// use height with no consensus state set
				height = height.(clienttypes.Height).Increment()
			}, false,
		},
		{
			"not a consensus state interface", func() {
				// marshal an empty client state and set as consensus state
				store := suite.chainA.App.ICPKeeper.ClientKeeper.ClientStore(suite.chainA.GetContext(), clientA)
				clientStateBz := suite.chainA.App.ICPKeeper.ClientKeeper.MustMarshalClientState(&types.ClientState{})
				store.Set(host.ConsensusStateKey(height), clientStateBz)
			}, false,
		},
		{
			"invalid consensus state (solomachine)", func() {
				// marshal and set solomachine consensus state
				store := suite.chainA.App.ICPKeeper.ClientKeeper.ClientStore(suite.chainA.GetContext(), clientA)
				consensusStateBz := suite.chainA.App.ICPKeeper.ClientKeeper.MustMarshalConsensusState(&solomachinetypes.ConsensusState{})
				store.Set(host.ConsensusStateKey(height), consensusStateBz)
			}, false,
		},
	}

	for _, tc := range testCases {
		tc := tc

		suite.Run(tc.name, func() {
			suite.SetupTest()

			clientA, _, _, _, _, _ = suite.coordinator.Setup(suite.chainA, suite.chainB, channeltypes.UNORDERED)
			clientState := suite.chainA.GetClientState(clientA)
			height = clientState.GetLatestHeight()

			tc.malleate() // change vars as necessary

			store := suite.chainA.App.ICPKeeper.ClientKeeper.ClientStore(suite.chainA.GetContext(), clientA)
			consensusState, err := types.GetConsensusState(store, suite.chainA.Codec, height)

			if tc.expPass {
				suite.Require().NoError(err)
				expConsensusState, found := suite.chainA.GetConsensusState(clientA, height)
				suite.Require().True(found)
				suite.Require().Equal(expConsensusState, consensusState)
			} else {
				suite.Require().Error(err)
				suite.Require().Nil(consensusState)
			}
		})
	}
}

func (suite *AugusteumTestSuite) TestGetProcessedTime() {
	// Verify ProcessedTime on CreateClient
	// coordinator increments time before creating client
	expectedTime := suite.chainA.CurrentHeader.Time.Add(icptesting.TimeIncrement)

	clientA, err := suite.coordinator.CreateClient(suite.chainA, suite.chainB, exported.Augusteum)
	suite.Require().NoError(err)

	clientState := suite.chainA.GetClientState(clientA)
	height := clientState.GetLatestHeight()

	store := suite.chainA.App.ICPKeeper.ClientKeeper.ClientStore(suite.chainA.GetContext(), clientA)
	actualTime, ok := types.GetProcessedTime(store, height)
	suite.Require().True(ok, "could not retrieve processed time for stored consensus state")
	suite.Require().Equal(uint64(expectedTime.UnixNano()), actualTime, "retrieved processed time is not expected value")

	// Verify ProcessedTime on UpdateClient
	// coordinator increments time before updating client
	expectedTime = suite.chainA.CurrentHeader.Time.Add(icptesting.TimeIncrement)

	err = suite.coordinator.UpdateClient(suite.chainA, suite.chainB, clientA, exported.Augusteum)
	suite.Require().NoError(err)

	clientState = suite.chainA.GetClientState(clientA)
	height = clientState.GetLatestHeight()

	store = suite.chainA.App.ICPKeeper.ClientKeeper.ClientStore(suite.chainA.GetContext(), clientA)
	actualTime, ok = types.GetProcessedTime(store, height)
	suite.Require().True(ok, "could not retrieve processed time for stored consensus state")
	suite.Require().Equal(uint64(expectedTime.UnixNano()), actualTime, "retrieved processed time is not expected value")

	// try to get processed time for height that doesn't exist in store
	_, ok = types.GetProcessedTime(store, clienttypes.NewHeight(1, 1))
	suite.Require().False(ok, "retrieved processed time for a non-existent consensus state")
}
