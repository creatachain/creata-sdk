package client

import (
	"github.com/gogo/protobuf/grpc"
	"github.com/spf13/cobra"

	"github.com/creatachain/creata-sdk/x/icp/core/02-client/client/cli"
	"github.com/creatachain/creata-sdk/x/icp/core/02-client/types"
)

// Name returns the ICP client name
func Name() string {
	return types.SubModuleName
}

// GetQueryCmd returns no root query command for the ICP client
func GetQueryCmd() *cobra.Command {
	return cli.GetQueryCmd()
}

// RegisterQueryService registers the gRPC query service for ICP client.
func RegisterQueryService(server grpc.Server, queryServer types.QueryServer) {
	types.RegisterQueryServer(server, queryServer)
}
