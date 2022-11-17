package types_test

import (
	"testing"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/suite"

	"github.com/creatachain/creata-sdk/x/icp/core/02-client/types"
	commitmenttypes "github.com/creatachain/creata-sdk/x/icp/core/23-commitment/types"
	"github.com/creatachain/creata-sdk/x/icp/core/exported"
	solomachinetypes "github.com/creatachain/creata-sdk/x/icp/light-clients/06-solomachine/types"
	icptmtypes "github.com/creatachain/creata-sdk/x/icp/light-clients/07-augusteum/types"
	icptesting "github.com/creatachain/creata-sdk/x/icp/testing"
)

type TypesTestSuite struct {
	suite.Suite

	coordinator *icptesting.Coordinator

	chainA *icptesting.TestChain
	chainB *icptesting.TestChain
}

func (suite *TypesTestSuite) SetupTest() {
	suite.coordinator = icptesting.NewCoordinator(suite.T(), 2)
	suite.chainA = suite.coordinator.GetChain(icptesting.GetChainID(0))
	suite.chainB = suite.coordinator.GetChain(icptesting.GetChainID(1))
}

func TestTypesTestSuite(t *testing.T) {
	suite.Run(t, new(TypesTestSuite))
}

// tests that different clients within MsgCreateClient can be marshaled
// and unmarshaled.
func (suite *TypesTestSuite) TestMarshalMsgCreateClient() {
	var (
		msg *types.MsgCreateClient
		err error
	)

	testCases := []struct {
		name     string
		malleate func()
	}{
		{
			"solo machine client", func() {
				soloMachine := icptesting.NewSolomachine(suite.T(), suite.chainA.Codec, "solomachine", "", 2)
				msg, err = types.NewMsgCreateClient(soloMachine.ClientState(), soloMachine.ConsensusState(), suite.chainA.SenderAccount.GetAddress())
				suite.Require().NoError(err)
			},
		},
		{
			"augusteum client", func() {
				augusteumClient := icptmtypes.NewClientState(suite.chainA.ChainID, icptesting.DefaultTrustLevel, icptesting.TrustingPeriod, icptesting.UnbondingPeriod, icptesting.MaxClockDrift, clientHeight, commitmenttypes.GetSDKSpecs(), icptesting.UpgradePath, false, false)
				msg, err = types.NewMsgCreateClient(augusteumClient, suite.chainA.CurrentTMClientHeader().ConsensusState(), suite.chainA.SenderAccount.GetAddress())
				suite.Require().NoError(err)
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
			bz, err := cdc.MarshalJSON(msg)
			suite.Require().NoError(err)

			// unmarshal message
			newMsg := &types.MsgCreateClient{}
			err = cdc.UnmarshalJSON(bz, newMsg)
			suite.Require().NoError(err)

			suite.Require().True(proto.Equal(msg, newMsg))
		})
	}
}

