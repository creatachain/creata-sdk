package types

import (
	"testing"

	"github.com/stretchr/testify/require"

	sdk "github.com/creatachain/creata-sdk/types"
)

func TestValidateGenesis(t *testing.T) {

	fp := InitialFeePool()
	require.Nil(t, fp.ValidateGenesis())

	fp2 := FeePool{CommunityPool: sdk.DecCoins{{Denom: "ucta", Amount: sdk.NewDec(-1)}}}
	require.NotNil(t, fp2.ValidateGenesis())

}
