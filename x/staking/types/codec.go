package types

import (
	"github.com/creatachain/creata-sdk/codec"
	"github.com/creatachain/creata-sdk/codec/types"
	cryptocodec "github.com/creatachain/creata-sdk/crypto/codec"
	sdk "github.com/creatachain/creata-sdk/types"
	"github.com/creatachain/creata-sdk/types/msgservice"
)

// RegisterLegacyAminoCodec registers the necessary x/staking interfaces and concrete types
// on the provided LegacyAmino codec. These types are used for Amino JSON serialization.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateValidator{}, "creata-sdk/MsgCreateValidator", nil)
	cdc.RegisterConcrete(&MsgEditValidator{}, "creata-sdk/MsgEditValidator", nil)
	cdc.RegisterConcrete(&MsgDelegate{}, "creata-sdk/MsgDelegate", nil)
	cdc.RegisterConcrete(&MsgUndelegate{}, "creata-sdk/MsgUndelegate", nil)
	cdc.RegisterConcrete(&MsgBeginRedelegate{}, "creata-sdk/MsgBeginRedelegate", nil)
}

// RegisterInterfaces registers the x/staking interfaces types with the interface registry
func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateValidator{},
		&MsgEditValidator{},
		&MsgDelegate{},
		&MsgUndelegate{},
		&MsgBeginRedelegate{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino = codec.NewLegacyAmino()

	// ModuleCdc references the global x/staking module codec. Note, the codec should
	// ONLY be used in certain instances of tests and for JSON encoding as Amino is
	// still used for that purpose.
	//
	// The actual codec used for serialization should be provided to x/staking and
	// defined at the application level.
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}
