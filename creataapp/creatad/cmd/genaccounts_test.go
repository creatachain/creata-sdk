package cmd_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/creatachain/augusteum/libs/log"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"

	"github.com/creatachain/creata-sdk/client"
	"github.com/creatachain/creata-sdk/client/flags"
	"github.com/creatachain/creata-sdk/creataapp"
	simcmd "github.com/creatachain/creata-sdk/creataapp/creatad/cmd"
	"github.com/creatachain/creata-sdk/server"
	"github.com/creatachain/creata-sdk/testutil/testdata"
	"github.com/creatachain/creata-sdk/types/module"
	"github.com/creatachain/creata-sdk/x/genutil"
	genutiltest "github.com/creatachain/creata-sdk/x/genutil/client/testutil"
)

var testMbm = module.NewBasicManager(genutil.AppModuleBasic{})

func TestAddGenesisAccountCmd(t *testing.T) {
	_, _, addr1 := testdata.KeyTestPubAddr()
	tests := []struct {
		name      string
		addr      string
		denom     string
		expectErr bool
	}{
		{
			name:      "invalid address",
			addr:      "",
			denom:     "1000cta",
			expectErr: true,
		},
		{
			name:      "valid address",
			addr:      addr1.String(),
			denom:     "1000cta",
			expectErr: false,
		},
		{
			name:      "multiple denoms",
			addr:      addr1.String(),
			denom:     "1000cta, 2000fcta",
			expectErr: false,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			home := t.TempDir()
			logger := log.NewNopLogger()
			cfg, err := genutiltest.CreateDefaultAugusteumConfig(home)
			require.NoError(t, err)

			appCodec, _ := creataapp.MakeCodecs()
			err = genutiltest.ExecInitCmd(testMbm, home, appCodec)
			require.NoError(t, err)

			serverCtx := server.NewContext(viper.New(), cfg, logger)
			clientCtx := client.Context{}.WithJSONMarshaler(appCodec).WithHomeDir(home)

			ctx := context.Background()
			ctx = context.WithValue(ctx, client.ClientContextKey, &clientCtx)
			ctx = context.WithValue(ctx, server.ServerContextKey, serverCtx)

			cmd := simcmd.AddGenesisAccountCmd(home)
			cmd.SetArgs([]string{
				tc.addr,
				tc.denom,
				fmt.Sprintf("--%s=home", flags.FlagHome)})

			if tc.expectErr {
				require.Error(t, cmd.ExecuteContext(ctx))
			} else {
				require.NoError(t, cmd.ExecuteContext(ctx))
			}
		})
	}
}
