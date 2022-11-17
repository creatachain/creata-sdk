package v040_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/creatachain/creata-sdk/client"
	"github.com/creatachain/creata-sdk/creataapp"
	sdk "github.com/creatachain/creata-sdk/types"
	v038evidence "github.com/creatachain/creata-sdk/x/evidence/legacy/v038"
	v040evidence "github.com/creatachain/creata-sdk/x/evidence/legacy/v040"
)

func TestMigrate(t *testing.T) {
	encodingConfig := creataapp.MakeTestEncodingConfig()
	clientCtx := client.Context{}.
		WithInterfaceRegistry(encodingConfig.InterfaceRegistry).
		WithTxConfig(encodingConfig.TxConfig).
		WithLegacyAmino(encodingConfig.Amino).
		WithJSONMarshaler(encodingConfig.Marshaler)

	addr1, _ := sdk.AccAddressFromBech32("creata1xxkueklal9vejv9unqu80w9vptyepfa95pd53u")

	evidenceGenState := v038evidence.GenesisState{
		Params: v038evidence.Params{MaxEvidenceAge: v038evidence.DefaultMaxEvidenceAge},
		Evidence: []v038evidence.Evidence{v038evidence.Equivocation{
			Height:           20,
			Power:            100,
			ConsensusAddress: addr1.Bytes(),
		}},
	}

	migrated := v040evidence.Migrate(evidenceGenState)
	expected := `{"evidence":[{"@type":"/creata.evidence.v1beta1.Equivocation","height":"20","time":"0001-01-01T00:00:00Z","power":"100","consensus_address":"creatavalcons1xxkueklal9vejv9unqu80w9vptyepfa99x2a3w"}]}`

	bz, err := clientCtx.JSONMarshaler.MarshalJSON(migrated)
	require.NoError(t, err)
	require.Equal(t, expected, string(bz))
}
