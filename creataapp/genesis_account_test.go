package creataapp_test

import (
	"testing"
	"time"

	"github.com/creatachain/creata-sdk/creataapp"
	"github.com/creatachain/creata-sdk/crypto/keys/secp256k1"
	sdk "github.com/creatachain/creata-sdk/types"
	authtypes "github.com/creatachain/creata-sdk/x/auth/types"

	"github.com/creatachain/augusteum/crypto"
	"github.com/stretchr/testify/require"
)

func TestSimGenesisAccountValidate(t *testing.T) {
	pubkey := secp256k1.GenPrivKey().PubKey()
	addr := sdk.AccAddress(pubkey.Address())

	vestingStart := time.Now().UTC()

	coins := sdk.NewCoins(sdk.NewInt64Coin("test", 1000))
	baseAcc := authtypes.NewBaseAccount(addr, pubkey, 0, 0)

	testCases := []struct {
		name    string
		sga     creataapp.SimGenesisAccount
		wantErr bool
	}{
		{
			"valid basic account",
			creataapp.SimGenesisAccount{
				BaseAccount: baseAcc,
			},
			false,
		},
		{
			"invalid basic account with mismatching address/pubkey",
			creataapp.SimGenesisAccount{
				BaseAccount: authtypes.NewBaseAccount(addr, secp256k1.GenPrivKey().PubKey(), 0, 0),
			},
			true,
		},
		{
			"valid basic account with module name",
			creataapp.SimGenesisAccount{
				BaseAccount: authtypes.NewBaseAccount(sdk.AccAddress(crypto.AddressHash([]byte("testmod"))), nil, 0, 0),
				ModuleName:  "testmod",
			},
			false,
		},
		{
			"valid basic account with invalid module name/pubkey pair",
			creataapp.SimGenesisAccount{
				BaseAccount: baseAcc,
				ModuleName:  "testmod",
			},
			true,
		},
		{
			"valid basic account with valid vesting attributes",
			creataapp.SimGenesisAccount{
				BaseAccount:     baseAcc,
				OriginalVesting: coins,
				StartTime:       vestingStart.Unix(),
				EndTime:         vestingStart.Add(1 * time.Hour).Unix(),
			},
			false,
		},
		{
			"valid basic account with invalid vesting end time",
			creataapp.SimGenesisAccount{
				BaseAccount:     baseAcc,
				OriginalVesting: coins,
				StartTime:       vestingStart.Add(2 * time.Hour).Unix(),
				EndTime:         vestingStart.Add(1 * time.Hour).Unix(),
			},
			true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.wantErr, tc.sga.Validate() != nil)
		})
	}
}