func (suite *TypesTestSuite) TestMsgCreateClient_ValidateBasic() {
	var (
		msg = &types.MsgCreateClient{}
		err error
	)

	cases := []struct {
		name     string
		malleate func()
		expPass  bool
	}{
		{
			"valid - augusteum client",
			func() {
				augusteumClient := icptmtypes.NewClientState(suite.chainA.ChainID, icptesting.DefaultTrustLevel, icptesting.TrustingPeriod, icptesting.UnbondingPeriod, icptesting.MaxClockDrift, clientHeight, commitmenttypes.GetSDKSpecs(), icptesting.UpgradePath, false, false)
				msg, err = types.NewMsgCreateClient(augusteumClient, suite.chainA.CurrentTMClientHeader().ConsensusState(), suite.chainA.SenderAccount.GetAddress())
				suite.Require().NoError(err)
			},
			true,
		},
		{
			"invalid augusteum client",
			func() {
				msg, err = types.NewMsgCreateClient(&icptmtypes.ClientState{}, suite.chainA.CurrentTMClientHeader().ConsensusState(), suite.chainA.SenderAccount.GetAddress())
				suite.Require().NoError(err)
			},
			false,
		},
		{
			"failed to unpack client",
			func() {
				msg.ClientState = nil
			},
			false,
		},
		{
			"failed to unpack consensus state",
			func() {
				augusteumClient := icptmtypes.NewClientState(suite.chainA.ChainID, icptesting.DefaultTrustLevel, icptesting.TrustingPeriod, icptesting.UnbondingPeriod, icptesting.MaxClockDrift, clientHeight, commitmenttypes.GetSDKSpecs(), icptesting.UpgradePath, false, false)
				msg, err = types.NewMsgCreateClient(augusteumClient, suite.chainA.CurrentTMClientHeader().ConsensusState(), suite.chainA.SenderAccount.GetAddress())
				suite.Require().NoError(err)
				msg.ConsensusState = nil
			},
			false,
		},
		{
			"invalid signer",
			func() {
				msg.Signer = ""
			},
			false,
		},
		{
			"valid - solomachine client",
			func() {
				soloMachine := icptesting.NewSolomachine(suite.T(), suite.chainA.Codec, "solomachine", "", 2)
				msg, err = types.NewMsgCreateClient(soloMachine.ClientState(), soloMachine.ConsensusState(), suite.chainA.SenderAccount.GetAddress())
				suite.Require().NoError(err)
			},
			true,
		},
		{
			"invalid solomachine client",
			func() {
				soloMachine := icptesting.NewSolomachine(suite.T(), suite.chainA.Codec, "solomachine", "", 2)
				msg, err = types.NewMsgCreateClient(&solomachinetypes.ClientState{}, soloMachine.ConsensusState(), suite.chainA.SenderAccount.GetAddress())
				suite.Require().NoError(err)
			},
			false,
		},
		{
			"invalid solomachine consensus state",
			func() {
				soloMachine := icptesting.NewSolomachine(suite.T(), suite.chainA.Codec, "solomachine", "", 2)
				msg, err = types.NewMsgCreateClient(soloMachine.ClientState(), &solomachinetypes.ConsensusState{}, suite.chainA.SenderAccount.GetAddress())
				suite.Require().NoError(err)
			},
			false,
		},
		{
			"invalid - client state and consensus state client types do not match",
			func() {
				augusteumClient := icptmtypes.NewClientState(suite.chainA.ChainID, icptesting.DefaultTrustLevel, icptesting.TrustingPeriod, icptesting.UnbondingPeriod, icptesting.MaxClockDrift, clientHeight, commitmenttypes.GetSDKSpecs(), icptesting.UpgradePath, false, false)
				soloMachine := icptesting.NewSolomachine(suite.T(), suite.chainA.Codec, "solomachine", "", 2)
				msg, err = types.NewMsgCreateClient(augusteumClient, soloMachine.ConsensusState(), suite.chainA.SenderAccount.GetAddress())
				suite.Require().NoError(err)
			},
			false,
		},
	}

	for _, tc := range cases {
		tc.malleate()
		err = msg.ValidateBasic()
		if tc.expPass {
			suite.Require().NoError(err, tc.name)
		} else {
			suite.Require().Error(err, tc.name)
		}
	}
}

// tests that different header within MsgUpdateClient can be marshaled
// and unmarshaled.
func (suite *TypesTestSuite) TestMarshalMsgUpdateClient() {
	var (
		msg *types.MsgUpdateClient
		err error
	)

	testCases := []struct {
		name     string
		malleate func()
	}{
		{
			"solo machine client", func() {
				soloMachine := icptesting.NewSolomachine(suite.T(), suite.chainA.Codec, "solomachine", "", 2)
				msg, err = types.NewMsgUpdateClient(soloMachine.ClientID, soloMachine.CreateHeader(), suite.chainA.SenderAccount.GetAddress())
				suite.Require().NoError(err)
			},
		},
		{
			"augusteum client", func() {
				msg, err = types.NewMsgUpdateClient("augusteum", suite.chainA.CurrentTMClientHeader(), suite.chainA.SenderAccount.GetAddress())
				suite.Require().NoError(err)

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
			bz, err := cdc.MarshalJSON(msg)
			suite.Require().NoError(err)

			// unmarshal message
			newMsg := &types.MsgUpdateClient{}
			err = cdc.UnmarshalJSON(bz, newMsg)
			suite.Require().NoError(err)

			suite.Require().True(proto.Equal(msg, newMsg))
		})
	}
}

