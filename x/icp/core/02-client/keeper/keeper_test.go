package keeper_test

import (
	"math/rand"
	"testing"
	"time"

	tmbytes "github.com/creatachain/augusteum/libs/bytes"
	tmproto "github.com/creatachain/augusteum/proto/augusteum/types"
	tmtypes "github.com/creatachain/augusteum/types"
	"github.com/stretchr/testify/suite"

	"github.com/creatachain/creata-sdk/baseapp"
	"github.com/creatachain/creata-sdk/codec"
	"github.com/creatachain/creata-sdk/creataapp"
	cryptocodec "github.com/creatachain/creata-sdk/crypto/codec"
	sdk "github.com/creatachain/creata-sdk/types"
	"github.com/creatachain/creata-sdk/x/icp/core/02-client/keeper"
	"github.com/creatachain/creata-sdk/x/icp/core/02-client/types"
	commitmenttypes "github.com/creatachain/creata-sdk/x/icp/core/23-commitment/types"
	"github.com/creatachain/creata-sdk/x/icp/core/exported"
	icptmtypes "github.com/creatachain/creata-sdk/x/icp/light-clients/07-augusteum/types"
	localhosttypes "github.com/creatachain/creata-sdk/x/icp/light-clients/09-localhost/types"
	icptesting "github.com/creatachain/creata-sdk/x/icp/testing"
	icptestingmock "github.com/creatachain/creata-sdk/x/icp/testing/mock"
	stakingtypes "github.com/creatachain/creata-sdk/x/staking/types"
)

const (
	testChainID          = "creatahub-0"
	testChainIDRevision1 = "creatahub-1"

	testClientID  = "augusteum-0"
	testClientID2 = "augusteum-1"
	testClientID3 = "augusteum-2"

	height = 5

	trustingPeriod time.Duration = time.Hour * 24 * 7 * 2
	ubdPeriod      time.Duration = time.Hour * 24 * 7 * 3
	maxClockDrift  time.Duration = time.Second * 10
)

var (
	testClientHeight          = types.NewHeight(0, 5)
	testClientHeightRevision1 = types.NewHeight(1, 5)
	newClientHeight           = types.NewHeight(1, 1)
)

type KeeperTestSuite struct {
	suite.Suite

	coordinator *icptesting.Coordinator

	chainA *icptesting.TestChain
	chainB *icptesting.TestChain

	cdc            codec.Marshaler
	ctx            sdk.Context
	keeper         *keeper.Keeper
	consensusState *icptmtypes.ConsensusState
	header         *icptmtypes.Header
	valSet         *tmtypes.ValidatorSet
	valSetHash     tmbytes.HexBytes
	privVal        tmtypes.PrivValidator
	now            time.Time
	past           time.Time

	queryClient types.QueryClient
}

