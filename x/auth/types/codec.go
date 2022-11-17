package types

import (
	"github.com/creatachain/creata-sdk/codec"
	"github.com/creatachain/creata-sdk/codec/types"
	cryptocodec "github.com/creatachain/creata-sdk/crypto/codec"
	"github.com/creatachain/creata-sdk/x/auth/legacy/legacytx"
)

// RegisterLegacyAminoCodec registers the account interfaces and concrete types on the
// provided LegacyAmino codec. These types are used for Amino JSON serialization
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterInterface((*ModuleAccountI)(nil), nil)
	cdc.RegisterInterface((*GenesisAccount)(nil), nil)
	cdc.RegisterInterface((*AccountI)(nil), nil)
	cdc.RegisterConcrete(&BaseAccount{}, "creata-sdk/BaseAccount", nil)
	cdc.RegisterConcrete(&ModuleAccount{}, "creata-sdk/ModuleAccount", nil)

	legacytx.RegisterLegacyAminoCodec(cdc)
}

// RegisterInterfaces associates protoName with AccountI interface
// and creates a registry of it's concrete implementations
func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterInterface(
		"creata.auth.v1beta1.AccountI",
		(*AccountI)(nil),
		&BaseAccount{},
		&ModuleAccount{},
	)

	registry.RegisterInterface(
		"creata.auth.v1beta1.GenesisAccount",
		(*GenesisAccount)(nil),
		&BaseAccount{},
		&ModuleAccount{},
	)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
}