func (suite *TypesTestSuite) TestMsgUpdateClient_ValidateBasic() {
	var (
		msg = &types.MsgUpdateClient{}
		err error
	)

	cases := []struct {
		name     string
		malleate func()
		expPass  bool
	}{
		{
			"invalid client-id",
			func() {
				msg.ClientId = ""
			},
			false,
		},
		{
			"valid - augusteum header",
			func() {
				msg, err = types.NewMsgUpdateClient("augusteum", suite.chainA.CurrentTMClientHeader(), suite.chainA.SenderAccount.GetAddress())
				suite.Require().NoError(err)
			},
			true,
		},
		{
			"invalid augusteum header",
			func() {
				msg, err = types.NewMsgUpdateClient("augusteum", &icptmtypes.Header{}, suite.chainA.SenderAccount.GetAddress())
				suite.Require().NoError(err)
			},
			false,
		},
		{
			"failed to unpack header",
			func() {
				msg.Header = nil
			},
			false,
		},
		{
			"invalid signer",
			func() {
				msg.Signer = ""
			},
			false,
		},
		{
			"valid - solomachine header",
			func() {
				soloMachine := icptesting.NewSolomachine(suite.T(), suite.chainA.Codec, "solomachine", "", 2)
				msg, err = types.NewMsgUpdateClient(soloMachine.ClientID, soloMachine.CreateHeader(), suite.chainA.SenderAccount.GetAddress())
				suite.Require().NoError(err)
			},
			true,
		},
		{
			"invalid solomachine header",
			func() {
				msg, err = types.NewMsgUpdateClient("solomachine", &solomachinetypes.Header{}, suite.chainA.SenderAccount.GetAddress())
				suite.Require().NoError(err)
			},
			false,
		},
		{
			"unsupported - localhost",
			func() {
				msg, err = types.NewMsgUpdateClient(exported.Localhost, suite.chainA.CurrentTMClientHeader(), suite.chainA.SenderAccount.GetAddress())
				suite.Require().NoError(err)
			},
			false,
		},
	}

	for _, tc := range cases {
		tc.malleate()
		err = msg.ValidateBasic()
		if tc.expPass {
			suite.Require().NoError(err, tc.name)
		} else {
			suite.Require().Error(err, tc.name)
		}
	}
}

func (suite *TypesTestSuite) TestMarshalMsgUpgradeClient() {
	var (
		msg *types.MsgUpgradeClient
		err error
	)

	testCases := []struct {
		name     string
		malleate func()
	}{
		{
			"client upgrades to new augusteum client",
			func() {
				augusteumClient := icptmtypes.NewClientState(suite.chainA.ChainID, icptesting.DefaultTrustLevel, icptesting.TrustingPeriod, icptesting.UnbondingPeriod, icptesting.MaxClockDrift, clientHeight, commitmenttypes.GetSDKSpecs(), icptesting.UpgradePath, false, false)
				augusteumConsState := &icptmtypes.ConsensusState{NextValidatorsHash: []byte("nextValsHash")}
				msg, err = types.NewMsgUpgradeClient("clientid", augusteumClient, augusteumConsState, []byte("proofUpgradeClient"), []byte("proofUpgradeConsState"), suite.chainA.SenderAccount.GetAddress())
				suite.Require().NoError(err)
			},
		},
		{
			"client upgrades to new solomachine client",
			func() {
				soloMachine := icptesting.NewSolomachine(suite.T(), suite.chainA.Codec, "solomachine", "", 1)
				msg, err = types.NewMsgUpgradeClient("clientid", soloMachine.ClientState(), soloMachine.ConsensusState(), []byte("proofUpgradeClient"), []byte("proofUpgradeConsState"), suite.chainA.SenderAccount.GetAddress())
				suite.Require().NoError(err)
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
			bz, err := cdc.MarshalJSON(msg)
			suite.Require().NoError(err)

			// unmarshal message
			newMsg := &types.MsgUpgradeClient{}
			err = cdc.UnmarshalJSON(bz, newMsg)
			suite.Require().NoError(err)
		})
	}
}

