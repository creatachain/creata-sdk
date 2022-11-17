package types

import (
	"github.com/gogo/protobuf/grpc"

	client "github.com/creatachain/creata-sdk/x/icp/core/02-client"
	clienttypes "github.com/creatachain/creata-sdk/x/icp/core/02-client/types"
	connection "github.com/creatachain/creata-sdk/x/icp/core/03-connection"
	connectiontypes "github.com/creatachain/creata-sdk/x/icp/core/03-connection/types"
	channel "github.com/creatachain/creata-sdk/x/icp/core/04-channel"
	channeltypes "github.com/creatachain/creata-sdk/x/icp/core/04-channel/types"
)

// QueryServer defines the ICP interfaces that the gRPC query server must implement
type QueryServer interface {
	clienttypes.QueryServer
	connectiontypes.QueryServer
	channeltypes.QueryServer
}

// RegisterQueryService registers each individual ICP submodule query service
func RegisterQueryService(server grpc.Server, queryService QueryServer) {
	client.RegisterQueryService(server, queryService)
	connection.RegisterQueryService(server, queryService)
	channel.RegisterQueryService(server, queryService)
}
