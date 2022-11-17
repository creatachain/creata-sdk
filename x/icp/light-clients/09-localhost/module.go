package localhost

import (
	"github.com/creatachain/creata-sdk/x/icp/light-clients/09-localhost/types"
)

// Name returns the ICP client name
func Name() string {
	return types.SubModuleName
}
