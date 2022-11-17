package cli

import (
	"github.com/spf13/cobra"

	"github.com/creatachain/creata-sdk/client"
	icpclient "github.com/creatachain/creata-sdk/x/icp/core/02-client"
	connection "github.com/creatachain/creata-sdk/x/icp/core/03-connection"
	channel "github.com/creatachain/creata-sdk/x/icp/core/04-channel"
	host "github.com/creatachain/creata-sdk/x/icp/core/24-host"
	solomachine "github.com/creatachain/creata-sdk/x/icp/light-clients/06-solomachine"
	augusteum "github.com/creatachain/creata-sdk/x/icp/light-clients/07-augusteum"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	icpTxCmd := &cobra.Command{
		Use:                        host.ModuleName,
		Short:                      "ICP transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	icpTxCmd.AddCommand(
		solomachine.GetTxCmd(),
		augusteum.GetTxCmd(),
		connection.GetTxCmd(),
		channel.GetTxCmd(),
	)

	return icpTxCmd
}

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd() *cobra.Command {
	// Group icp queries under a subcommand
	icpQueryCmd := &cobra.Command{
		Use:                        host.ModuleName,
		Short:                      "Querying commands for the ICP module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	icpQueryCmd.AddCommand(
		icpclient.GetQueryCmd(),
		connection.GetQueryCmd(),
		channel.GetQueryCmd(),
	)

	return icpQueryCmd
}
