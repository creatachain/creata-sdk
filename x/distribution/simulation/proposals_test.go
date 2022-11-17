package simulation_test

import (
	"math/rand"
	"testing"

	tmproto "github.com/creatachain/augusteum/proto/augusteum/types"
	"github.com/stretchr/testify/require"

	"github.com/creatachain/creata-sdk/creataapp"
	creataappparams "github.com/creatachain/creata-sdk/creataapp/params"
	sdk "github.com/creatachain/creata-sdk/types"
	simtypes "github.com/creatachain/creata-sdk/types/simulation"
	"github.com/creatachain/creata-sdk/x/distribution/simulation"
)

func TestProposalContents(t *testing.T) {
	app := creataapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	// initialize parameters
	s := rand.NewSource(1)
	r := rand.New(s)

	accounts := simtypes.RandomAccounts(r, 3)

	// execute ProposalContents function
	weightedProposalContent := simulation.ProposalContents(app.DistrKeeper)
	require.Len(t, weightedProposalContent, 1)

	w0 := weightedProposalContent[0]

	// tests w0 interface:
	require.Equal(t, simulation.OpWeightSubmitCommunitySpendProposal, w0.AppParamsKey())
	require.Equal(t, creataappparams.DefaultWeightTextProposal, w0.DefaultWeight())

	amount := sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(1)), sdk.NewCoin("atoken", sdk.NewInt(2)))

	feePool := app.DistrKeeper.GetFeePool(ctx)
	feePool.CommunityPool = sdk.NewDecCoinsFromCoins(amount...)
	app.DistrKeeper.SetFeePool(ctx, feePool)

	content := w0.ContentSimulatorFn()(r, ctx, accounts)

	require.Equal(t, "sTxPjfweXhSUkMhPjMaxKlMIJMOXcnQfyzeOcbWwNbeHVIkPZBSpYuLyYggwexjxusrBqDOTtGTOWeLrQKjLxzIivHSlcxgdXhhu", content.GetDescription())
	require.Equal(t, "xKGLwQvuyN", content.GetTitle())
	require.Equal(t, "distribution", content.ProposalRoute())
	require.Equal(t, "CommunityPoolSpend", content.ProposalType())
}