func (suite *KeeperTestSuite) SetupTest() {
	suite.coordinator = icptesting.NewCoordinator(suite.T(), 2)

	suite.chainA = suite.coordinator.GetChain(icptesting.GetChainID(0))
	suite.chainB = suite.coordinator.GetChain(icptesting.GetChainID(1))

	isCheckTx := false
	suite.now = time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC)
	suite.past = time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	now2 := suite.now.Add(time.Hour)
	app := creataapp.Setup(isCheckTx)

	suite.cdc = app.AppCodec()
	suite.ctx = app.BaseApp.NewContext(isCheckTx, tmproto.Header{Height: height, ChainID: testClientID, Time: now2})
	suite.keeper = &app.ICPKeeper.ClientKeeper
	suite.privVal = icptestingmock.NewPV()

	pubKey, err := suite.privVal.GetPubKey()
	suite.Require().NoError(err)

	testClientHeightMinus1 := types.NewHeight(0, height-1)

	validator := tmtypes.NewValidator(pubKey, 1)
	suite.valSet = tmtypes.NewValidatorSet([]*tmtypes.Validator{validator})
	suite.valSetHash = suite.valSet.Hash()
	suite.header = suite.chainA.CreateTMClientHeader(testChainID, int64(testClientHeight.RevisionHeight), testClientHeightMinus1, now2, suite.valSet, suite.valSet, []tmtypes.PrivValidator{suite.privVal})
	suite.consensusState = icptmtypes.NewConsensusState(suite.now, commitmenttypes.NewMerkleRoot([]byte("hash")), suite.valSetHash)

	var validators stakingtypes.Validators
	for i := 1; i < 11; i++ {
		privVal := icptestingmock.NewPV()
		tmPk, err := privVal.GetPubKey()
		suite.Require().NoError(err)
		pk, err := cryptocodec.FromTmPubKeyInterface(tmPk)
		suite.Require().NoError(err)
		val, err := stakingtypes.NewValidator(sdk.ValAddress(pk.Address()), pk, stakingtypes.Description{})
		suite.Require().NoError(err)

		val.Status = stakingtypes.Bonded
		val.Tokens = sdk.NewInt(rand.Int63())
		validators = append(validators, val)

		hi := stakingtypes.NewHistoricalInfo(suite.ctx.BlockHeader(), validators)
		app.StakingKeeper.SetHistoricalInfo(suite.ctx, int64(i), &hi)
	}

	// add localhost client
	revision := types.ParseChainID(suite.chainA.ChainID)
	localHostClient := localhosttypes.NewClientState(
		suite.chainA.ChainID, types.NewHeight(revision, uint64(suite.chainA.GetContext().BlockHeight())),
	)
	suite.chainA.App.ICPKeeper.ClientKeeper.SetClientState(suite.chainA.GetContext(), exported.Localhost, localHostClient)

	queryHelper := baseapp.NewQueryServerTestHelper(suite.ctx, app.InterfaceRegistry())
	types.RegisterQueryServer(queryHelper, app.ICPKeeper.ClientKeeper)
	suite.queryClient = types.NewQueryClient(queryHelper)
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (suite *KeeperTestSuite) TestSetClientState() {
	clientState := icptmtypes.NewClientState(testChainID, icptmtypes.DefaultTrustLevel, trustingPeriod, ubdPeriod, maxClockDrift, types.ZeroHeight(), commitmenttypes.GetSDKSpecs(), icptesting.UpgradePath, false, false)
	suite.keeper.SetClientState(suite.ctx, testClientID, clientState)

	retrievedState, found := suite.keeper.GetClientState(suite.ctx, testClientID)
	suite.Require().True(found, "GetClientState failed")
	suite.Require().Equal(clientState, retrievedState, "Client states are not equal")
}

func (suite *KeeperTestSuite) TestSetClientConsensusState() {
	suite.keeper.SetClientConsensusState(suite.ctx, testClientID, testClientHeight, suite.consensusState)

	retrievedConsState, found := suite.keeper.GetClientConsensusState(suite.ctx, testClientID, testClientHeight)
	suite.Require().True(found, "GetConsensusState failed")

	tmConsState, ok := retrievedConsState.(*icptmtypes.ConsensusState)
	suite.Require().True(ok)
	suite.Require().Equal(suite.consensusState, tmConsState, "ConsensusState not stored correctly")
}

