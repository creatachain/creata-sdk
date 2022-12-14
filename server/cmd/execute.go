package cmd

import (
	"context"

	tmcfg "github.com/creatachain/augusteum/config"
	tmcli "github.com/creatachain/augusteum/libs/cli"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"

	"github.com/creatachain/creata-sdk/client"
	"github.com/creatachain/creata-sdk/client/flags"
	"github.com/creatachain/creata-sdk/server"
)

// Execute executes the root command of an application. It handles creating a
// server context object with the appropriate server and client objects injected
// into the underlying stdlib Context. It also handles adding core CLI flags,
// specifically the logging flags. It returns an error upon execution failure.
func Execute(rootCmd *cobra.Command, defaultHome string) error {
	// Create and set a client.Context on the command's Context. During the pre-run
	// of the root command, a default initialized client.Context is provided to
	// seed child command execution with values such as AccountRetriver, Keyring,
	// and a Augusteum RPC. This requires the use of a pointer reference when
	// getting and setting the client.Context. Ideally,
	srvCtx := server.NewDefaultContext()
	ctx := context.Background()
	ctx = context.WithValue(ctx, client.ClientContextKey, &client.Context{})
	ctx = context.WithValue(ctx, server.ServerContextKey, srvCtx)

	rootCmd.PersistentFlags().String(flags.FlagLogLevel, zerolog.InfoLevel.String(), "The logging level (trace|debug|info|warn|error|fatal|panic)")
	rootCmd.PersistentFlags().String(flags.FlagLogFormat, tmcfg.LogFormatPlain, "The logging format (json|plain)")

	executor := tmcli.PrepareBaseCmd(rootCmd, "", defaultHome)
	return executor.ExecuteContext(ctx)
}
