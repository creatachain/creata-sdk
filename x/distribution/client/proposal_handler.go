package client

import (
	"github.com/creatachain/creata-sdk/x/distribution/client/cli"
	"github.com/creatachain/creata-sdk/x/distribution/client/rest"
	govclient "github.com/creatachain/creata-sdk/x/gov/client"
)

// ProposalHandler is the community spend proposal handler.
var (
	ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitProposal, rest.ProposalRESTHandler)
)
