package mock

import (
	"github.com/creatachain/augusteum/crypto"
	tmproto "github.com/creatachain/augusteum/proto/augusteum/types"
	tmtypes "github.com/creatachain/augusteum/types"

	cryptocodec "github.com/creatachain/creata-sdk/crypto/codec"
	"github.com/creatachain/creata-sdk/crypto/keys/ed25519"
	cryptotypes "github.com/creatachain/creata-sdk/crypto/types"
)

var _ tmtypes.PrivValidator = PV{}

// MockPV implements PrivValidator without any safety or persistence.
// Only use it for testing.
type PV struct {
	PrivKey cryptotypes.PrivKey
}

func NewPV() PV {
	return PV{ed25519.GenPrivKey()}
}

// GetPubKey implements PrivValidator interface
func (pv PV) GetPubKey() (crypto.PubKey, error) {
	return cryptocodec.ToTmPubKeyInterface(pv.PrivKey.PubKey())
}

// SignVote implements PrivValidator interface
func (pv PV) SignVote(chainID string, vote *tmproto.Vote) error {
	signBytes := tmtypes.VoteSignBytes(chainID, vote)
	sig, err := pv.PrivKey.Sign(signBytes)
	if err != nil {
		return err
	}
	vote.Signature = sig
	return nil
}

// SignProposal implements PrivValidator interface
func (pv PV) SignProposal(chainID string, proposal *tmproto.Proposal) error {
	signBytes := tmtypes.ProposalSignBytes(chainID, proposal)
	sig, err := pv.PrivKey.Sign(signBytes)
	if err != nil {
		return err
	}
	proposal.Signature = sig
	return nil
}
