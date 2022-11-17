package simulation_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/creatachain/creata-sdk/creataapp"
	"github.com/creatachain/creata-sdk/types/kv"
	"github.com/creatachain/creata-sdk/x/icp/core/02-client/simulation"
	"github.com/creatachain/creata-sdk/x/icp/core/02-client/types"
	host "github.com/creatachain/creata-sdk/x/icp/core/24-host"
	icptmtypes "github.com/creatachain/creata-sdk/x/icp/light-clients/07-augusteum/types"
)

func TestDecodeStore(t *testing.T) {
	app := creataapp.Setup(false)
	clientID := "clientidone"

	height := types.NewHeight(0, 10)

	clientState := &icptmtypes.ClientState{
		FrozenHeight: height,
	}

	consState := &icptmtypes.ConsensusState{
		Timestamp: time.Now().UTC(),
	}

	kvPairs := kv.Pairs{
		Pairs: []kv.Pair{
			{
				Key:   host.FullClientStateKey(clientID),
				Value: app.ICPKeeper.ClientKeeper.MustMarshalClientState(clientState),
			},
			{
				Key:   host.FullConsensusStateKey(clientID, height),
				Value: app.ICPKeeper.ClientKeeper.MustMarshalConsensusState(consState),
			},
			{
				Key:   []byte{0x99},
				Value: []byte{0x99},
			},
		},
	}
	tests := []struct {
		name        string
		expectedLog string
	}{
		{"ClientState", fmt.Sprintf("ClientState A: %v\nClientState B: %v", clientState, clientState)},
		{"ConsensusState", fmt.Sprintf("ConsensusState A: %v\nConsensusState B: %v", consState, consState)},
		{"other", ""},
	}

	for i, tt := range tests {
		i, tt := i, tt
		t.Run(tt.name, func(t *testing.T) {
			res, found := simulation.NewDecodeStore(app.ICPKeeper.ClientKeeper, kvPairs.Pairs[i], kvPairs.Pairs[i])
			if i == len(tests)-1 {
				require.False(t, found, string(kvPairs.Pairs[i].Key))
				require.Empty(t, res, string(kvPairs.Pairs[i].Key))
			} else {
				require.True(t, found, string(kvPairs.Pairs[i].Key))
				require.Equal(t, tt.expectedLog, res, string(kvPairs.Pairs[i].Key))
			}
		})
	}
}
