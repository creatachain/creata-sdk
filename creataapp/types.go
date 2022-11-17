package creataapp

import (
	msm "github.com/creatachain/augusteum/msm/types"

	"github.com/creatachain/creata-sdk/codec"
	"github.com/creatachain/creata-sdk/server/types"
	sdk "github.com/creatachain/creata-sdk/types"
	"github.com/creatachain/creata-sdk/types/module"
)

// App implements the common methods for a Creata SDK-based application
// specific blockchain.
type App interface {
	// The assigned name of the app.
	Name() string

	// The application types codec.
	// NOTE: This shoult be sealed before being returned.
	LegacyAmino() *codec.LegacyAmino

	// Application updates every begin block.
	BeginBlocker(ctx sdk.Context, req msm.RequestBeginBlock) msm.ResponseBeginBlock

	// Application updates every end block.
	EndBlocker(ctx sdk.Context, req msm.RequestEndBlock) msm.ResponseEndBlock

	// Application update at chain (i.e app) initialization.
	InitChainer(ctx sdk.Context, req msm.RequestInitChain) msm.ResponseInitChain

	// Loads the app at a given height.
	LoadHeight(height int64) error

	// Exports the state of the application for a genesis file.
	ExportAppStateAndValidators(
		forZeroHeight bool, jailAllowedAddrs []string,
	) (types.ExportedApp, error)

	// All the registered module account addreses.
	ModuleAccountAddrs() map[string]bool

	// Helper for the simulation framework.
	SimulationManager() *module.SimulationManager
}
