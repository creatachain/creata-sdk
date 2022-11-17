package types_test

import (
	"github.com/creatachain/creata-sdk/codec"
	codectypes "github.com/creatachain/creata-sdk/codec/types"
	govtypes "github.com/creatachain/creata-sdk/x/gov/types"
	"github.com/creatachain/creata-sdk/x/icp/core/02-client/types"
	icptmtypes "github.com/creatachain/creata-sdk/x/icp/light-clients/07-augusteum/types"
	icptesting "github.com/creatachain/creata-sdk/x/icp/testing"
)

func (suite *TypesTestSuite) TestNewUpdateClientProposal() {
	p, err := types.NewClientUpdateProposal(icptesting.Title, icptesting.Description, clientID, &icptmtypes.Header{})
	suite.Require().NoError(err)
	suite.Require().NotNil(p)

	p, err = types.NewClientUpdateProposal(icptesting.Title, icptesting.Description, clientID, nil)
	suite.Require().Error(err)
	suite.Require().Nil(p)
}

func (suite *TypesTestSuite) TestValidateBasic() {
	// use solo machine header for testing
	solomachine := icptesting.NewSolomachine(suite.T(), suite.chainA.Codec, clientID, "", 2)
	smHeader := solomachine.CreateHeader()
	header, err := types.PackHeader(smHeader)
	suite.Require().NoError(err)

	// use a different pointer so we don't modify 'header'
	smInvalidHeader := solomachine.CreateHeader()

	// a sequence of 0 will fail basic validation
	smInvalidHeader.Sequence = 0

	invalidHeader, err := types.PackHeader(smInvalidHeader)
	suite.Require().NoError(err)

	testCases := []struct {
		name     string
		proposal govtypes.Content
		expPass  bool
	}{
		{
			"success",
			&types.ClientUpdateProposal{icptesting.Title, icptesting.Description, clientID, header},
			true,
		},
		{
			"fails validate abstract - empty title",
			&types.ClientUpdateProposal{"", icptesting.Description, clientID, header},
			false,
		},
		{
			"fails to unpack header",
			&types.ClientUpdateProposal{icptesting.Title, icptesting.Description, clientID, nil},
			false,
		},
		{
			"fails header validate basic",
			&types.ClientUpdateProposal{icptesting.Title, icptesting.Description, clientID, invalidHeader},
			false,
		},
	}

	for _, tc := range testCases {

		err := tc.proposal.ValidateBasic()

		if tc.expPass {
			suite.Require().NoError(err, tc.name)
		} else {
			suite.Require().Error(err, tc.name)
		}
	}
}

// tests a client update proposal can be marshaled and unmarshaled, and the
// client state can be unpacked
func (suite *TypesTestSuite) TestMarshalClientUpdateProposalProposal() {
	_, err := types.PackHeader(&icptmtypes.Header{})
	suite.Require().NoError(err)

	// create proposal
	header := suite.chainA.CurrentTMClientHeader()
	proposal, err := types.NewClientUpdateProposal("update ICP client", "description", "client-id", header)
	suite.Require().NoError(err)

	// create codec
	ir := codectypes.NewInterfaceRegistry()
	types.RegisterInterfaces(ir)
	govtypes.RegisterInterfaces(ir)
	icptmtypes.RegisterInterfaces(ir)
	cdc := codec.NewProtoCodec(ir)

	// marshal message
	bz, err := cdc.MarshalJSON(proposal)
	suite.Require().NoError(err)

	// unmarshal proposal
	newProposal := &types.ClientUpdateProposal{}
	err = cdc.UnmarshalJSON(bz, newProposal)
	suite.Require().NoError(err)

	// unpack client state
	_, err = types.UnpackHeader(newProposal.Header)
	suite.Require().NoError(err)
}
