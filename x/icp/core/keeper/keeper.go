package keeper

import (
	"github.com/creatachain/creata-sdk/codec"
	sdk "github.com/creatachain/creata-sdk/types"
	capabilitykeeper "github.com/creatachain/creata-sdk/x/capability/keeper"
	clientkeeper "github.com/creatachain/creata-sdk/x/icp/core/02-client/keeper"
	clienttypes "github.com/creatachain/creata-sdk/x/icp/core/02-client/types"
	connectionkeeper "github.com/creatachain/creata-sdk/x/icp/core/03-connection/keeper"
	channelkeeper "github.com/creatachain/creata-sdk/x/icp/core/04-channel/keeper"
	portkeeper "github.com/creatachain/creata-sdk/x/icp/core/05-port/keeper"
	porttypes "github.com/creatachain/creata-sdk/x/icp/core/05-port/types"
	"github.com/creatachain/creata-sdk/x/icp/core/types"
	paramtypes "github.com/creatachain/creata-sdk/x/params/types"
)

var _ types.QueryServer = (*Keeper)(nil)

// Keeper defines each ICS keeper for ICP
type Keeper struct {
	// implements gRPC QueryServer interface
	types.QueryServer

	cdc codec.BinaryMarshaler

	ClientKeeper     clientkeeper.Keeper
	ConnectionKeeper connectionkeeper.Keeper
	ChannelKeeper    channelkeeper.Keeper
	PortKeeper       portkeeper.Keeper
	Router           *porttypes.Router
}

// NewKeeper creates a new icp Keeper
func NewKeeper(
	cdc codec.BinaryMarshaler, key sdk.StoreKey, paramSpace paramtypes.Subspace,
	stakingKeeper clienttypes.StakingKeeper, scopedKeeper capabilitykeeper.ScopedKeeper,
) *Keeper {
	clientKeeper := clientkeeper.NewKeeper(cdc, key, paramSpace, stakingKeeper)
	connectionKeeper := connectionkeeper.NewKeeper(cdc, key, clientKeeper)
	portKeeper := portkeeper.NewKeeper(scopedKeeper)
	channelKeeper := channelkeeper.NewKeeper(cdc, key, clientKeeper, connectionKeeper, portKeeper, scopedKeeper)

	return &Keeper{
		cdc:              cdc,
		ClientKeeper:     clientKeeper,
		ConnectionKeeper: connectionKeeper,
		ChannelKeeper:    channelKeeper,
		PortKeeper:       portKeeper,
	}
}

// Codec returns the ICP module codec.
func (k Keeper) Codec() codec.BinaryMarshaler {
	return k.cdc
}

// SetRouter sets the Router in ICP Keeper and seals it. The method panics if
// there is an existing router that's already sealed.
func (k *Keeper) SetRouter(rtr *porttypes.Router) {
	if k.Router != nil && k.Router.Sealed() {
		panic("cannot reset a sealed router")
	}
	k.Router = rtr
	k.Router.Seal()
}
