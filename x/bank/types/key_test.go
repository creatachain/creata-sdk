package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	sdk "github.com/creatachain/creata-sdk/types"
	"github.com/creatachain/creata-sdk/x/bank/types"
)

func cloneAppend(bz []byte, tail []byte) (res []byte) {
	res = make([]byte, len(bz)+len(tail))
	copy(res, bz)
	copy(res[len(bz):], tail)
	return
}

func TestAddressFromBalancesStore(t *testing.T) {
	addr, err := sdk.AccAddressFromBech32("creata1n88uc38xhjgxzw9nwre4ep2c8ga4fjxcar6mn7")
	require.NoError(t, err)

	key := cloneAppend(addr.Bytes(), []byte("ucta"))
	res := types.AddressFromBalancesStore(key)
	require.Equal(t, res, addr)
}