func (suite *KeeperTestSuite) TestValidateSelfClient() {
	testClientHeight := types.NewHeight(0, uint64(suite.chainA.GetContext().BlockHeight()-1))

	testCases := []struct {
		name        string
		clientState exported.ClientState
		expPass     bool
	}{
		{
			"success",
			icptmtypes.NewClientState(suite.chainA.ChainID, icptmtypes.DefaultTrustLevel, trustingPeriod, ubdPeriod, maxClockDrift, testClientHeight, commitmenttypes.GetSDKSpecs(), icptesting.UpgradePath, false, false),
			true,
		},
		{
			"success with nil UpgradePath",
			icptmtypes.NewClientState(suite.chainA.ChainID, icptmtypes.DefaultTrustLevel, trustingPeriod, ubdPeriod, maxClockDrift, testClientHeight, commitmenttypes.GetSDKSpecs(), nil, false, false),
			true,
		},
		{
			"invalid client type",
			localhosttypes.NewClientState(suite.chainA.ChainID, testClientHeight),
			false,
		},
		{
			"frozen client",
			&icptmtypes.ClientState{suite.chainA.ChainID, icptmtypes.DefaultTrustLevel, trustingPeriod, ubdPeriod, maxClockDrift, testClientHeight, testClientHeight, commitmenttypes.GetSDKSpecs(), icptesting.UpgradePath, false, false},
			false,
		},
		{
			"incorrect chainID",
			icptmtypes.NewClientState("creatatestnet", icptmtypes.DefaultTrustLevel, trustingPeriod, ubdPeriod, maxClockDrift, testClientHeight, commitmenttypes.GetSDKSpecs(), icptesting.UpgradePath, false, false),
			false,
		},
		{
			"invalid client height",
			icptmtypes.NewClientState(suite.chainA.ChainID, icptmtypes.DefaultTrustLevel, trustingPeriod, ubdPeriod, maxClockDrift, types.NewHeight(0, uint64(suite.chainA.GetContext().BlockHeight())), commitmenttypes.GetSDKSpecs(), icptesting.UpgradePath, false, false),
			false,
		},
		{
			"invalid client revision",
			icptmtypes.NewClientState(suite.chainA.ChainID, icptmtypes.DefaultTrustLevel, trustingPeriod, ubdPeriod, maxClockDrift, testClientHeightRevision1, commitmenttypes.GetSDKSpecs(), icptesting.UpgradePath, false, false),
			false,
		},
		{
			"invalid proof specs",
			icptmtypes.NewClientState(suite.chainA.ChainID, icptmtypes.DefaultTrustLevel, trustingPeriod, ubdPeriod, maxClockDrift, testClientHeight, nil, icptesting.UpgradePath, false, false),
			false,
		},
		{
			"invalid trust level",
			icptmtypes.NewClientState(suite.chainA.ChainID, icptmtypes.Fraction{0, 1}, trustingPeriod, ubdPeriod, maxClockDrift, testClientHeight, commitmenttypes.GetSDKSpecs(), icptesting.UpgradePath, false, false),
			false,
		},
		{
			"invalid unbonding period",
			icptmtypes.NewClientState(suite.chainA.ChainID, icptmtypes.DefaultTrustLevel, trustingPeriod, ubdPeriod+10, maxClockDrift, testClientHeight, commitmenttypes.GetSDKSpecs(), icptesting.UpgradePath, false, false),
			false,
		},
		{
			"invalid trusting period",
			icptmtypes.NewClientState(suite.chainA.ChainID, icptmtypes.DefaultTrustLevel, ubdPeriod+10, ubdPeriod, maxClockDrift, testClientHeight, commitmenttypes.GetSDKSpecs(), icptesting.UpgradePath, false, false),
			false,
		},
		{
			"invalid upgrade path",
			icptmtypes.NewClientState(suite.chainA.ChainID, icptmtypes.DefaultTrustLevel, trustingPeriod, ubdPeriod, maxClockDrift, testClientHeight, commitmenttypes.GetSDKSpecs(), []string{"bad", "upgrade", "path"}, false, false),
			false,
		},
	}

	for _, tc := range testCases {
		err := suite.chainA.App.ICPKeeper.ClientKeeper.ValidateSelfClient(suite.chainA.GetContext(), tc.clientState)
		if tc.expPass {
			suite.Require().NoError(err, "expected valid client for case: %s", tc.name)
		} else {
			suite.Require().Error(err, "expected invalid client for case: %s", tc.name)
		}
	}
}

