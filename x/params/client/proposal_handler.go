package client

import (
	govclient "github.com/creatachain/creata-sdk/x/gov/client"
	"github.com/creatachain/creata-sdk/x/params/client/cli"
	"github.com/creatachain/creata-sdk/x/params/client/rest"
)

// ProposalHandler is the param change proposal handler.
var ProposalHandler = govclient.NewProposalHandler(cli.NewSubmitParamChangeProposalTxCmd, rest.ProposalRESTHandler)
