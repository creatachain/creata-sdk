package tx

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/creatachain/creata-sdk/codec"
	codectypes "github.com/creatachain/creata-sdk/codec/types"
	"github.com/creatachain/creata-sdk/std"
	"github.com/creatachain/creata-sdk/testutil/testdata"
	sdk "github.com/creatachain/creata-sdk/types"
	"github.com/creatachain/creata-sdk/x/auth/testutil"
)

func TestGenerator(t *testing.T) {
	interfaceRegistry := codectypes.NewInterfaceRegistry()
	std.RegisterInterfaces(interfaceRegistry)
	interfaceRegistry.RegisterImplementations((*sdk.Msg)(nil), &testdata.TestMsg{})
	protoCodec := codec.NewProtoCodec(interfaceRegistry)
	suite.Run(t, testutil.NewTxConfigTestSuite(NewTxConfig(protoCodec, DefaultSignModes)))
}
