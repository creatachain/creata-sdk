package teststaking

import (
	tmcrypto "github.com/creatachain/augusteum/crypto"
	tmtypes "github.com/creatachain/augusteum/types"

	cryptocodec "github.com/creatachain/creata-sdk/crypto/codec"
	"github.com/creatachain/creata-sdk/x/staking/types"
)

// GetTmConsPubKey gets the validator's public key as a tmcrypto.PubKey.
func GetTmConsPubKey(v types.Validator) (tmcrypto.PubKey, error) {
	pk, err := v.ConsPubKey()
	if err != nil {
		return nil, err
	}

	return cryptocodec.ToTmPubKeyInterface(pk)
}

// ToTmValidator casts an SDK validator to a augusteum type Validator.
func ToTmValidator(v types.Validator) (*tmtypes.Validator, error) {
	tmPk, err := GetTmConsPubKey(v)
	if err != nil {
		return nil, err
	}

	return tmtypes.NewValidator(tmPk, v.ConsensusPower()), nil
}

// ToTmValidators casts all validators to the corresponding augusteum type.
func ToTmValidators(v types.Validators) ([]*tmtypes.Validator, error) {
	validators := make([]*tmtypes.Validator, len(v))
	var err error
	for i, val := range v {
		validators[i], err = ToTmValidator(val)
		if err != nil {
			return nil, err
		}
	}

	return validators, nil
}