func (suite *TypesTestSuite) TestMsgUpgradeClient_ValidateBasic() {
	cases := []struct {
		name     string
		malleate func(*types.MsgUpgradeClient)
		expPass  bool
	}{
		{
			name:     "success",
			malleate: func(msg *types.MsgUpgradeClient) {},
			expPass:  true,
		},
		{
			name: "client id empty",
			malleate: func(msg *types.MsgUpgradeClient) {
				msg.ClientId = ""
			},
			expPass: false,
		},
		{
			name: "invalid client id",
			malleate: func(msg *types.MsgUpgradeClient) {
				msg.ClientId = "invalid~chain/id"
			},
			expPass: false,
		},
		{
			name: "unpacking clientstate fails",
			malleate: func(msg *types.MsgUpgradeClient) {
				msg.ClientState = nil
			},
			expPass: false,
		},
		{
			name: "unpacking consensus state fails",
			malleate: func(msg *types.MsgUpgradeClient) {
				msg.ConsensusState = nil
			},
			expPass: false,
		},
		{
			name: "client and consensus type does not match",
			malleate: func(msg *types.MsgUpgradeClient) {
				soloMachine := icptesting.NewSolomachine(suite.T(), suite.chainA.Codec, "solomachine", "", 2)
				soloConsensus, err := types.PackConsensusState(soloMachine.ConsensusState())
				suite.Require().NoError(err)
				msg.ConsensusState = soloConsensus
			},
			expPass: false,
		},
		{
			name: "empty client proof",
			malleate: func(msg *types.MsgUpgradeClient) {
				msg.ProofUpgradeClient = nil
			},
			expPass: false,
		},
		{
			name: "empty consensus state proof",
			malleate: func(msg *types.MsgUpgradeClient) {
				msg.ProofUpgradeConsensusState = nil
			},
			expPass: false,
		},
		{
			name: "empty signer",
			malleate: func(msg *types.MsgUpgradeClient) {
				msg.Signer = "  "
			},
			expPass: false,
		},
	}

	for _, tc := range cases {
		tc := tc

		clientState := icptmtypes.NewClientState(suite.chainA.ChainID, icptesting.DefaultTrustLevel, icptesting.TrustingPeriod, icptesting.UnbondingPeriod, icptesting.MaxClockDrift, clientHeight, commitmenttypes.GetSDKSpecs(), icptesting.UpgradePath, false, false)
		consState := &icptmtypes.ConsensusState{NextValidatorsHash: []byte("nextValsHash")}
		msg, err := types.NewMsgUpgradeClient("testclientid", clientState, consState, []byte("proofUpgradeClient"), []byte("proofUpgradeConsState"), suite.chainA.SenderAccount.GetAddress())
		suite.Require().NoError(err)

		tc.malleate(msg)
		err = msg.ValidateBasic()
		if tc.expPass {
			suite.Require().NoError(err, "valid case %s failed", tc.name)
		} else {
			suite.Require().Error(err, "invalid case %s passed", tc.name)
		}
	}
}

// tests that different misbehaviours within MsgSubmitMisbehaviour can be marshaled
// and unmarshaled.
func (suite *TypesTestSuite) TestMarshalMsgSubmitMisbehaviour() {
	var (
		msg *types.MsgSubmitMisbehaviour
		err error
	)

	testCases := []struct {
		name     string
		malleate func()
	}{
		{
			"solo machine client", func() {
				soloMachine := icptesting.NewSolomachine(suite.T(), suite.chainA.Codec, "solomachine", "", 2)
				msg, err = types.NewMsgSubmitMisbehaviour(soloMachine.ClientID, soloMachine.CreateMisbehaviour(), suite.chainA.SenderAccount.GetAddress())
				suite.Require().NoError(err)
			},
		},
		{
			"augusteum client", func() {
				height := types.NewHeight(0, uint64(suite.chainA.CurrentHeader.Height))
				heightMinus1 := types.NewHeight(0, uint64(suite.chainA.CurrentHeader.Height)-1)
				header1 := suite.chainA.CreateTMClientHeader(suite.chainA.ChainID, int64(height.RevisionHeight), heightMinus1, suite.chainA.CurrentHeader.Time, suite.chainA.Vals, suite.chainA.Vals, suite.chainA.Signers)
				header2 := suite.chainA.CreateTMClientHeader(suite.chainA.ChainID, int64(height.RevisionHeight), heightMinus1, suite.chainA.CurrentHeader.Time.Add(time.Minute), suite.chainA.Vals, suite.chainA.Vals, suite.chainA.Signers)

				misbehaviour := icptmtypes.NewMisbehaviour("augusteum", header1, header2)
				msg, err = types.NewMsgSubmitMisbehaviour("augusteum", misbehaviour, suite.chainA.SenderAccount.GetAddress())
				suite.Require().NoError(err)

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
			bz, err := cdc.MarshalJSON(msg)
			suite.Require().NoError(err)

			// unmarshal message
			newMsg := &types.MsgSubmitMisbehaviour{}
			err = cdc.UnmarshalJSON(bz, newMsg)
			suite.Require().NoError(err)

			suite.Require().True(proto.Equal(msg, newMsg))
		})
	}
}

