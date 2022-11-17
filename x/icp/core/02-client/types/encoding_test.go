package types_test

import (
	"github.com/creatachain/creata-sdk/x/icp/core/02-client/types"
	icptmtypes "github.com/creatachain/creata-sdk/x/icp/light-clients/07-augusteum/types"
)

func (suite *TypesTestSuite) TestMarshalHeader() {

	cdc := suite.chainA.App.AppCodec()
	h := &icptmtypes.Header{
		TrustedHeight: types.NewHeight(4, 100),
	}

	// marshal header
	bz, err := types.MarshalHeader(cdc, h)
	suite.Require().NoError(err)

	// unmarshal header
	newHeader, err := types.UnmarshalHeader(cdc, bz)
	suite.Require().NoError(err)

	suite.Require().Equal(h, newHeader)

	// use invalid bytes
	invalidHeader, err := types.UnmarshalHeader(cdc, []byte("invalid bytes"))
	suite.Require().Error(err)
	suite.Require().Nil(invalidHeader)

}
