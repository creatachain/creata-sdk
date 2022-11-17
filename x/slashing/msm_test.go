package slashing_test

import (
	"testing"
	"time"

	msm "github.com/creatachain/augusteum/msm/types"
	tmproto "github.com/creatachain/augusteum/proto/augusteum/types"
	"github.com/stretchr/testify/require"

	"github.com/creatachain/creata-sdk/creataapp"
	sdk "github.com/creatachain/creata-sdk/types"
	"github.com/creatachain/creata-sdk/x/slashing"
	"github.com/creatachain/creata-sdk/x/staking"
	"github.com/creatachain/creata-sdk/x/staking/teststaking"
	stakingtypes "github.com/creatachain/creata-sdk/x/staking/types"
)

func TestBeginBlocker(t *testing.T) {
	app := creataapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	pks := creataapp.CreateTestPubKeys(1)
	creataapp.AddTestAddrsFromPubKeys(app, ctx, pks, sdk.TokensFromConsensusPower(200))
	addr, pk := sdk.ValAddress(pks[0].Address()), pks[0]
	tstaking := teststaking.NewHelper(t, ctx, app.StakingKeeper)

	// bond the validator
	power := int64(100)
	amt := tstaking.CreateValidatorWithValPower(addr, pk, power, true)
	staking.EndBlocker(ctx, app.StakingKeeper)
	require.Equal(
		t, app.BankKeeper.GetAllBalances(ctx, sdk.AccAddress(addr)),
		sdk.NewCoins(sdk.NewCoin(app.StakingKeeper.GetParams(ctx).BondDenom, InitTokens.Sub(amt))),
	)
	require.Equal(t, amt, app.StakingKeeper.Validator(ctx, addr).GetBondedTokens())

	val := msm.Validator{
		Address: pk.Address(),
		Power:   power,
	}

	// mark the validator as having signed
	req := msm.RequestBeginBlock{
		LastCommitInfo: msm.LastCommitInfo{
			Votes: []msm.VoteInfo{{
				Validator:       val,
				SignedLastBlock: true,
			}},
		},
	}

	slashing.BeginBlocker(ctx, req, app.SlashingKeeper)

	info, found := app.SlashingKeeper.GetValidatorSigningInfo(ctx, sdk.ConsAddress(pk.Address()))
	require.True(t, found)
	require.Equal(t, ctx.BlockHeight(), info.StartHeight)
	require.Equal(t, int64(1), info.IndexOffset)
	require.Equal(t, time.Unix(0, 0).UTC(), info.JailedUntil)
	require.Equal(t, int64(0), info.MissedBlocksCounter)

	height := int64(0)

	// for 1000 blocks, mark the validator as having signed
	for ; height < app.SlashingKeeper.SignedBlocksWindow(ctx); height++ {
		ctx = ctx.WithBlockHeight(height)
		req = msm.RequestBeginBlock{
			LastCommitInfo: msm.LastCommitInfo{
				Votes: []msm.VoteInfo{{
					Validator:       val,
					SignedLastBlock: true,
				}},
			},
		}

		slashing.BeginBlocker(ctx, req, app.SlashingKeeper)
	}

	// for 500 blocks, mark the validator as having not signed
	for ; height < ((app.SlashingKeeper.SignedBlocksWindow(ctx) * 2) - app.SlashingKeeper.MinSignedPerWindow(ctx) + 1); height++ {
		ctx = ctx.WithBlockHeight(height)
		req = msm.RequestBeginBlock{
			LastCommitInfo: msm.LastCommitInfo{
				Votes: []msm.VoteInfo{{
					Validator:       val,
					SignedLastBlock: false,
				}},
			},
		}

		slashing.BeginBlocker(ctx, req, app.SlashingKeeper)
	}

	// end block
	staking.EndBlocker(ctx, app.StakingKeeper)

	// validator should be jailed
	validator, found := app.StakingKeeper.GetValidatorByConsAddr(ctx, sdk.GetConsAddress(pk))
	require.True(t, found)
	require.Equal(t, stakingtypes.Unbonding, validator.GetStatus())
}
