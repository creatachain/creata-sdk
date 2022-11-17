package tmservice

import (
	"context"

	ctypes "github.com/creatachain/augusteum/rpc/core/types"

	"github.com/creatachain/creata-sdk/client"
)

func getNodeStatus(clientCtx client.Context) (*ctypes.ResultStatus, error) {
	node, err := clientCtx.GetNode()
	if err != nil {
		return &ctypes.ResultStatus{}, err
	}
	return node.Status(context.Background())
}
