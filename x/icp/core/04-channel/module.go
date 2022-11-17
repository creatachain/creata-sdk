package channel

import (
	"github.com/gogo/protobuf/grpc"
	"github.com/spf13/cobra"

	"github.com/creatachain/creata-sdk/x/icp/core/04-channel/client/cli"
	"github.com/creatachain/creata-sdk/x/icp/core/04-channel/types"
)

// Name returns the ICP channel ICS name.
func Name() string {
	return types.SubModuleName
}

// GetTxCmd returns the root tx command for ICP channels.
func GetTxCmd() *cobra.Command {
	return cli.NewTxCmd()
}

// GetQueryCmd returns the root query command for ICP channels.
func GetQueryCmd() *cobra.Command {
	return cli.GetQueryCmd()
}

// RegisterQueryService registers the gRPC query service for ICP channels.
func RegisterQueryService(server grpc.Server, queryServer types.QueryServer) {
	types.RegisterQueryServer(server, queryServer)
}
