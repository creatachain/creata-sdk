package client

import (
	"context"
	"fmt"
	"testing"

	"github.com/creatachain/augusteum/crypto/tmhash"
	"github.com/creatachain/augusteum/mempool"
	"github.com/creatachain/augusteum/rpc/client/mock"
	ctypes "github.com/creatachain/augusteum/rpc/core/types"
	tmtypes "github.com/creatachain/augusteum/types"
	"github.com/stretchr/testify/require"

	"github.com/creatachain/creata-sdk/client/flags"
	sdkerrors "github.com/creatachain/creata-sdk/types/errors"
)

type MockClient struct {
	mock.Client
	err error
}

func (c MockClient) BroadcastTxCommit(ctx context.Context, tx tmtypes.Tx) (*ctypes.ResultBroadcastTxCommit, error) {
	return nil, c.err
}

func (c MockClient) BroadcastTxAsync(ctx context.Context, tx tmtypes.Tx) (*ctypes.ResultBroadcastTx, error) {
	return nil, c.err
}

func (c MockClient) BroadcastTxSync(ctx context.Context, tx tmtypes.Tx) (*ctypes.ResultBroadcastTx, error) {
	return nil, c.err
}

func CreateContextWithErrorAndMode(err error, mode string) Context {
	return Context{
		Client:        MockClient{err: err},
		BroadcastMode: mode,
	}
}

// Test the correct code is returned when
func TestBroadcastError(t *testing.T) {
	errors := map[error]uint32{
		mempool.ErrTxInCache:       sdkerrors.ErrTxInMempoolCache.MSMCode(),
		mempool.ErrTxTooLarge{}:    sdkerrors.ErrTxTooLarge.MSMCode(),
		mempool.ErrMempoolIsFull{}: sdkerrors.ErrMempoolIsFull.MSMCode(),
	}

	modes := []string{
		flags.BroadcastAsync,
		flags.BroadcastBlock,
		flags.BroadcastSync,
	}

	txBytes := []byte{0xA, 0xB}
	txHash := fmt.Sprintf("%X", tmhash.Sum(txBytes))

	for _, mode := range modes {
		for err, code := range errors {
			ctx := CreateContextWithErrorAndMode(err, mode)
			resp, returnedErr := ctx.BroadcastTx(txBytes)
			require.NoError(t, returnedErr)
			require.Equal(t, code, resp.Code)
			require.NotEmpty(t, resp.Codespace)
			require.Equal(t, txHash, resp.TxHash)
		}
	}

}
