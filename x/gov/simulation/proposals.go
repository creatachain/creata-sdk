package simulation

import (
	"math/rand"

	creataappparams "github.com/creatachain/creata-sdk/creataapp/params"
	sdk "github.com/creatachain/creata-sdk/types"
	simtypes "github.com/creatachain/creata-sdk/types/simulation"
	"github.com/creatachain/creata-sdk/x/gov/types"
	"github.com/creatachain/creata-sdk/x/simulation"
)

// OpWeightSubmitTextProposal app params key for text proposal
const OpWeightSubmitTextProposal = "op_weight_submit_text_proposal"

// ProposalContents defines the module weighted proposals' contents
func ProposalContents() []simtypes.WeightedProposalContent {
	return []simtypes.WeightedProposalContent{
		simulation.NewWeightedProposalContent(
			OpWeightMsgDeposit,
			creataappparams.DefaultWeightTextProposal,
			SimulateTextProposalContent,
		),
	}
}

// SimulateTextProposalContent returns a random text proposal content.
func SimulateTextProposalContent(r *rand.Rand, _ sdk.Context, _ []simtypes.Account) simtypes.Content {
	return types.NewTextProposal(
		simtypes.RandStringOfLength(r, 140),
		simtypes.RandStringOfLength(r, 5000),
	)
}
