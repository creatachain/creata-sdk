package icp

import (
	sdk "github.com/creatachain/creata-sdk/types"
	client "github.com/creatachain/creata-sdk/x/icp/core/02-client"
	connection "github.com/creatachain/creata-sdk/x/icp/core/03-connection"
	channel "github.com/creatachain/creata-sdk/x/icp/core/04-channel"
	"github.com/creatachain/creata-sdk/x/icp/core/keeper"
	"github.com/creatachain/creata-sdk/x/icp/core/types"
)

// InitGenesis initializes the icp state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, createLocalhost bool, gs *types.GenesisState) {
	client.InitGenesis(ctx, k.ClientKeeper, gs.ClientGenesis)
	connection.InitGenesis(ctx, k.ConnectionKeeper, gs.ConnectionGenesis)
	channel.InitGenesis(ctx, k.ChannelKeeper, gs.ChannelGenesis)
}

// ExportGenesis returns the icp exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	return &types.GenesisState{
		ClientGenesis:     client.ExportGenesis(ctx, k.ClientKeeper),
		ConnectionGenesis: connection.ExportGenesis(ctx, k.ConnectionKeeper),
		ChannelGenesis:    channel.ExportGenesis(ctx, k.ChannelKeeper),
	}
}
