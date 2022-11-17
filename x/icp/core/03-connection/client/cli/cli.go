package cli

import (
	"github.com/spf13/cobra"

	"github.com/creatachain/creata-sdk/client"
	"github.com/creatachain/creata-sdk/x/icp/core/03-connection/types"
)

// GetQueryCmd returns the query commands for ICP connections
func GetQueryCmd() *cobra.Command {
	queryCmd := &cobra.Command{
		Use:                        types.SubModuleName,
		Short:                      "ICP connection query subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
	}

	queryCmd.AddCommand(
		GetCmdQueryConnections(),
		GetCmdQueryConnection(),
		GetCmdQueryClientConnections(),
	)

	return queryCmd
}

// NewTxCmd returns a CLI command handler for all x/icp connection transaction commands.
func NewTxCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.SubModuleName,
		Short:                      "ICP connection transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	txCmd.AddCommand(
		NewConnectionOpenInitCmd(),
		NewConnectionOpenTryCmd(),
		NewConnectionOpenAckCmd(),
		NewConnectionOpenConfirmCmd(),
	)

	return txCmd
}
