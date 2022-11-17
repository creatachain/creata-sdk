package simulation

import (
	"math/rand"

	simtypes "github.com/creatachain/creata-sdk/types/simulation"
	"github.com/creatachain/creata-sdk/x/icp/core/02-client/types"
)

// GenClientGenesis returns the default client genesis state.
func GenClientGenesis(_ *rand.Rand, _ []simtypes.Account) types.GenesisState {
	return types.DefaultGenesisState()
}
