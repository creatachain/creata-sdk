package types

import (
	"fmt"

	host "github.com/creatachain/creata-sdk/x/icp/core/24-host"
)

// ICP client events
const (
	AttributeKeyClientID        = "client_id"
	AttributeKeyClientType      = "client_type"
	AttributeKeyConsensusHeight = "consensus_height"
	AttributeKeyHeader          = "header"
)

// ICP client events vars
var (
	EventTypeCreateClient         = "create_client"
	EventTypeUpdateClient         = "update_client"
	EventTypeUpgradeClient        = "upgrade_client"
	EventTypeSubmitMisbehaviour   = "client_misbehaviour"
	EventTypeUpdateClientProposal = "update_client_proposal"

	AttributeValueCategory = fmt.Sprintf("%s_%s", host.ModuleName, SubModuleName)
)
