package legacy

import (
	"github.com/creatachain/creata-sdk/codec"
	cryptocodec "github.com/creatachain/creata-sdk/crypto/codec"
	cryptotypes "github.com/creatachain/creata-sdk/crypto/types"
)

// Cdc defines a global generic sealed Amino codec to be used throughout sdk. It
// has all Augusteum crypto and evidence types registered.

var Cdc *codec.LegacyAmino

func init() {
	Cdc = codec.NewLegacyAmino()
	cryptocodec.RegisterCrypto(Cdc)
	codec.RegisterEvidences(Cdc)
}

// PrivKeyFromBytes unmarshals private key bytes and returns a PrivKey
func PrivKeyFromBytes(privKeyBytes []byte) (privKey cryptotypes.PrivKey, err error) {
	err = Cdc.UnmarshalBinaryBare(privKeyBytes, &privKey)
	return
}

// PubKeyFromBytes unmarshals public key bytes and returns a PubKey
func PubKeyFromBytes(pubKeyBytes []byte) (pubKey cryptotypes.PubKey, err error) {
	err = Cdc.UnmarshalBinaryBare(pubKeyBytes, &pubKey)
	return
}
