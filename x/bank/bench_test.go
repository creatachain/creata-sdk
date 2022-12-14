package bank_test

import (
	"testing"

	msm "github.com/creatachain/augusteum/msm/types"
	tmproto "github.com/creatachain/augusteum/proto/augusteum/types"
	"github.com/stretchr/testify/require"

	"github.com/creatachain/creata-sdk/creataapp"
	creataappparams "github.com/creatachain/creata-sdk/creataapp/params"
	sdk "github.com/creatachain/creata-sdk/types"
	"github.com/creatachain/creata-sdk/x/auth/types"
	authtypes "github.com/creatachain/creata-sdk/x/auth/types"
	stakingtypes "github.com/creatachain/creata-sdk/x/staking/types"
)

var moduleAccAddr = authtypes.NewModuleAddress(stakingtypes.BondedPoolName)

func BenchmarkOneBankSendTxPerBlock(b *testing.B) {
	// Add an account at genesis
	acc := authtypes.BaseAccount{
		Address: addr1.String(),
	}

	// construct genesis state
	genAccs := []types.GenesisAccount{&acc}
	benchmarkApp := creataapp.SetupWithGenesisAccounts(genAccs)
	ctx := benchmarkApp.BaseApp.NewContext(false, tmproto.Header{})

	// some value conceivably higher than the benchmarks would ever go
	err := benchmarkApp.BankKeeper.SetBalances(ctx, addr1, sdk.NewCoins(sdk.NewInt64Coin("foocoin", 100000000000)))
	require.NoError(b, err)

	benchmarkApp.Commit()
	txGen := creataappparams.MakeTestEncodingConfig().TxConfig

	// Precompute all txs
	txs, err := creataapp.GenSequenceOfTxs(txGen, []sdk.Msg{sendMsg1}, []uint64{0}, []uint64{uint64(0)}, b.N, priv1)
	require.NoError(b, err)
	b.ResetTimer()

	height := int64(3)

	// Run this with a profiler, so its easy to distinguish what time comes from
	// Committing, and what time comes from Check/Deliver Tx.
	for i := 0; i < b.N; i++ {
		benchmarkApp.BeginBlock(msm.RequestBeginBlock{Header: tmproto.Header{Height: height}})
		_, _, err := benchmarkApp.Check(txGen.TxEncoder(), txs[i])
		if err != nil {
			panic("something is broken in checking transaction")
		}

		_, _, err = benchmarkApp.Deliver(txGen.TxEncoder(), txs[i])
		require.NoError(b, err)
		benchmarkApp.EndBlock(msm.RequestEndBlock{Height: height})
		benchmarkApp.Commit()
		height++
	}
}

func BenchmarkOneBankMultiSendTxPerBlock(b *testing.B) {
	b.ReportAllocs()
	// Add an account at genesis
	acc := authtypes.BaseAccount{
		Address: addr1.String(),
	}

	// Construct genesis state
	genAccs := []authtypes.GenesisAccount{&acc}
	benchmarkApp := creataapp.SetupWithGenesisAccounts(genAccs)
	ctx := benchmarkApp.BaseApp.NewContext(false, tmproto.Header{})

	// some value conceivably higher than the benchmarks would ever go
	err := benchmarkApp.BankKeeper.SetBalances(ctx, addr1, sdk.NewCoins(sdk.NewInt64Coin("foocoin", 100000000000)))
	require.NoError(b, err)

	benchmarkApp.Commit()
	txGen := creataappparams.MakeTestEncodingConfig().TxConfig

	// Precompute all txs
	txs, err := creataapp.GenSequenceOfTxs(txGen, []sdk.Msg{multiSendMsg1}, []uint64{0}, []uint64{uint64(0)}, b.N, priv1)
	require.NoError(b, err)
	b.ResetTimer()

	height := int64(3)

	// Run this with a profiler, so its easy to distinguish what time comes from
	// Committing, and what time comes from Check/Deliver Tx.
	for i := 0; i < b.N; i++ {
		benchmarkApp.BeginBlock(msm.RequestBeginBlock{Header: tmproto.Header{Height: height}})
		_, _, err := benchmarkApp.Check(txGen.TxEncoder(), txs[i])
		if err != nil {
			panic("something is broken in checking transaction")
		}

		_, _, err = benchmarkApp.Deliver(txGen.TxEncoder(), txs[i])
		require.NoError(b, err)
		benchmarkApp.EndBlock(msm.RequestEndBlock{Height: height})
		benchmarkApp.Commit()
		height++
	}
}
