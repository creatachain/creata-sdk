package augusteum

import (
	"github.com/spf13/cobra"

	"github.com/creatachain/creata-sdk/x/icp/light-clients/07-augusteum/client/cli"
	"github.com/creatachain/creata-sdk/x/icp/light-clients/07-augusteum/types"
)

// Name returns the ICP client name
func Name() string {
	return types.SubModuleName
}

// GetTxCmd returns the root tx command for the ICP client
func GetTxCmd() *cobra.Command {
	return cli.NewTxCmd()
}
