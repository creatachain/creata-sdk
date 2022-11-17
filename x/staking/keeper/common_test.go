package keeper_test

import (
	"math/big"
	"testing"

	tmproto "github.com/creatachain/augusteum/proto/augusteum/types"

	"github.com/creatachain/creata-sdk/codec"
	"github.com/creatachain/creata-sdk/creataapp"
	sdk "github.com/creatachain/creata-sdk/types"
	"github.com/creatachain/creata-sdk/x/staking/keeper"
	"github.com/creatachain/creata-sdk/x/staking/types"
)

var (
	PKs = creataapp.CreateTestPubKeys(500)
)

func init() {
	sdk.PowerReduction = sdk.NewIntFromBigInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil))
}

// createTestInput Returns a creataapp with custom StakingKeeper
// to avoid messing with the hooks.
func createTestInput() (*codec.LegacyAmino, *creataapp.CreataApp, sdk.Context) {
	app := creataapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	app.StakingKeeper = keeper.NewKeeper(
		app.AppCodec(),
		app.GetKey(types.StoreKey),
		app.AccountKeeper,
		app.BankKeeper,
		app.GetSubspace(types.ModuleName),
	)
	return app.LegacyAmino(), app, ctx
}

// intended to be used with require/assert:  require.True(ValEq(...))
func ValEq(t *testing.T, exp, got types.Validator) (*testing.T, bool, string, types.Validator, types.Validator) {
	return t, exp.MinEqual(&got), "expected:\n%v\ngot:\n%v", exp, got
}

// generateAddresses generates numAddrs of normal AccAddrs and ValAddrs
func generateAddresses(app *creataapp.CreataApp, ctx sdk.Context, numAddrs int) ([]sdk.AccAddress, []sdk.ValAddress) {
	addrDels := creataapp.AddTestAddrsIncremental(app, ctx, numAddrs, sdk.NewInt(10000))
	addrVals := creataapp.ConvertAddrsToValAddrs(addrDels)

	return addrDels, addrVals
}
