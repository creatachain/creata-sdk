package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/creatachain/creata-sdk/x/icp/core/02-client/types"
	"github.com/creatachain/creata-sdk/x/icp/core/exported"
	icptesting "github.com/creatachain/creata-sdk/x/icp/testing"
)

func (suite *TypesTestSuite) TestMarshalConsensusStateWithHeight() {
	var (
		cswh types.ConsensusStateWithHeight
	)

	testCases := []struct {
		name     string
		malleate func()
	}{
		{
			"solo machine client", func() {
				soloMachine := icptesting.NewSolomachine(suite.T(), suite.chainA.Codec, "solomachine", "", 1)
				cswh = types.NewConsensusStateWithHeight(types.NewHeight(0, soloMachine.Sequence), soloMachine.ConsensusState())
			},
		},
		{
			"augusteum client", func() {
				clientA, _ := suite.coordinator.SetupClients(suite.chainA, suite.chainB, exported.Augusteum)
				clientState := suite.chainA.GetClientState(clientA)
				consensusState, ok := suite.chainA.GetConsensusState(clientA, clientState.GetLatestHeight())
				suite.Require().True(ok)

				cswh = types.NewConsensusStateWithHeight(clientState.GetLatestHeight().(types.Height), consensusState)
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		suite.Run(tc.name, func() {
			suite.SetupTest()

			tc.malleate()

			cdc := suite.chainA.App.AppCodec()

			// marshal message
			bz, err := cdc.MarshalJSON(&cswh)
			suite.Require().NoError(err)

			// unmarshal message
			newCswh := &types.ConsensusStateWithHeight{}
			err = cdc.UnmarshalJSON(bz, newCswh)
			suite.Require().NoError(err)
		})
	}
}

func TestValidateClientType(t *testing.T) {
	testCases := []struct {
		name       string
		clientType string
		expPass    bool
	}{
		{"valid", "augusteum", true},
		{"valid solomachine", "solomachine-v1", true},
		{"too large", "augusteumaugusteumaugusteumaugusteumaugusteumt", false},
		{"too short", "t", false},
		{"blank id", "               ", false},
		{"empty id", "", false},
		{"ends with dash", "augusteum-", false},
	}

	for _, tc := range testCases {

		err := types.ValidateClientType(tc.clientType)

		if tc.expPass {
			require.NoError(t, err, tc.name)
		} else {
			require.Error(t, err, tc.name)
		}
	}
}
