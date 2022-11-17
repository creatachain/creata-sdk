package simulation_test

import (
	"encoding/json"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/creatachain/creata-sdk/codec"
	codectypes "github.com/creatachain/creata-sdk/codec/types"
	"github.com/creatachain/creata-sdk/types/module"
	simtypes "github.com/creatachain/creata-sdk/types/simulation"
	host "github.com/creatachain/creata-sdk/x/icp/core/24-host"
	"github.com/creatachain/creata-sdk/x/icp/core/simulation"
	"github.com/creatachain/creata-sdk/x/icp/core/types"
)

// TestRandomizedGenState tests the normal scenario of applying RandomizedGenState.
// Abonormal scenarios are not tested here.
func TestRandomizedGenState(t *testing.T) {
	interfaceRegistry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(interfaceRegistry)

	s := rand.NewSource(1)
	r := rand.New(s)

	simState := module.SimulationState{
		AppParams:    make(simtypes.AppParams),
		Cdc:          cdc,
		Rand:         r,
		NumBonded:    3,
		Accounts:     simtypes.RandomAccounts(r, 3),
		InitialStake: 1000,
		GenState:     make(map[string]json.RawMessage),
	}

	// Remark: the current RandomizedGenState function
	// is actually not random as it does not utilize concretely the random value r.
	// This tests will pass for any value of r.
	simulation.RandomizedGenState(&simState)

	var icpGenesis types.GenesisState
	simState.Cdc.MustUnmarshalJSON(simState.GenState[host.ModuleName], &icpGenesis)

	require.NotNil(t, icpGenesis.ClientGenesis)
	require.NotNil(t, icpGenesis.ConnectionGenesis)
	require.NotNil(t, icpGenesis.ChannelGenesis)
}
