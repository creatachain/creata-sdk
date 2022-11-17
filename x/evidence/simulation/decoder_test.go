package simulation_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/creatachain/creata-sdk/creataapp"
	"github.com/creatachain/creata-sdk/crypto/keys/ed25519"
	sdk "github.com/creatachain/creata-sdk/types"
	"github.com/creatachain/creata-sdk/types/kv"
	"github.com/creatachain/creata-sdk/x/evidence/simulation"
	"github.com/creatachain/creata-sdk/x/evidence/types"
)

func TestDecodeStore(t *testing.T) {
	app := creataapp.Setup(false)
	dec := simulation.NewDecodeStore(app.EvidenceKeeper)

	delPk1 := ed25519.GenPrivKey().PubKey()

	ev := &types.Equivocation{
		Height:           10,
		Time:             time.Now().UTC(),
		Power:            1000,
		ConsensusAddress: sdk.ConsAddress(delPk1.Address()).String(),
	}

	evBz, err := app.EvidenceKeeper.MarshalEvidence(ev)
	require.NoError(t, err)

	kvPairs := kv.Pairs{
		Pairs: []kv.Pair{
			{
				Key:   types.KeyPrefixEvidence,
				Value: evBz,
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
		{"Evidence", fmt.Sprintf("%v\n%v", ev, ev)},
		{"other", ""},
	}

	for i, tt := range tests {
		i, tt := i, tt
		t.Run(tt.name, func(t *testing.T) {
			switch i {
			case len(tests) - 1:
				require.Panics(t, func() { dec(kvPairs.Pairs[i], kvPairs.Pairs[i]) }, tt.name)
			default:
				require.Equal(t, tt.expectedLog, dec(kvPairs.Pairs[i], kvPairs.Pairs[i]), tt.name)
			}
		})
	}
}
