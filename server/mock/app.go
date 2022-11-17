package mock

import (
	"encoding/json"
	"errors"
	"fmt"
	"path/filepath"

	"github.com/creatachain/augusteum/types"

	"github.com/creatachain/augusteum/libs/log"
	msm "github.com/creatachain/augusteum/msm/types"

	bam "github.com/creatachain/creata-sdk/baseapp"
	"github.com/creatachain/creata-sdk/codec"
	sdk "github.com/creatachain/creata-sdk/types"
)

// NewApp creates a simple mock kvstore app for testing. It should work
// similar to a real app. Make sure rootDir is empty before running the test,
// in order to guarantee consistent results
func NewApp(rootDir string, logger log.Logger) (msm.Application, error) {
	db, err := sdk.NewLevelDB("mock", filepath.Join(rootDir, "data"))
	if err != nil {
		return nil, err
	}

	// Capabilities key to access the main KVStore.
	capKeyMainStore := sdk.NewKVStoreKey("main")

	// Create BaseApp.
	baseApp := bam.NewBaseApp("kvstore", logger, db, decodeTx)

	// Set mounts for BaseApp's MultiStore.
	baseApp.MountStores(capKeyMainStore)

	baseApp.SetInitChainer(InitChainer(capKeyMainStore))

	// Set a Route.
	baseApp.Router().AddRoute(sdk.NewRoute("kvstore", KVStoreHandler(capKeyMainStore)))

	// Load latest version.
	if err := baseApp.LoadLatestVersion(); err != nil {
		return nil, err
	}

	return baseApp, nil
}

// KVStoreHandler is a simple handler that takes kvstoreTx and writes
// them to the db
func KVStoreHandler(storeKey sdk.StoreKey) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		dTx, ok := msg.(kvstoreTx)
		if !ok {
			return nil, errors.New("KVStoreHandler should only receive kvstoreTx")
		}

		// tx is already unmarshalled
		key := dTx.key
		value := dTx.value

		store := ctx.KVStore(storeKey)
		store.Set(key, value)

		return &sdk.Result{
			Log: fmt.Sprintf("set %s=%s", key, value),
		}, nil
	}
}

// basic KV structure
type KV struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// What Genesis JSON is formatted as
type GenesisJSON struct {
	Values []KV `json:"values"`
}

// InitChainer returns a function that can initialize the chain
// with key/value pairs
func InitChainer(key sdk.StoreKey) func(sdk.Context, msm.RequestInitChain) msm.ResponseInitChain {
	return func(ctx sdk.Context, req msm.RequestInitChain) msm.ResponseInitChain {
		stateJSON := req.AppStateBytes

		genesisState := new(GenesisJSON)
		err := json.Unmarshal(stateJSON, genesisState)
		if err != nil {
			panic(err)
		}

		for _, val := range genesisState.Values {
			store := ctx.KVStore(key)
			store.Set([]byte(val.Key), []byte(val.Value))
		}
		return msm.ResponseInitChain{}
	}
}

// AppGenState can be passed into InitCmd, returns a static string of a few
// key-values that can be parsed by InitChainer
func AppGenState(_ *codec.LegacyAmino, _ types.GenesisDoc, _ []json.RawMessage) (appState json.
	RawMessage, err error) {
	appState = json.RawMessage(`{
  "values": [
    {
        "key": "hello",
        "value": "goodbye"
    },
    {
        "key": "foo",
        "value": "bar"
    }
  ]
}`)
	return
}

// AppGenStateEmpty returns an empty transaction state for mocking.
func AppGenStateEmpty(_ *codec.LegacyAmino, _ types.GenesisDoc, _ []json.RawMessage) (
	appState json.RawMessage, err error) {
	appState = json.RawMessage(``)
	return
}
