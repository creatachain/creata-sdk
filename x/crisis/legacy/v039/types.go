package v039

import sdk "github.com/creatachain/creata-sdk/types"

const (
	ModuleName = "crisis"
)

type (
	GenesisState struct {
		ConstantFee sdk.Coin `json:"constant_fee" yaml:"constant_fee"`
	}
)