func (suite KeeperTestSuite) TestGetAllGenesisClients() {
	clientIDs := []string{
		testClientID2, testClientID3, testClientID,
	}
	expClients := []exported.ClientState{
		icptmtypes.NewClientState(testChainID, icptmtypes.DefaultTrustLevel, trustingPeriod, ubdPeriod, maxClockDrift, types.ZeroHeight(), commitmenttypes.GetSDKSpecs(), icptesting.UpgradePath, false, false),
		icptmtypes.NewClientState(testChainID, icptmtypes.DefaultTrustLevel, trustingPeriod, ubdPeriod, maxClockDrift, types.ZeroHeight(), commitmenttypes.GetSDKSpecs(), icptesting.UpgradePath, false, false),
		icptmtypes.NewClientState(testChainID, icptmtypes.DefaultTrustLevel, trustingPeriod, ubdPeriod, maxClockDrift, types.ZeroHeight(), commitmenttypes.GetSDKSpecs(), icptesting.UpgradePath, false, false),
	}

	expGenClients := make(types.IdentifiedClientStates, len(expClients))

	for i := range expClients {
		suite.chainA.App.ICPKeeper.ClientKeeper.SetClientState(suite.chainA.GetContext(), clientIDs[i], expClients[i])
		expGenClients[i] = types.NewIdentifiedClientState(clientIDs[i], expClients[i])
	}

	// add localhost client
	localHostClient, found := suite.chainA.App.ICPKeeper.ClientKeeper.GetClientState(suite.chainA.GetContext(), exported.Localhost)
	suite.Require().True(found)
	expGenClients = append(expGenClients, types.NewIdentifiedClientState(exported.Localhost, localHostClient))

	genClients := suite.chainA.App.ICPKeeper.ClientKeeper.GetAllGenesisClients(suite.chainA.GetContext())

	suite.Require().Equal(expGenClients.Sort(), genClients)
}

func (suite KeeperTestSuite) TestGetAllGenesisMetadata() {
	expectedGenMetadata := []types.IdentifiedGenesisMetadata{
		types.NewIdentifiedGenesisMetadata(
			"clientA",
			[]types.GenesisMetadata{
				types.NewGenesisMetadata(icptmtypes.ProcessedTimeKey(types.NewHeight(0, 1)), []byte("foo")),
				types.NewGenesisMetadata(icptmtypes.ProcessedTimeKey(types.NewHeight(0, 2)), []byte("bar")),
				types.NewGenesisMetadata(icptmtypes.ProcessedTimeKey(types.NewHeight(0, 3)), []byte("baz")),
			},
		),
		types.NewIdentifiedGenesisMetadata(
			"clientB",
			[]types.GenesisMetadata{
				types.NewGenesisMetadata(icptmtypes.ProcessedTimeKey(types.NewHeight(1, 100)), []byte("val1")),
				types.NewGenesisMetadata(icptmtypes.ProcessedTimeKey(types.NewHeight(2, 300)), []byte("val2")),
			},
		),
	}

	genClients := []types.IdentifiedClientState{
		types.NewIdentifiedClientState("clientA", &icptmtypes.ClientState{}), types.NewIdentifiedClientState("clientB", &icptmtypes.ClientState{}),
		types.NewIdentifiedClientState("clientC", &icptmtypes.ClientState{}), types.NewIdentifiedClientState("clientD", &localhosttypes.ClientState{}),
	}

	suite.chainA.App.ICPKeeper.ClientKeeper.SetAllClientMetadata(suite.chainA.GetContext(), expectedGenMetadata)

	actualGenMetadata, err := suite.chainA.App.ICPKeeper.ClientKeeper.GetAllClientMetadata(suite.chainA.GetContext(), genClients)
	suite.Require().NoError(err, "get client metadata returned error unexpectedly")
	suite.Require().Equal(expectedGenMetadata, actualGenMetadata, "retrieved metadata is unexpected")
}

func (suite KeeperTestSuite) TestGetConsensusState() {
	suite.ctx = suite.ctx.WithBlockHeight(10)
	cases := []struct {
		name    string
		height  types.Height
		expPass bool
	}{
		{"zero height", types.ZeroHeight(), false},
		{"height > latest height", types.NewHeight(0, uint64(suite.ctx.BlockHeight())+1), false},
		{"latest height - 1", types.NewHeight(0, uint64(suite.ctx.BlockHeight())-1), true},
		{"latest height", types.GetSelfHeight(suite.ctx), true},
	}

	for i, tc := range cases {
		tc := tc
		cs, found := suite.keeper.GetSelfConsensusState(suite.ctx, tc.height)
		if tc.expPass {
			suite.Require().True(found, "Case %d should have passed: %s", i, tc.name)
			suite.Require().NotNil(cs, "Case %d should have passed: %s", i, tc.name)
		} else {
			suite.Require().False(found, "Case %d should have failed: %s", i, tc.name)
			suite.Require().Nil(cs, "Case %d should have failed: %s", i, tc.name)
		}
	}
}

