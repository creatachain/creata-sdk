package simulation

import (
	"math/rand"

	simtypes "github.com/creatachain/creata-sdk/types/simulation"
	"github.com/creatachain/creata-sdk/x/icp/core/03-connection/types"
)

// GenConnectionGenesis returns the default connection genesis state.
func GenConnectionGenesis(_ *rand.Rand, _ []simtypes.Account) types.GenesisState {
	return types.DefaultGenesisState()
}
