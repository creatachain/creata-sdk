package legacytx

import (
	"github.com/creatachain/creata-sdk/codec"
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(StdTx{}, "creata-sdk/StdTx", nil)
}
