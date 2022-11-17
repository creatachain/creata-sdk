package baseapp_test

import (
	"testing"

	msm "github.com/creatachain/augusteum/msm/types"
	tmproto "github.com/creatachain/augusteum/proto/augusteum/types"
	"github.com/stretchr/testify/require"

	"github.com/creatachain/creata-sdk/baseapp"
)

func TestValidateBlockParams(t *testing.T) {
	testCases := []struct {
		arg       interface{}
		expectErr bool
	}{
		{nil, true},
		{&msm.BlockParams{}, true},
		{msm.BlockParams{}, true},
		{msm.BlockParams{MaxBytes: -1, MaxGas: -1}, true},
		{msm.BlockParams{MaxBytes: 2000000, MaxGas: -5}, true},
		{msm.BlockParams{MaxBytes: 2000000, MaxGas: 300000}, false},
	}

	for _, tc := range testCases {
		require.Equal(t, tc.expectErr, baseapp.ValidateBlockParams(tc.arg) != nil)
	}
}

func TestValidateEvidenceParams(t *testing.T) {
	testCases := []struct {
		arg       interface{}
		expectErr bool
	}{
		{nil, true},
		{&tmproto.EvidenceParams{}, true},
		{tmproto.EvidenceParams{}, true},
		{tmproto.EvidenceParams{MaxAgeNumBlocks: -1, MaxAgeDuration: 18004000, MaxBytes: 5000000}, true},
		{tmproto.EvidenceParams{MaxAgeNumBlocks: 360000, MaxAgeDuration: -1, MaxBytes: 5000000}, true},
		{tmproto.EvidenceParams{MaxAgeNumBlocks: 360000, MaxAgeDuration: 18004000, MaxBytes: -1}, true},
		{tmproto.EvidenceParams{MaxAgeNumBlocks: 360000, MaxAgeDuration: 18004000, MaxBytes: 5000000}, false},
		{tmproto.EvidenceParams{MaxAgeNumBlocks: 360000, MaxAgeDuration: 18004000, MaxBytes: 0}, false},
	}

	for _, tc := range testCases {
		require.Equal(t, tc.expectErr, baseapp.ValidateEvidenceParams(tc.arg) != nil)
	}
}

func TestValidateValidatorParams(t *testing.T) {
	testCases := []struct {
		arg       interface{}
		expectErr bool
	}{
		{nil, true},
		{&tmproto.ValidatorParams{}, true},
		{tmproto.ValidatorParams{}, true},
		{tmproto.ValidatorParams{PubKeyTypes: []string{}}, true},
		{tmproto.ValidatorParams{PubKeyTypes: []string{"secp256k1"}}, false},
	}

	for _, tc := range testCases {
		require.Equal(t, tc.expectErr, baseapp.ValidateValidatorParams(tc.arg) != nil)
	}
}