func (suite *TypesTestSuite) TestMsgSubmitMisbehaviour_ValidateBasic() {
	var (
		msg = &types.MsgSubmitMisbehaviour{}
		err error
	)

	cases := []struct {
		name     string
		malleate func()
		expPass  bool
	}{
		{
			"invalid client-id",
			func() {
				msg.ClientId = ""
			},
			false,
		},
		{
			"valid - augusteum misbehaviour",
			func() {
				height := types.NewHeight(0, uint64(suite.chainA.CurrentHeader.Height))
				heightMinus1 := types.NewHeight(0, uint64(suite.chainA.CurrentHeader.Height)-1)
				header1 := suite.chainA.CreateTMClientHeader(suite.chainA.ChainID, int64(height.RevisionHeight), heightMinus1, suite.chainA.CurrentHeader.Time, suite.chainA.Vals, suite.chainA.Vals, suite.chainA.Signers)
				header2 := suite.chainA.CreateTMClientHeader(suite.chainA.ChainID, int64(height.RevisionHeight), heightMinus1, suite.chainA.CurrentHeader.Time.Add(time.Minute), suite.chainA.Vals, suite.chainA.Vals, suite.chainA.Signers)

				misbehaviour := icptmtypes.NewMisbehaviour("augusteum", header1, header2)
				msg, err = types.NewMsgSubmitMisbehaviour("augusteum", misbehaviour, suite.chainA.SenderAccount.GetAddress())
				suite.Require().NoError(err)
			},
			true,
		},
		{
			"invalid augusteum misbehaviour",
			func() {
				msg, err = types.NewMsgSubmitMisbehaviour("augusteum", &icptmtypes.Misbehaviour{}, suite.chainA.SenderAccount.GetAddress())
				suite.Require().NoError(err)
			},
			false,
		},
		{
			"failed to unpack misbehaviourt",
			func() {
				msg.Misbehaviour = nil
			},
			false,
		},
		{
			"invalid signer",
			func() {
				msg.Signer = ""
			},
			false,
		},
		{
			"valid - solomachine misbehaviour",
			func() {
				soloMachine := icptesting.NewSolomachine(suite.T(), suite.chainA.Codec, "solomachine", "", 2)
				msg, err = types.NewMsgSubmitMisbehaviour(soloMachine.ClientID, soloMachine.CreateMisbehaviour(), suite.chainA.SenderAccount.GetAddress())
				suite.Require().NoError(err)
			},
			true,
		},
		{
			"invalid solomachine misbehaviour",
			func() {
				msg, err = types.NewMsgSubmitMisbehaviour("solomachine", &solomachinetypes.Misbehaviour{}, suite.chainA.SenderAccount.GetAddress())
				suite.Require().NoError(err)
			},
			false,
		},
		{
			"client-id mismatch",
			func() {
				soloMachineMisbehaviour := icptesting.NewSolomachine(suite.T(), suite.chainA.Codec, "solomachine", "", 2).CreateMisbehaviour()
				msg, err = types.NewMsgSubmitMisbehaviour("external", soloMachineMisbehaviour, suite.chainA.SenderAccount.GetAddress())
				suite.Require().NoError(err)
			},
			false,
		},
	}

	for _, tc := range cases {
		tc.malleate()
		err = msg.ValidateBasic()
		if tc.expPass {
			suite.Require().NoError(err, tc.name)
		} else {
			suite.Require().Error(err, tc.name)
		}
	}
}
