package client

import (
	govclient "github.com/creatachain/creata-sdk/x/gov/client"
	"github.com/creatachain/creata-sdk/x/upgrade/client/cli"
	"github.com/creatachain/creata-sdk/x/upgrade/client/rest"
)

var ProposalHandler = govclient.NewProposalHandler(cli.NewCmdSubmitUpgradeProposal, rest.ProposalRESTHandler)
var CancelProposalHandler = govclient.NewProposalHandler(cli.NewCmdSubmitCancelUpgradeProposal, rest.ProposalCancelRESTHandler)
