package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/creatachain/creata-sdk/codec"
	"github.com/creatachain/creata-sdk/x/bank/types"
)

func TestMetadataValidate(t *testing.T) {
	testCases := []struct {
		name     string
		metadata types.Metadata
		expErr   bool
	}{
		{
			"non-empty coins",
			types.Metadata{
				Description: "The native staking token of the Creata Hub.",
				DenomUnits: []*types.DenomUnit{
					{"ucta", uint32(0), []string{"microcta"}},
					{"mcta", uint32(3), []string{"millicta"}},
					{"cta", uint32(6), nil},
				},
				Base:    "ucta",
				Display: "cta",
			},
			false,
		},
		{"empty metadata", types.Metadata{}, true},
		{
			"invalid base denom",
			types.Metadata{
				Base: "",
			},
			true,
		},
		{
			"invalid display denom",
			types.Metadata{
				Base:    "ucta",
				Display: "",
			},
			true,
		},
		{
			"duplicate denom unit",
			types.Metadata{
				Description: "The native staking token of the Creata Hub.",
				DenomUnits: []*types.DenomUnit{
					{"ucta", uint32(0), []string{"microcta"}},
					{"ucta", uint32(1), []string{"microcta"}},
				},
				Base:    "ucta",
				Display: "cta",
			},
			true,
		},
		{
			"invalid denom unit",
			types.Metadata{
				Description: "The native staking token of the Creata Hub.",
				DenomUnits: []*types.DenomUnit{
					{"", uint32(0), []string{"microcta"}},
				},
				Base:    "ucta",
				Display: "cta",
			},
			true,
		},
		{
			"invalid denom unit alias",
			types.Metadata{
				Description: "The native staking token of the Creata Hub.",
				DenomUnits: []*types.DenomUnit{
					{"ucta", uint32(0), []string{""}},
				},
				Base:    "ucta",
				Display: "cta",
			},
			true,
		},
		{
			"duplicate denom unit alias",
			types.Metadata{
				Description: "The native staking token of the Creata Hub.",
				DenomUnits: []*types.DenomUnit{
					{"ucta", uint32(0), []string{"microcta", "microcta"}},
				},
				Base:    "ucta",
				Display: "cta",
			},
			true,
		},
		{
			"no base denom unit",
			types.Metadata{
				Description: "The native staking token of the Creata Hub.",
				DenomUnits: []*types.DenomUnit{
					{"mcta", uint32(3), []string{"millicta"}},
					{"cta", uint32(6), nil},
				},
				Base:    "ucta",
				Display: "cta",
			},
			true,
		},
		{
			"base denom exponent not zero",
			types.Metadata{
				Description: "The native staking token of the Creata Hub.",
				DenomUnits: []*types.DenomUnit{
					{"ucta", uint32(1), []string{"microcta"}},
					{"mcta", uint32(3), []string{"millicta"}},
					{"cta", uint32(6), nil},
				},
				Base:    "ucta",
				Display: "cta",
			},
			true,
		},
		{
			"no display denom unit",
			types.Metadata{
				Description: "The native staking token of the Creata Hub.",
				DenomUnits: []*types.DenomUnit{
					{"ucta", uint32(0), []string{"microcta"}},
				},
				Base:    "ucta",
				Display: "cta",
			},
			true,
		},
		{
			"denom units not sorted",
			types.Metadata{
				Description: "The native staking token of the Creata Hub.",
				DenomUnits: []*types.DenomUnit{
					{"ucta", uint32(0), []string{"microcta"}},
					{"cta", uint32(6), nil},
					{"mcta", uint32(3), []string{"millicta"}},
				},
				Base:    "ucta",
				Display: "cta",
			},
			true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {

			err := tc.metadata.Validate()

			if tc.expErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestMarshalJSONMetaData(t *testing.T) {
	cdc := codec.NewLegacyAmino()

	testCases := []struct {
		name      string
		input     []types.Metadata
		strOutput string
	}{
		{"nil metadata", nil, `null`},
		{"empty metadata", []types.Metadata{}, `[]`},
		{"non-empty coins", []types.Metadata{{
			Description: "The native staking token of the Creata Hub.",
			DenomUnits: []*types.DenomUnit{
				{"ucta", uint32(0), []string{"microcta"}}, // The default exponent value 0 is omitted in the json
				{"mcta", uint32(3), []string{"millicta"}},
				{"cta", uint32(6), nil},
			},
			Base:    "ucta",
			Display: "cta",
		},
		},
			`[{"description":"The native staking token of the Creata Hub.","denom_units":[{"denom":"ucta","aliases":["microcta"]},{"denom":"mcta","exponent":3,"aliases":["millicta"]},{"denom":"cta","exponent":6}],"base":"ucta","display":"cta"}]`},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			bz, err := cdc.MarshalJSON(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.strOutput, string(bz))

			var newMetadata []types.Metadata
			require.NoError(t, cdc.UnmarshalJSON(bz, &newMetadata))

			if len(tc.input) == 0 {
				require.Nil(t, newMetadata)
			} else {
				require.Equal(t, tc.input, newMetadata)
			}
		})
	}
}
