package genutil

import (
	msm "github.com/creatachain/augusteum/msm/types"

	"github.com/creatachain/creata-sdk/client"
	sdk "github.com/creatachain/creata-sdk/types"
	"github.com/creatachain/creata-sdk/x/genutil/types"
)

// InitGenesis - initialize accounts and deliver genesis transactions
func InitGenesis(
	ctx sdk.Context, stakingKeeper types.StakingKeeper,
	deliverTx deliverTxfn, genesisState types.GenesisState,
	txEncodingConfig client.TxEncodingConfig,
) (validators []msm.ValidatorUpdate, err error) {
	if len(genesisState.GenTxs) > 0 {
		validators, err = DeliverGenTxs(ctx, genesisState.GenTxs, stakingKeeper, deliverTx, txEncodingConfig)
	}
	return
}
