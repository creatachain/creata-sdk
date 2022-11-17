package staking_test

import (
	"fmt"
	"testing"

	msm "github.com/creatachain/augusteum/msm/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	codectypes "github.com/creatachain/creata-sdk/codec/types"
	"github.com/creatachain/creata-sdk/creataapp"
	"github.com/creatachain/creata-sdk/crypto/keys/ed25519"
	sdk "github.com/creatachain/creata-sdk/types"
	banktypes "github.com/creatachain/creata-sdk/x/bank/types"
	"github.com/creatachain/creata-sdk/x/staking"
	"github.com/creatachain/creata-sdk/x/staking/teststaking"
	"github.com/creatachain/creata-sdk/x/staking/types"
)

func bootstrapGenesisTest(t *testing.T, power int64, numAddrs int) (*creataapp.CreataApp, sdk.Context, []sdk.AccAddress) {
	_, app, ctx := getBaseCreataappWithCustomKeeper()

	addrDels, _ := generateAddresses(app, ctx, numAddrs, sdk.NewInt(10000))

	amt := sdk.TokensFromConsensusPower(power)
	totalSupply := sdk.NewCoins(sdk.NewCoin(app.StakingKeeper.BondDenom(ctx), amt.MulRaw(int64(len(addrDels)))))

	notBondedPool := app.StakingKeeper.GetNotBondedPool(ctx)
	err := app.BankKeeper.SetBalances(ctx, notBondedPool.GetAddress(), totalSupply)
	require.NoError(t, err)

	app.AccountKeeper.SetModuleAccount(ctx, notBondedPool)
	app.BankKeeper.SetSupply(ctx, banktypes.NewSupply(totalSupply))

	return app, ctx, addrDels
}

func TestInitGenesis(t *testing.T) {
	app, ctx, addrs := bootstrapGenesisTest(t, 1000, 10)

	valTokens := sdk.TokensFromConsensusPower(1)

	params := app.StakingKeeper.GetParams(ctx)
	validators := make([]types.Validator, 2)
	var delegations []types.Delegation

	pk0, err := codectypes.NewAnyWithValue(PKs[0])
	require.NoError(t, err)

	pk1, err := codectypes.NewAnyWithValue(PKs[1])
	require.NoError(t, err)

	// initialize the validators
	validators[0].OperatorAddress = sdk.ValAddress(addrs[0]).String()
	validators[0].ConsensusPubkey = pk0
	validators[0].Description = types.NewDescription("hoop", "", "", "", "")
	validators[0].Status = types.Bonded
	validators[0].Tokens = valTokens
	validators[0].DelegatorShares = valTokens.ToDec()
	validators[1].OperatorAddress = sdk.ValAddress(addrs[1]).String()
	validators[1].ConsensusPubkey = pk1
	validators[1].Description = types.NewDescription("bloop", "", "", "", "")
	validators[1].Status = types.Bonded
	validators[1].Tokens = valTokens
	validators[1].DelegatorShares = valTokens.ToDec()

	genesisState := types.NewGenesisState(params, validators, delegations)
	vals := staking.InitGenesis(ctx, app.StakingKeeper, app.AccountKeeper, app.BankKeeper, genesisState)

	actualGenesis := staking.ExportGenesis(ctx, app.StakingKeeper)
	require.Equal(t, genesisState.Params, actualGenesis.Params)
	require.Equal(t, genesisState.Delegations, actualGenesis.Delegations)
	require.EqualValues(t, app.StakingKeeper.GetAllValidators(ctx), actualGenesis.Validators)

	// Ensure validators have addresses.
	vals2, err := staking.WriteValidators(ctx, app.StakingKeeper)
	require.NoError(t, err)
	for _, val := range vals2 {
		require.NotEmpty(t, val.Address)
	}

	// now make sure the validators are bonded and intra-tx counters are correct
	resVal, found := app.StakingKeeper.GetValidator(ctx, sdk.ValAddress(addrs[0]))
	require.True(t, found)
	require.Equal(t, types.Bonded, resVal.Status)

	resVal, found = app.StakingKeeper.GetValidator(ctx, sdk.ValAddress(addrs[1]))
	require.True(t, found)
	require.Equal(t, types.Bonded, resVal.Status)

	msmvals := make([]msm.ValidatorUpdate, len(vals))
	for i, val := range validators {
		msmvals[i] = val.MSMValidatorUpdate()
	}

	require.Equal(t, msmvals, vals)
}

func TestInitGenesisLargeValidatorSet(t *testing.T) {
	size := 200
	require.True(t, size > 100)

	app, ctx, addrs := bootstrapGenesisTest(t, 1000, 200)

	params := app.StakingKeeper.GetParams(ctx)
	delegations := []types.Delegation{}
	validators := make([]types.Validator, size)
	var err error
	for i := range validators {
		validators[i], err = types.NewValidator(sdk.ValAddress(addrs[i]),
			PKs[i], types.NewDescription(fmt.Sprintf("#%d", i), "", "", "", ""))
		require.NoError(t, err)
		validators[i].Status = types.Bonded

		tokens := sdk.TokensFromConsensusPower(1)
		if i < 100 {
			tokens = sdk.TokensFromConsensusPower(2)
		}
		validators[i].Tokens = tokens
		validators[i].DelegatorShares = tokens.ToDec()
	}

	genesisState := types.NewGenesisState(params, validators, delegations)
	vals := staking.InitGenesis(ctx, app.StakingKeeper, app.AccountKeeper, app.BankKeeper, genesisState)

	msmvals := make([]msm.ValidatorUpdate, 100)
	for i, val := range validators[:100] {
		msmvals[i] = val.MSMValidatorUpdate()
	}

	require.Equal(t, msmvals, vals)
}

func TestValidateGenesis(t *testing.T) {
	genValidators1 := make([]types.Validator, 1, 5)
	pk := ed25519.GenPrivKey().PubKey()
	genValidators1[0] = teststaking.NewValidator(t, sdk.ValAddress(pk.Address()), pk)
	genValidators1[0].Tokens = sdk.OneInt()
	genValidators1[0].DelegatorShares = sdk.OneDec()

	tests := []struct {
		name    string
		mutate  func(*types.GenesisState)
		wantErr bool
	}{
		{"default", func(*types.GenesisState) {}, false},
		// validate genesis validators
		{"duplicate validator", func(data *types.GenesisState) {
			data.Validators = genValidators1
			data.Validators = append(data.Validators, genValidators1[0])
		}, true},
		{"no delegator shares", func(data *types.GenesisState) {
			data.Validators = genValidators1
			data.Validators[0].DelegatorShares = sdk.ZeroDec()
		}, true},
		{"jailed and bonded validator", func(data *types.GenesisState) {
			data.Validators = genValidators1
			data.Validators[0].Jailed = true
			data.Validators[0].Status = types.Bonded
		}, true},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			genesisState := types.DefaultGenesisState()
			tt.mutate(genesisState)
			if tt.wantErr {
				assert.Error(t, staking.ValidateGenesis(genesisState))
			} else {
				assert.NoError(t, staking.ValidateGenesis(genesisState))
			}
		})
	}
}
