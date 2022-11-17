package types_test

import (
	"time"

	tmprotocrypto "github.com/creatachain/augusteum/proto/augusteum/crypto"

	clienttypes "github.com/creatachain/creata-sdk/x/icp/core/02-client/types"
	"github.com/creatachain/creata-sdk/x/icp/core/exported"
	"github.com/creatachain/creata-sdk/x/icp/light-clients/07-augusteum/types"
)

func (suite *AugusteumTestSuite) TestGetHeight() {
	header := suite.chainA.LastHeader
	suite.Require().NotEqual(uint64(0), header.GetHeight())
}

func (suite *AugusteumTestSuite) TestGetTime() {
	header := suite.chainA.LastHeader
	suite.Require().NotEqual(time.Time{}, header.GetTime())
}

func (suite *AugusteumTestSuite) TestHeaderValidateBasic() {
	var (
		header *types.Header
	)
	testCases := []struct {
		name     string
		malleate func()
		expPass  bool
	}{
		{"valid header", func() {}, true},
		{"header is nil", func() {
			header.Header = nil
		}, false},
		{"signed header is nil", func() {
			header.SignedHeader = nil
		}, false},
		{"SignedHeaderFromProto failed", func() {
			header.SignedHeader.Commit.Height = -1
		}, false},
		{"signed header failed augusteum ValidateBasic", func() {
			header = suite.chainA.LastHeader
			header.SignedHeader.Commit = nil
		}, false},
		{"trusted height is greater than header height", func() {
			header.TrustedHeight = header.GetHeight().(clienttypes.Height).Increment().(clienttypes.Height)
		}, false},
		{"validator set nil", func() {
			header.ValidatorSet = nil
		}, false},
		{"ValidatorSetFromProto failed", func() {
			header.ValidatorSet.Validators[0].PubKey = tmprotocrypto.PublicKey{}
		}, false},
		{"header validator hash does not equal hash of validator set", func() {
			// use chainB's randomly generated validator set
			header.ValidatorSet = suite.chainB.LastHeader.ValidatorSet
		}, false},
	}

	suite.Require().Equal(exported.Augusteum, suite.header.ClientType())

	for _, tc := range testCases {
		tc := tc

		suite.Run(tc.name, func() {
			suite.SetupTest()

			header = suite.chainA.LastHeader // must be explicitly changed in malleate

			tc.malleate()

			err := header.ValidateBasic()

			if tc.expPass {
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(err)
			}
		})
	}
}
