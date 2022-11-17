package types

import (
	"github.com/creatachain/creata-sdk/codec"
	codectypes "github.com/creatachain/creata-sdk/codec/types"
	sdk "github.com/creatachain/creata-sdk/types"
	"github.com/creatachain/creata-sdk/types/msgservice"
	"github.com/creatachain/creata-sdk/x/icp/core/exported"
)

// RegisterInterfaces register the icp channel submodule interfaces to protobuf
// Any.
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterInterface(
		"icp.core.channel.v1.ChannelI",
		(*exported.ChannelI)(nil),
	)
	registry.RegisterInterface(
		"icp.core.channel.v1.CounterpartyChannelI",
		(*exported.CounterpartyChannelI)(nil),
	)
	registry.RegisterInterface(
		"icp.core.channel.v1.PacketI",
		(*exported.PacketI)(nil),
	)
	registry.RegisterImplementations(
		(*exported.ChannelI)(nil),
		&Channel{},
	)
	registry.RegisterImplementations(
		(*exported.CounterpartyChannelI)(nil),
		&Counterparty{},
	)
	registry.RegisterImplementations(
		(*exported.PacketI)(nil),
		&Packet{},
	)
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgChannelOpenInit{},
		&MsgChannelOpenTry{},
		&MsgChannelOpenAck{},
		&MsgChannelOpenConfirm{},
		&MsgChannelCloseInit{},
		&MsgChannelCloseConfirm{},
		&MsgRecvPacket{},
		&MsgAcknowledgement{},
		&MsgTimeout{},
		&MsgTimeoutOnClose{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

// SubModuleCdc references the global x/icp/core/04-channel module codec. Note, the codec should
// ONLY be used in certain instances of tests and for JSON encoding.
//
// The actual codec used for serialization should be provided to x/icp/core/04-channel and
// defined at the application level.
var SubModuleCdc = codec.NewProtoCodec(codectypes.NewInterfaceRegistry())
