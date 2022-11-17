package legacytx_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/creatachain/creata-sdk/codec"
	cryptoAmino "github.com/creatachain/creata-sdk/crypto/codec"
	"github.com/creatachain/creata-sdk/testutil/testdata"
	sdk "github.com/creatachain/creata-sdk/types"
	"github.com/creatachain/creata-sdk/x/auth/legacy/legacytx"
	"github.com/creatachain/creata-sdk/x/auth/testutil"
)

func testCodec() *codec.LegacyAmino {
	cdc := codec.NewLegacyAmino()
	sdk.RegisterLegacyAminoCodec(cdc)
	cryptoAmino.RegisterCrypto(cdc)
	cdc.RegisterConcrete(&testdata.TestMsg{}, "creata-sdk/Test", nil)
	return cdc
}

func TestStdTxConfig(t *testing.T) {
	cdc := testCodec()
	txGen := legacytx.StdTxConfig{Cdc: cdc}
	suite.Run(t, testutil.NewTxConfigTestSuite(txGen))
}
