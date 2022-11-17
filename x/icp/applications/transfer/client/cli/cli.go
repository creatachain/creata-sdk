package cli

import (
	"github.com/spf13/cobra"

	"github.com/creatachain/creata-sdk/client"
)

// GetQueryCmd returns the query commands for ICP connections
func GetQueryCmd() *cobra.Command {
	queryCmd := &cobra.Command{
		Use:                        "icp-transfer",
		Short:                      "ICP fungible token transfer query subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
	}

	queryCmd.AddCommand(
		GetCmdQueryDenomTrace(),
		GetCmdQueryDenomTraces(),
		GetCmdParams(),
	)

	return queryCmd
}

// NewTxCmd returns the transaction commands for ICP fungible token transfer
func NewTxCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        "icp-transfer",
		Short:                      "ICP fungible token transfer transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	txCmd.AddCommand(
		NewTransferTxCmd(),
	)

	return txCmd
}
