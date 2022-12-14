package client

import (
	"context"
	"fmt"
	"strings"

	"github.com/pkg/errors"

	tmbytes "github.com/creatachain/augusteum/libs/bytes"
	msm "github.com/creatachain/augusteum/msm/types"
	rpcclient "github.com/creatachain/augusteum/rpc/client"

	"github.com/creatachain/creata-sdk/store/rootmulti"
	sdk "github.com/creatachain/creata-sdk/types"
)

// GetNode returns an RPC client. If the context's client is not defined, an
// error is returned.
func (ctx Context) GetNode() (rpcclient.Client, error) {
	if ctx.Client == nil {
		return nil, errors.New("no RPC client is defined in offline mode")
	}

	return ctx.Client, nil
}

// Query performs a query to a Augusteum node with the provided path.
// It returns the result and height of the query upon success or an error if
// the query fails.
func (ctx Context) Query(path string) ([]byte, int64, error) {
	return ctx.query(path, nil)
}

// QueryWithData performs a query to a Augusteum node with the provided path
// and a data payload. It returns the result and height of the query upon success
// or an error if the query fails.
func (ctx Context) QueryWithData(path string, data []byte) ([]byte, int64, error) {
	return ctx.query(path, data)
}

// QueryStore performs a query to a Augusteum node with the provided key and
// store name. It returns the result and height of the query upon success
// or an error if the query fails.
func (ctx Context) QueryStore(key tmbytes.HexBytes, storeName string) ([]byte, int64, error) {
	return ctx.queryStore(key, storeName, "key")
}

// QueryMSM performs a query to a Augusteum node with the provide RequestQuery.
// It returns the ResultQuery obtained from the query.
func (ctx Context) QueryMSM(req msm.RequestQuery) (msm.ResponseQuery, error) {
	return ctx.queryMSM(req)
}

// GetFromAddress returns the from address from the context's name.
func (ctx Context) GetFromAddress() sdk.AccAddress {
	return ctx.FromAddress
}

// GetFromName returns the key name for the current context.
func (ctx Context) GetFromName() string {
	return ctx.FromName
}

func (ctx Context) queryMSM(req msm.RequestQuery) (msm.ResponseQuery, error) {
	node, err := ctx.GetNode()
	if err != nil {
		return msm.ResponseQuery{}, err
	}

	opts := rpcclient.MSMQueryOptions{
		Height: ctx.Height,
		Prove:  req.Prove,
	}

	result, err := node.MSMQueryWithOptions(context.Background(), req.Path, req.Data, opts)
	if err != nil {
		return msm.ResponseQuery{}, err
	}

	if !result.Response.IsOK() {
		return msm.ResponseQuery{}, errors.New(result.Response.Log)
	}

	// data from trusted node or subspace query doesn't need verification
	if !opts.Prove || !isQueryStoreWithProof(req.Path) {
		return result.Response, nil
	}

	return result.Response, nil
}

// query performs a query to a Augusteum node with the provided store name
// and path. It returns the result and height of the query upon success
// or an error if the query fails.
func (ctx Context) query(path string, key tmbytes.HexBytes) ([]byte, int64, error) {
	resp, err := ctx.queryMSM(msm.RequestQuery{
		Path: path,
		Data: key,
	})
	if err != nil {
		return nil, 0, err
	}

	return resp.Value, resp.Height, nil
}

// queryStore performs a query to a Augusteum node with the provided a store
// name and path. It returns the result and height of the query upon success
// or an error if the query fails.
func (ctx Context) queryStore(key tmbytes.HexBytes, storeName, endPath string) ([]byte, int64, error) {
	path := fmt.Sprintf("/store/%s/%s", storeName, endPath)
	return ctx.query(path, key)
}

// isQueryStoreWithProof expects a format like /<queryType>/<storeName>/<subpath>
// queryType must be "store" and subpath must be "key" to require a proof.
func isQueryStoreWithProof(path string) bool {
	if !strings.HasPrefix(path, "/") {
		return false
	}

	paths := strings.SplitN(path[1:], "/", 3)

	switch {
	case len(paths) != 3:
		return false
	case paths[0] != "store":
		return false
	case rootmulti.RequireProof("/" + paths[2]):
		return true
	}

	return false
}
