package testutil

import (
	"context"
	"fmt"

	tmcfg "github.com/creatachain/augusteum/config"
	"github.com/creatachain/augusteum/libs/cli"
	"github.com/creatachain/augusteum/libs/log"
	"github.com/spf13/viper"

	"github.com/creatachain/creata-sdk/client"
	"github.com/creatachain/creata-sdk/codec"
	"github.com/creatachain/creata-sdk/server"
	"github.com/creatachain/creata-sdk/testutil"
	"github.com/creatachain/creata-sdk/types/module"
	genutilcli "github.com/creatachain/creata-sdk/x/genutil/client/cli"
)

func ExecInitCmd(testMbm module.BasicManager, home string, cdc codec.JSONMarshaler) error {
	logger := log.NewNopLogger()
	cfg, err := CreateDefaultAugusteumConfig(home)
	if err != nil {
		return err
	}

	cmd := genutilcli.InitCmd(testMbm, home)
	serverCtx := server.NewContext(viper.New(), cfg, logger)
	clientCtx := client.Context{}.WithJSONMarshaler(cdc).WithHomeDir(home)

	_, out := testutil.ApplyMockIO(cmd)
	clientCtx = clientCtx.WithOutput(out)

	ctx := context.Background()
	ctx = context.WithValue(ctx, client.ClientContextKey, &clientCtx)
	ctx = context.WithValue(ctx, server.ServerContextKey, serverCtx)

	cmd.SetArgs([]string{"appnode-test", fmt.Sprintf("--%s=%s", cli.HomeFlag, home)})

	return cmd.ExecuteContext(ctx)
}

func CreateDefaultAugusteumConfig(rootDir string) (*tmcfg.Config, error) {
	conf := tmcfg.DefaultConfig()
	conf.SetRoot(rootDir)
	tmcfg.EnsureRoot(rootDir)

	if err := conf.ValidateBasic(); err != nil {
		return nil, fmt.Errorf("error in config file: %v", err)
	}

	return conf, nil
}
