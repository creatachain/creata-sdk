package types

import (
	codectypes "github.com/creatachain/creata-sdk/codec/types"
	clienttypes "github.com/creatachain/creata-sdk/x/icp/core/02-client/types"
	connectiontypes "github.com/creatachain/creata-sdk/x/icp/core/03-connection/types"
	channeltypes "github.com/creatachain/creata-sdk/x/icp/core/04-channel/types"
	commitmenttypes "github.com/creatachain/creata-sdk/x/icp/core/23-commitment/types"
	solomachinetypes "github.com/creatachain/creata-sdk/x/icp/light-clients/06-solomachine/types"
	icptmtypes "github.com/creatachain/creata-sdk/x/icp/light-clients/07-augusteum/types"
	localhosttypes "github.com/creatachain/creata-sdk/x/icp/light-clients/09-localhost/types"
)

// RegisterInterfaces registers x/icp interfaces into protobuf Any.
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	clienttypes.RegisterInterfaces(registry)
	connectiontypes.RegisterInterfaces(registry)
	channeltypes.RegisterInterfaces(registry)
	solomachinetypes.RegisterInterfaces(registry)
	icptmtypes.RegisterInterfaces(registry)
	localhosttypes.RegisterInterfaces(registry)
	commitmenttypes.RegisterInterfaces(registry)
}
