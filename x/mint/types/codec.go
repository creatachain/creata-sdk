package types

import (
	"github.com/creatachain/creata-sdk/codec"
	cryptocodec "github.com/creatachain/creata-sdk/crypto/codec"
)

var (
	amino = codec.NewLegacyAmino()
)

func init() {
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}
