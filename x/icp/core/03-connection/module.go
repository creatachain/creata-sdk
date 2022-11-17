package connection

import (
	"github.com/gogo/protobuf/grpc"
	"github.com/spf13/cobra"

	"github.com/creatachain/creata-sdk/x/icp/core/03-connection/client/cli"
	"github.com/creatachain/creata-sdk/x/icp/core/03-connection/types"
)

// Name returns the ICP connection ICS name.
func Name() string {
	return types.SubModuleName
}

// GetTxCmd returns the root tx command for the ICP connections.
func GetTxCmd() *cobra.Command {
	return cli.NewTxCmd()
}

// GetQueryCmd returns the root query command for the ICP connections.
func GetQueryCmd() *cobra.Command {
	return cli.GetQueryCmd()
}

// RegisterQueryService registers the gRPC query service for ICP connections.
func RegisterQueryService(server grpc.Server, queryServer types.QueryServer) {
	types.RegisterQueryServer(server, queryServer)
}
