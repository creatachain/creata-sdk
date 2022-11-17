package types

import (
	"github.com/creatachain/creata-sdk/codec"
	codectypes "github.com/creatachain/creata-sdk/codec/types"
	sdk "github.com/creatachain/creata-sdk/types"
	"github.com/creatachain/creata-sdk/types/msgservice"
	"github.com/creatachain/creata-sdk/x/icp/core/exported"
)

// RegisterInterfaces register the icp interfaces submodule implementations to protobuf
// Any.
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterInterface(
		"icp.core.connection.v1.ConnectionI",
		(*exported.ConnectionI)(nil),
		&ConnectionEnd{},
	)
	registry.RegisterInterface(
		"icp.core.connection.v1.CounterpartyConnectionI",
		(*exported.CounterpartyConnectionI)(nil),
		&Counterparty{},
	)
	registry.RegisterInterface(
		"icp.core.connection.v1.Version",
		(*exported.Version)(nil),
		&Version{},
	)
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgConnectionOpenInit{},
		&MsgConnectionOpenTry{},
		&MsgConnectionOpenAck{},
		&MsgConnectionOpenConfirm{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	// SubModuleCdc references the global x/icp/core/03-connection module codec. Note, the codec should
	// ONLY be used in certain instances of tests and for JSON encoding.
	//
	// The actual codec used for serialization should be provided to x/icp/core/03-connection and
	// defined at the application level.
	SubModuleCdc = codec.NewProtoCodec(codectypes.NewInterfaceRegistry())
)
