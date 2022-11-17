package types_test

import (
	"testing"

	tmproto "github.com/creatachain/augusteum/proto/augusteum/types"
	"github.com/stretchr/testify/suite"

	"github.com/creatachain/creata-sdk/codec"
	"github.com/creatachain/creata-sdk/creataapp"
	sdk "github.com/creatachain/creata-sdk/types"
	clienttypes "github.com/creatachain/creata-sdk/x/icp/core/02-client/types"
	"github.com/creatachain/creata-sdk/x/icp/core/exported"
)

const (
	height = 4
)

var (
	clientHeight = clienttypes.NewHeight(0, 10)
)

type LocalhostTestSuite struct {
	suite.Suite

	cdc   codec.Marshaler
	ctx   sdk.Context
	store sdk.KVStore
}

func (suite *LocalhostTestSuite) SetupTest() {
	isCheckTx := false
	app := creataapp.Setup(isCheckTx)

	suite.cdc = app.AppCodec()
	suite.ctx = app.BaseApp.NewContext(isCheckTx, tmproto.Header{Height: 1, ChainID: "icp-chain"})
	suite.store = app.ICPKeeper.ClientKeeper.ClientStore(suite.ctx, exported.Localhost)
}

func TestLocalhostTestSuite(t *testing.T) {
	suite.Run(t, new(LocalhostTestSuite))
}
