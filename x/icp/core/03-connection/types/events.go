package types

import (
	"fmt"

	host "github.com/creatachain/creata-sdk/x/icp/core/24-host"
)

// ICP connection events
const (
	AttributeKeyConnectionID             = "connection_id"
	AttributeKeyClientID                 = "client_id"
	AttributeKeyCounterpartyClientID     = "counterparty_client_id"
	AttributeKeyCounterpartyConnectionID = "counterparty_connection_id"
)

// ICP connection events vars
var (
	EventTypeConnectionOpenInit    = MsgConnectionOpenInit{}.Type()
	EventTypeConnectionOpenTry     = MsgConnectionOpenTry{}.Type()
	EventTypeConnectionOpenAck     = MsgConnectionOpenAck{}.Type()
	EventTypeConnectionOpenConfirm = MsgConnectionOpenConfirm{}.Type()

	AttributeValueCategory = fmt.Sprintf("%s_%s", host.ModuleName, SubModuleName)
)
