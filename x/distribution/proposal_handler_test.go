package distribution_test

import (
	"testing"

	tmproto "github.com/creatachain/augusteum/proto/augusteum/types"
	"github.com/stretchr/testify/require"

	"github.com/creatachain/creata-sdk/creataapp"
	"github.com/creatachain/creata-sdk/crypto/keys/ed25519"
	sdk "github.com/creatachain/creata-sdk/types"
	"github.com/creatachain/creata-sdk/x/distribution"
	"github.com/creatachain/creata-sdk/x/distribution/types"
)

var (
	delPk1   = ed25519.GenPrivKey().PubKey()
	delAddr1 = sdk.AccAddress(delPk1.Address())

	amount = sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(1)))
)

func testProposal(recipient sdk.AccAddress, amount sdk.Coins) *types.CommunityPoolSpendProposal {
	return types.NewCommunityPoolSpendProposal("Test", "description", recipient, amount)
}

func TestProposalHandlerPassed(t *testing.T) {
	app := creataapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	recipient := delAddr1

	// add coins to the module account
	macc := app.DistrKeeper.GetDistributionAccount(ctx)
	balances := app.BankKeeper.GetAllBalances(ctx, macc.GetAddress())
	err := app.BankKeeper.SetBalances(ctx, macc.GetAddress(), balances.Add(amount...))
	require.NoError(t, err)

	app.AccountKeeper.SetModuleAccount(ctx, macc)

	account := app.AccountKeeper.NewAccountWithAddress(ctx, recipient)
	app.AccountKeeper.SetAccount(ctx, account)
	require.True(t, app.BankKeeper.GetAllBalances(ctx, account.GetAddress()).IsZero())

	feePool := app.DistrKeeper.GetFeePool(ctx)
	feePool.CommunityPool = sdk.NewDecCoinsFromCoins(amount...)
	app.DistrKeeper.SetFeePool(ctx, feePool)

	tp := testProposal(recipient, amount)
	hdlr := distribution.NewCommunityPoolSpendProposalHandler(app.DistrKeeper)
	require.NoError(t, hdlr(ctx, tp))

	balances = app.BankKeeper.GetAllBalances(ctx, recipient)
	require.Equal(t, balances, amount)
}

func TestProposalHandlerFailed(t *testing.T) {
	app := creataapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	recipient := delAddr1

	account := app.AccountKeeper.NewAccountWithAddress(ctx, recipient)
	app.AccountKeeper.SetAccount(ctx, account)
	require.True(t, app.BankKeeper.GetAllBalances(ctx, account.GetAddress()).IsZero())

	tp := testProposal(recipient, amount)
	hdlr := distribution.NewCommunityPoolSpendProposalHandler(app.DistrKeeper)
	require.Error(t, hdlr(ctx, tp))

	balances := app.BankKeeper.GetAllBalances(ctx, recipient)
	require.True(t, balances.IsZero())
}
