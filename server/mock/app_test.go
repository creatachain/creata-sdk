package mock

import (
	"testing"

	msm "github.com/creatachain/augusteum/msm/types"
	tmproto "github.com/creatachain/augusteum/proto/augusteum/types"
	"github.com/creatachain/augusteum/types"
	"github.com/stretchr/testify/require"
)

// TestInitApp makes sure we can initialize this thing without an error
func TestInitApp(t *testing.T) {
	// set up an app
	app, closer, err := SetupApp()

	// closer may need to be run, even when error in later stage
	if closer != nil {
		defer closer()
	}
	require.NoError(t, err)

	// initialize it future-way
	appState, err := AppGenState(nil, types.GenesisDoc{}, nil)
	require.NoError(t, err)

	//TODO test validators in the init chain?
	req := msm.RequestInitChain{
		AppStateBytes: appState,
	}
	app.InitChain(req)
	app.Commit()

	// make sure we can query these values
	query := msm.RequestQuery{
		Path: "/store/main/key",
		Data: []byte("foo"),
	}
	qres := app.Query(query)
	require.Equal(t, uint32(0), qres.Code, qres.Log)
	require.Equal(t, []byte("bar"), qres.Value)
}

// TextDeliverTx ensures we can write a tx
func TestDeliverTx(t *testing.T) {
	// set up an app
	app, closer, err := SetupApp()
	// closer may need to be run, even when error in later stage
	if closer != nil {
		defer closer()
	}
	require.NoError(t, err)

	key := "my-special-key"
	value := "top-secret-data!!"
	tx := NewTx(key, value)
	txBytes := tx.GetSignBytes()

	header := tmproto.Header{
		AppHash: []byte("apphash"),
		Height:  1,
	}
	app.BeginBlock(msm.RequestBeginBlock{Header: header})
	dres := app.DeliverTx(msm.RequestDeliverTx{Tx: txBytes})
	require.Equal(t, uint32(0), dres.Code, dres.Log)
	app.EndBlock(msm.RequestEndBlock{})
	cres := app.Commit()
	require.NotEmpty(t, cres.Data)

	// make sure we can query these values
	query := msm.RequestQuery{
		Path: "/store/main/key",
		Data: []byte(key),
	}
	qres := app.Query(query)
	require.Equal(t, uint32(0), qres.Code, qres.Log)
	require.Equal(t, []byte(value), qres.Value)
}
