package cli

import (
	"github.com/spf13/cobra"

	"github.com/creatachain/creata-sdk/client"
	"github.com/creatachain/creata-sdk/x/icp/core/02-client/types"
)

// GetQueryCmd returns the query commands for ICP clients
func GetQueryCmd() *cobra.Command {
	queryCmd := &cobra.Command{
		Use:                        types.SubModuleName,
		Short:                      "ICP client query subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	queryCmd.AddCommand(
		GetCmdQueryClientStates(),
		GetCmdQueryClientState(),
		GetCmdQueryConsensusStates(),
		GetCmdQueryConsensusState(),
		GetCmdQueryHeader(),
		GetCmdNodeConsensusState(),
		GetCmdParams(),
	)

	return queryCmd
}
