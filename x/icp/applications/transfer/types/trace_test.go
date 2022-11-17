package types

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseDenomTrace(t *testing.T) {
	testCases := []struct {
		name     string
		denom    string
		expTrace DenomTrace
	}{
		{"empty denom", "", DenomTrace{}},
		{"base denom", "ucta", DenomTrace{BaseDenom: "ucta"}},
		{"trace info", "transfer/channelToA/ucta", DenomTrace{BaseDenom: "ucta", Path: "transfer/channelToA"}},
		{"incomplete path", "transfer/ucta", DenomTrace{BaseDenom: "ucta", Path: "transfer"}},
		{"invalid path (1)", "transfer//ucta", DenomTrace{BaseDenom: "ucta", Path: "transfer/"}},
		{"invalid path (2)", "transfer/channelToA/ucta/", DenomTrace{BaseDenom: "", Path: "transfer/channelToA/ucta"}},
	}

	for _, tc := range testCases {
		trace := ParseDenomTrace(tc.denom)
		require.Equal(t, tc.expTrace, trace, tc.name)
	}
}

func TestDenomTrace_ICPDenom(t *testing.T) {
	testCases := []struct {
		name     string
		trace    DenomTrace
		expDenom string
	}{
		{"base denom", DenomTrace{BaseDenom: "ucta"}, "ucta"},
		{"trace info", DenomTrace{BaseDenom: "ucta", Path: "transfer/channelToA"}, "icp/7F1D3FCF4AE79E1554D670D1AD949A9BA4E4A3C76C63093E17E446A46061A7A2"},
	}

	for _, tc := range testCases {
		denom := tc.trace.ICPDenom()
		require.Equal(t, tc.expDenom, denom, tc.name)
	}
}

func TestDenomTrace_Validate(t *testing.T) {
	testCases := []struct {
		name     string
		trace    DenomTrace
		expError bool
	}{
		{"base denom only", DenomTrace{BaseDenom: "ucta"}, false},
		{"empty DenomTrace", DenomTrace{}, true},
		{"valid single trace info", DenomTrace{BaseDenom: "ucta", Path: "transfer/channelToA"}, false},
		{"valid multiple trace info", DenomTrace{BaseDenom: "ucta", Path: "transfer/channelToA/transfer/channelToB"}, false},
		{"single trace identifier", DenomTrace{BaseDenom: "ucta", Path: "transfer"}, true},
		{"invalid port ID", DenomTrace{BaseDenom: "ucta", Path: "(transfer)/channelToA"}, true},
		{"invalid channel ID", DenomTrace{BaseDenom: "ucta", Path: "transfer/(channelToA)"}, true},
		{"empty base denom with trace", DenomTrace{BaseDenom: "", Path: "transfer/channelToA"}, true},
	}

	for _, tc := range testCases {
		err := tc.trace.Validate()
		if tc.expError {
			require.Error(t, err, tc.name)
			continue
		}
		require.NoError(t, err, tc.name)
	}
}

func TestTraces_Validate(t *testing.T) {
	testCases := []struct {
		name     string
		traces   Traces
		expError bool
	}{
		{"empty Traces", Traces{}, false},
		{"valid multiple trace info", Traces{{BaseDenom: "ucta", Path: "transfer/channelToA/transfer/channelToB"}}, false},
		{
			"valid multiple trace info",
			Traces{
				{BaseDenom: "ucta", Path: "transfer/channelToA/transfer/channelToB"},
				{BaseDenom: "ucta", Path: "transfer/channelToA/transfer/channelToB"},
			},
			true,
		},
		{"empty base denom with trace", Traces{{BaseDenom: "", Path: "transfer/channelToA"}}, true},
	}

	for _, tc := range testCases {
		err := tc.traces.Validate()
		if tc.expError {
			require.Error(t, err, tc.name)
			continue
		}
		require.NoError(t, err, tc.name)
	}
}

func TestValidatePrefixedDenom(t *testing.T) {
	testCases := []struct {
		name     string
		denom    string
		expError bool
	}{
		{"prefixed denom", "transfer/channelToA/ucta", false},
		{"base denom", "ucta", false},
		{"empty denom", "", true},
		{"empty prefix", "/ucta", true},
		{"empty identifiers", "//ucta", true},
		{"single trace identifier", "transfer/", true},
		{"invalid port ID", "(transfer)/channelToA/ucta", true},
		{"invalid channel ID", "transfer/(channelToA)/ucta", true},
	}

	for _, tc := range testCases {
		err := ValidatePrefixedDenom(tc.denom)
		if tc.expError {
			require.Error(t, err, tc.name)
			continue
		}
		require.NoError(t, err, tc.name)
	}
}

func TestValidateICPDenom(t *testing.T) {
	testCases := []struct {
		name     string
		denom    string
		expError bool
	}{
		{"denom with trace hash", "icp/7F1D3FCF4AE79E1554D670D1AD949A9BA4E4A3C76C63093E17E446A46061A7A2", false},
		{"base denom", "ucta", false},
		{"empty denom", "", true},
		{"invalid prefixed denom", "transfer/channelToA/ucta", true},
		{"denom 'icp'", "icp", true},
		{"denom 'icp/'", "icp/", true},
		{"invald prefix", "noticp/7F1D3FCF4AE79E1554D670D1AD949A9BA4E4A3C76C63093E17E446A46061A7A2", true},
		{"invald hash", "icp/!@#$!@#", true},
	}

	for _, tc := range testCases {
		err := ValidateICPDenom(tc.denom)
		if tc.expError {
			require.Error(t, err, tc.name)
			continue
		}
		require.NoError(t, err, tc.name)
	}
}
