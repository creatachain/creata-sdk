package types

import (
	"github.com/creatachain/creata-sdk/codec"
	"github.com/creatachain/creata-sdk/codec/types"
	sdk "github.com/creatachain/creata-sdk/types"
	"github.com/creatachain/creata-sdk/types/msgservice"
	authtypes "github.com/creatachain/creata-sdk/x/auth/types"
	"github.com/creatachain/creata-sdk/x/auth/vesting/exported"
)

// RegisterLegacyAminoCodec registers the vesting interfaces and concrete types on the
// provided LegacyAmino codec. These types are used for Amino JSON serialization
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterInterface((*exported.VestingAccount)(nil), nil)
	cdc.RegisterConcrete(&BaseVestingAccount{}, "creata-sdk/BaseVestingAccount", nil)
	cdc.RegisterConcrete(&ContinuousVestingAccount{}, "creata-sdk/ContinuousVestingAccount", nil)
	cdc.RegisterConcrete(&DelayedVestingAccount{}, "creata-sdk/DelayedVestingAccount", nil)
	cdc.RegisterConcrete(&PeriodicVestingAccount{}, "creata-sdk/PeriodicVestingAccount", nil)
}

// RegisterInterface associates protoName with AccountI and VestingAccount
// Interfaces and creates a registry of it's concrete implementations
func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterInterface(
		"creata.vesting.v1beta1.VestingAccount",
		(*exported.VestingAccount)(nil),
		&ContinuousVestingAccount{},
		&DelayedVestingAccount{},
		&PeriodicVestingAccount{},
	)

	registry.RegisterImplementations(
		(*authtypes.AccountI)(nil),
		&BaseVestingAccount{},
		&DelayedVestingAccount{},
		&ContinuousVestingAccount{},
		&PeriodicVestingAccount{},
	)

	registry.RegisterImplementations(
		(*authtypes.GenesisAccount)(nil),
		&BaseVestingAccount{},
		&DelayedVestingAccount{},
		&ContinuousVestingAccount{},
		&PeriodicVestingAccount{},
	)

	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgCreateVestingAccount{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var amino = codec.NewLegacyAmino()

func init() {
	RegisterLegacyAminoCodec(amino)
	amino.Seal()
}
