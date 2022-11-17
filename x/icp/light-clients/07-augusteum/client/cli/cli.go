package cli

import (
	"github.com/spf13/cobra"

	"github.com/creatachain/creata-sdk/x/icp/light-clients/07-augusteum/types"
)

// NewTxCmd returns a root CLI command handler for all x/icp/light-clients/07-augusteum transaction commands.
func NewTxCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.SubModuleName,
		Short:                      "Augusteum client transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
	}

	txCmd.AddCommand(
		NewCreateClientCmd(),
		NewUpdateClientCmd(),
		NewSubmitMisbehaviourCmd(),
	)

	return txCmd
}
