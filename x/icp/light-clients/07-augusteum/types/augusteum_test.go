package types_test

import (
	"testing"
	"time"

	tmbytes "github.com/creatachain/augusteum/libs/bytes"
	tmproto "github.com/creatachain/augusteum/proto/augusteum/types"
	tmtypes "github.com/creatachain/augusteum/types"
	"github.com/stretchr/testify/suite"

	"github.com/creatachain/creata-sdk/codec"
	"github.com/creatachain/creata-sdk/creataapp"
	sdk "github.com/creatachain/creata-sdk/types"
	clienttypes "github.com/creatachain/creata-sdk/x/icp/core/02-client/types"
	icptmtypes "github.com/creatachain/creata-sdk/x/icp/light-clients/07-augusteum/types"
	icptesting "github.com/creatachain/creata-sdk/x/icp/testing"
	icptestingmock "github.com/creatachain/creata-sdk/x/icp/testing/mock"
)

const (
	chainID                        = "creata"
	chainIDRevision0               = "creata-revision-0"
	chainIDRevision1               = "creata-revision-1"
	clientID                       = "creatamainnet"
	trustingPeriod   time.Duration = time.Hour * 24 * 7 * 2
	ubdPeriod        time.Duration = time.Hour * 24 * 7 * 3
	maxClockDrift    time.Duration = time.Second * 10
)

var (
	height          = clienttypes.NewHeight(0, 4)
	newClientHeight = clienttypes.NewHeight(1, 1)
	upgradePath     = []string{"upgrade", "upgradedICPState"}
)

type AugusteumTestSuite struct {
	suite.Suite

	coordinator *icptesting.Coordinator

	// testing chains used for convenience and readability
	chainA *icptesting.TestChain
	chainB *icptesting.TestChain

	// TODO: deprecate usage in favor of testing package
	ctx        sdk.Context
	cdc        codec.Marshaler
	privVal    tmtypes.PrivValidator
	valSet     *tmtypes.ValidatorSet
	valsHash   tmbytes.HexBytes
	header     *icptmtypes.Header
	now        time.Time
	headerTime time.Time
	clientTime time.Time
}

func (suite *AugusteumTestSuite) SetupTest() {
	suite.coordinator = icptesting.NewCoordinator(suite.T(), 2)
	suite.chainA = suite.coordinator.GetChain(icptesting.GetChainID(0))
	suite.chainB = suite.coordinator.GetChain(icptesting.GetChainID(1))
	// commit some blocks so that QueryProof returns valid proof (cannot return valid query if height <= 1)
	suite.coordinator.CommitNBlocks(suite.chainA, 2)
	suite.coordinator.CommitNBlocks(suite.chainB, 2)

	// TODO: deprecate usage in favor of testing package
	checkTx := false
	app := creataapp.Setup(checkTx)

	suite.cdc = app.AppCodec()

	// now is the time of the current chain, must be after the updating header
	// mocks ctx.BlockTime()
	suite.now = time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC)
	suite.clientTime = time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	// Header time is intended to be time for any new header used for updates
	suite.headerTime = time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC)

	suite.privVal = icptestingmock.NewPV()

	pubKey, err := suite.privVal.GetPubKey()
	suite.Require().NoError(err)

	heightMinus1 := clienttypes.NewHeight(0, height.RevisionHeight-1)

	val := tmtypes.NewValidator(pubKey, 10)
	suite.valSet = tmtypes.NewValidatorSet([]*tmtypes.Validator{val})
	suite.valsHash = suite.valSet.Hash()
	suite.header = suite.chainA.CreateTMClientHeader(chainID, int64(height.RevisionHeight), heightMinus1, suite.now, suite.valSet, suite.valSet, []tmtypes.PrivValidator{suite.privVal})
	suite.ctx = app.BaseApp.NewContext(checkTx, tmproto.Header{Height: 1, Time: suite.now})
}

func TestAugusteumTestSuite(t *testing.T) {
	suite.Run(t, new(AugusteumTestSuite))
}
