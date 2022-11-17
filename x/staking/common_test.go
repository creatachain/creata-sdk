package staking_test

import (
	"math/big"

	tmproto "github.com/creatachain/augusteum/proto/augusteum/types"

	"github.com/creatachain/creata-sdk/codec"
	"github.com/creatachain/creata-sdk/creataapp"
	"github.com/creatachain/creata-sdk/crypto/keys/ed25519"
	"github.com/creatachain/creata-sdk/crypto/keys/secp256k1"
	sdk "github.com/creatachain/creata-sdk/types"
	"github.com/creatachain/creata-sdk/x/staking/keeper"
	"github.com/creatachain/creata-sdk/x/staking/types"
)

func init() {
	sdk.PowerReduction = sdk.NewIntFromBigInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil))
}

// nolint:deadcode,unused,varcheck
var (
	priv1 = secp256k1.GenPrivKey()
	addr1 = sdk.AccAddress(priv1.PubKey().Address())
	priv2 = secp256k1.GenPrivKey()
	addr2 = sdk.AccAddress(priv2.PubKey().Address())

	valKey  = ed25519.GenPrivKey()
	valAddr = sdk.AccAddress(valKey.PubKey().Address())

	commissionRates = types.NewCommissionRates(sdk.ZeroDec(), sdk.ZeroDec(), sdk.ZeroDec())

	PKs = creataapp.CreateTestPubKeys(500)
)

// getBaseCreataappWithCustomKeeper Returns a creataapp with custom StakingKeeper
// to avoid messing with the hooks.
func getBaseCreataappWithCustomKeeper() (*codec.LegacyAmino, *creataapp.CreataApp, sdk.Context) {
	app := creataapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	appCodec := app.AppCodec()

	app.StakingKeeper = keeper.NewKeeper(
		appCodec,
		app.GetKey(types.StoreKey),
		app.AccountKeeper,
		app.BankKeeper,
		app.GetSubspace(types.ModuleName),
	)
	app.StakingKeeper.SetParams(ctx, types.DefaultParams())

	return codec.NewLegacyAmino(), app, ctx
}

// generateAddresses generates numAddrs of normal AccAddrs and ValAddrs
func generateAddresses(app *creataapp.CreataApp, ctx sdk.Context, numAddrs int, accAmount sdk.Int) ([]sdk.AccAddress, []sdk.ValAddress) {
	addrDels := creataapp.AddTestAddrsIncremental(app, ctx, numAddrs, accAmount)
	addrVals := creataapp.ConvertAddrsToValAddrs(addrDels)

	return addrDels, addrVals
}
