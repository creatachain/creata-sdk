package keys

import (
	"github.com/creatachain/creata-sdk/codec"
	cryptocodec "github.com/creatachain/creata-sdk/crypto/codec"
)

// KeysCdc defines codec to be used with key operations
var KeysCdc *codec.LegacyAmino

func init() {
	KeysCdc = codec.NewLegacyAmino()
	cryptocodec.RegisterCrypto(KeysCdc)
	KeysCdc.Seal()
}

// marshal keys
func MarshalJSON(o interface{}) ([]byte, error) {
	return KeysCdc.MarshalJSON(o)
}

// unmarshal json
func UnmarshalJSON(bz []byte, ptr interface{}) error {
	return KeysCdc.UnmarshalJSON(bz, ptr)
}
