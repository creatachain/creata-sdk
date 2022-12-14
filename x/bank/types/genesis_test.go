package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	sdk "github.com/creatachain/creata-sdk/types"
	"github.com/creatachain/creata-sdk/x/bank/types"
)

func TestGenesisStateValidate(t *testing.T) {

	testCases := []struct {
		name         string
		genesisState types.GenesisState
		expErr       bool
	}{
		{
			"valid genesisState",
			types.GenesisState{
				Params: types.DefaultParams(),
				Balances: []types.Balance{
					{
						Address: "creata1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t",
						Coins:   sdk.Coins{sdk.NewInt64Coin("ucta", 1)},
					},
				},
				Supply: sdk.Coins{sdk.NewInt64Coin("ucta", 1)},
				DenomMetadata: []types.Metadata{
					{
						Description: "The native staking token of the Creata Hub.",
						DenomUnits: []*types.DenomUnit{
							{"ucta", uint32(0), []string{"microcta"}},
							{"mcta", uint32(3), []string{"millicta"}},
							{"cta", uint32(6), nil},
						},
						Base:    "ucta",
						Display: "cta",
					},
				},
			},
			false,
		},
		{"empty genesisState", types.GenesisState{}, false},
		{
			"invalid params ",
			types.GenesisState{
				Params: types.Params{
					SendEnabled: []*types.SendEnabled{
						{"", true},
					},
				},
			},
			true,
		},
		{
			"dup balances",
			types.GenesisState{
				Balances: []types.Balance{
					{
						Address: "creata1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t",
						Coins:   sdk.Coins{sdk.NewInt64Coin("ucta", 1)},
					},
					{
						Address: "creata1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t",
						Coins:   sdk.Coins{sdk.NewInt64Coin("ucta", 1)},
					},
				},
			},
			true,
		},
		{
			"0  balance",
			types.GenesisState{
				Balances: []types.Balance{
					{
						Address: "creata1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t",
					},
				},
			},
			false,
		},
		{
			"dup Metadata",
			types.GenesisState{
				DenomMetadata: []types.Metadata{
					{
						Description: "The native staking token of the Creata Hub.",
						DenomUnits: []*types.DenomUnit{
							{"ucta", uint32(0), []string{"microcta"}},
							{"mcta", uint32(3), []string{"millicta"}},
							{"cta", uint32(6), nil},
						},
						Base:    "ucta",
						Display: "cta",
					},
					{
						Description: "The native staking token of the Creata Hub.",
						DenomUnits: []*types.DenomUnit{
							{"ucta", uint32(0), []string{"microcta"}},
							{"mcta", uint32(3), []string{"millicta"}},
							{"cta", uint32(6), nil},
						},
						Base:    "ucta",
						Display: "cta",
					},
				},
			},
			true,
		},
		{
			"invalid Metadata",
			types.GenesisState{
				DenomMetadata: []types.Metadata{
					{
						Description: "The native staking token of the Creata Hub.",
						DenomUnits: []*types.DenomUnit{
							{"ucta", uint32(0), []string{"microcta"}},
							{"mcta", uint32(3), []string{"millicta"}},
							{"cta", uint32(6), nil},
						},
						Base:    "",
						Display: "",
					},
				},
			},
			true,
		},
		{
			"invalid supply",
			types.GenesisState{
				Supply: sdk.Coins{sdk.Coin{Denom: "", Amount: sdk.OneInt()}},
			},
			true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {

			err := tc.genesisState.Validate()

			if tc.expErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
