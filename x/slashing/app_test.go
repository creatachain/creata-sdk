package slashing_test

import (
	"errors"
	"testing"

	msm "github.com/creatachain/augusteum/msm/types"
	tmproto "github.com/creatachain/augusteum/proto/augusteum/types"
	"github.com/stretchr/testify/require"

	"github.com/creatachain/creata-sdk/creataapp"
	"github.com/creatachain/creata-sdk/crypto/keys/ed25519"
	"github.com/creatachain/creata-sdk/crypto/keys/secp256k1"
	sdk "github.com/creatachain/creata-sdk/types"
	authtypes "github.com/creatachain/creata-sdk/x/auth/types"
	banktypes "github.com/creatachain/creata-sdk/x/bank/types"
	"github.com/creatachain/creata-sdk/x/slashing/types"
	stakingtypes "github.com/creatachain/creata-sdk/x/staking/types"
)

var (
	priv1 = secp256k1.GenPrivKey()
	addr1 = sdk.AccAddress(priv1.PubKey().Address())

	valKey  = ed25519.GenPrivKey()
	valAddr = sdk.AccAddress(valKey.PubKey().Address())
)

func checkValidator(t *testing.T, app *creataapp.CreataApp, _ sdk.AccAddress, expFound bool) stakingtypes.Validator {
	ctxCheck := app.BaseApp.NewContext(true, tmproto.Header{})
	validator, found := app.StakingKeeper.GetValidator(ctxCheck, sdk.ValAddress(addr1))
	require.Equal(t, expFound, found)
	return validator
}

func checkValidatorSigningInfo(t *testing.T, app *creataapp.CreataApp, addr sdk.ConsAddress, expFound bool) types.ValidatorSigningInfo {
	ctxCheck := app.BaseApp.NewContext(true, tmproto.Header{})
	signingInfo, found := app.SlashingKeeper.GetValidatorSigningInfo(ctxCheck, addr)
	require.Equal(t, expFound, found)
	return signingInfo
}

func TestSlashingMsgs(t *testing.T) {
	genTokens := sdk.TokensFromConsensusPower(42)
	bondTokens := sdk.TokensFromConsensusPower(10)
	genCoin := sdk.NewCoin(sdk.DefaultBondDenom, genTokens)
	bondCoin := sdk.NewCoin(sdk.DefaultBondDenom, bondTokens)

	acc1 := &authtypes.BaseAccount{
		Address: addr1.String(),
	}
	accs := authtypes.GenesisAccounts{acc1}
	balances := []banktypes.Balance{
		{
			Address: addr1.String(),
			Coins:   sdk.Coins{genCoin},
		},
	}

	app := creataapp.SetupWithGenesisAccounts(accs, balances...)
	creataapp.CheckBalance(t, app, addr1, sdk.Coins{genCoin})

	description := stakingtypes.NewDescription("foo_moniker", "", "", "", "")
	commission := stakingtypes.NewCommissionRates(sdk.ZeroDec(), sdk.ZeroDec(), sdk.ZeroDec())

	createValidatorMsg, err := stakingtypes.NewMsgCreateValidator(
		sdk.ValAddress(addr1), valKey.PubKey(), bondCoin, description, commission, sdk.OneInt(),
	)
	require.NoError(t, err)

	header := tmproto.Header{Height: app.LastBlockHeight() + 1}
	txGen := creataapp.MakeTestEncodingConfig().TxConfig
	_, _, err = creataapp.SignCheckDeliver(t, txGen, app.BaseApp, header, []sdk.Msg{createValidatorMsg}, "", []uint64{0}, []uint64{0}, true, true, priv1)
	require.NoError(t, err)
	creataapp.CheckBalance(t, app, addr1, sdk.Coins{genCoin.Sub(bondCoin)})

	header = tmproto.Header{Height: app.LastBlockHeight() + 1}
	app.BeginBlock(msm.RequestBeginBlock{Header: header})

	validator := checkValidator(t, app, addr1, true)
	require.Equal(t, sdk.ValAddress(addr1).String(), validator.OperatorAddress)
	require.Equal(t, stakingtypes.Bonded, validator.Status)
	require.True(sdk.IntEq(t, bondTokens, validator.BondedTokens()))
	unjailMsg := &types.MsgUnjail{ValidatorAddr: sdk.ValAddress(addr1).String()}

	checkValidatorSigningInfo(t, app, sdk.ConsAddress(valAddr), true)

	// unjail should fail with unknown validator
	header = tmproto.Header{Height: app.LastBlockHeight() + 1}
	_, res, err := creataapp.SignCheckDeliver(t, txGen, app.BaseApp, header, []sdk.Msg{unjailMsg}, "", []uint64{0}, []uint64{1}, false, false, priv1)
	require.Error(t, err)
	require.Nil(t, res)
	require.True(t, errors.Is(types.ErrValidatorNotJailed, err))
}