func (suite KeeperTestSuite) TestConsensusStateHelpers() {
	// initial setup
	clientState := icptmtypes.NewClientState(testChainID, icptmtypes.DefaultTrustLevel, trustingPeriod, ubdPeriod, maxClockDrift, testClientHeight, commitmenttypes.GetSDKSpecs(), icptesting.UpgradePath, false, false)

	suite.keeper.SetClientState(suite.ctx, testClientID, clientState)
	suite.keeper.SetClientConsensusState(suite.ctx, testClientID, testClientHeight, suite.consensusState)

	nextState := icptmtypes.NewConsensusState(suite.now, commitmenttypes.NewMerkleRoot([]byte("next")), suite.valSetHash)

	testClientHeightPlus5 := types.NewHeight(0, height+5)

	header := suite.chainA.CreateTMClientHeader(testClientID, int64(testClientHeightPlus5.RevisionHeight), testClientHeight, suite.header.Header.Time.Add(time.Minute),
		suite.valSet, suite.valSet, []tmtypes.PrivValidator{suite.privVal})

	// mock update functionality
	clientState.LatestHeight = header.GetHeight().(types.Height)
	suite.keeper.SetClientConsensusState(suite.ctx, testClientID, header.GetHeight(), nextState)
	suite.keeper.SetClientState(suite.ctx, testClientID, clientState)

	latest, ok := suite.keeper.GetLatestClientConsensusState(suite.ctx, testClientID)
	suite.Require().True(ok)
	suite.Require().Equal(nextState, latest, "Latest client not returned correctly")
}

// 2 clients in total are created on chainA. The first client is updated so it contains an initial consensus state
// and a consensus state at the update height.
func (suite KeeperTestSuite) TestGetAllConsensusStates() {
	clientA, _ := suite.coordinator.SetupClients(suite.chainA, suite.chainB, exported.Augusteum)

	clientState := suite.chainA.GetClientState(clientA)
	expConsensusHeight0 := clientState.GetLatestHeight()
	consensusState0, ok := suite.chainA.GetConsensusState(clientA, expConsensusHeight0)
	suite.Require().True(ok)

	// update client to create a second consensus state
	err := suite.coordinator.UpdateClient(suite.chainA, suite.chainB, clientA, exported.Augusteum)
	suite.Require().NoError(err)

	clientState = suite.chainA.GetClientState(clientA)
	expConsensusHeight1 := clientState.GetLatestHeight()
	suite.Require().True(expConsensusHeight1.GT(expConsensusHeight0))
	consensusState1, ok := suite.chainA.GetConsensusState(clientA, expConsensusHeight1)
	suite.Require().True(ok)

	expConsensus := []exported.ConsensusState{
		consensusState0,
		consensusState1,
	}

	// create second client on chainA
	clientA2, _ := suite.coordinator.SetupClients(suite.chainA, suite.chainB, exported.Augusteum)
	clientState = suite.chainA.GetClientState(clientA2)

	expConsensusHeight2 := clientState.GetLatestHeight()
	consensusState2, ok := suite.chainA.GetConsensusState(clientA2, expConsensusHeight2)
	suite.Require().True(ok)

	expConsensus2 := []exported.ConsensusState{consensusState2}

	expConsensusStates := types.ClientsConsensusStates{
		types.NewClientConsensusStates(clientA, []types.ConsensusStateWithHeight{
			types.NewConsensusStateWithHeight(expConsensusHeight0.(types.Height), expConsensus[0]),
			types.NewConsensusStateWithHeight(expConsensusHeight1.(types.Height), expConsensus[1]),
		}),
		types.NewClientConsensusStates(clientA2, []types.ConsensusStateWithHeight{
			types.NewConsensusStateWithHeight(expConsensusHeight2.(types.Height), expConsensus2[0]),
		}),
	}.Sort()

	consStates := suite.chainA.App.ICPKeeper.ClientKeeper.GetAllConsensusStates(suite.chainA.GetContext())
	suite.Require().Equal(expConsensusStates, consStates, "%s \n\n%s", expConsensusStates, consStates)
}
