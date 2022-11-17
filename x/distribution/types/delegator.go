package types

import (
	sdk "github.com/creatachain/creata-sdk/types"
)

// create a new DelegatorStartingInfo
func NewDelegatorStartingInfo(previousPeriod uint64, ucta sdk.Dec, height uint64) DelegatorStartingInfo {
	return DelegatorStartingInfo{
		PreviousPeriod: previousPeriod,
		Ucta:           ucta,
		Height:         height,
	}
}
