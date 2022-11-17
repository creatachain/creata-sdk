package cmd_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/creatachain/creata-sdk/creataapp"
	"github.com/creatachain/creata-sdk/creataapp/creatad/cmd"
	svrcmd "github.com/creatachain/creata-sdk/server/cmd"
	"github.com/creatachain/creata-sdk/x/genutil/client/cli"
)

func TestInitCmd(t *testing.T) {
	rootCmd, _ := cmd.NewRootCmd()
	rootCmd.SetArgs([]string{
		"init",           // Test the init cmd
		"creataapp-test", // Moniker
		fmt.Sprintf("--%s=%s", cli.FlagOverwrite, "true"), // Overwrite genesis.json, in case it already exists
	})

	require.NoError(t, svrcmd.Execute(rootCmd, creataapp.DefaultNodeHome))
}
