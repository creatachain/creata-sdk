package keeper_test

import (
	"strings"

	"github.com/creatachain/creata-sdk/creataapp"

	"github.com/creatachain/creata-sdk/x/evidence/exported"
	"github.com/creatachain/creata-sdk/x/evidence/types"

	msm "github.com/creatachain/augusteum/msm/types"
)

const (
	custom = "custom"
)

func (suite *KeeperTestSuite) TestQuerier_QueryEvidence_Existing() {
	ctx := suite.ctx.WithIsCheckTx(false)
	numEvidence := 100
	_, cdc := creataapp.MakeCodecs()

	evidence := suite.populateEvidence(ctx, numEvidence)
	query := msm.RequestQuery{
		Path: strings.Join([]string{custom, types.QuerierRoute, types.QueryEvidence}, "/"),
		Data: cdc.MustMarshalJSON(types.NewQueryEvidenceRequest(evidence[0].Hash())),
	}

	bz, err := suite.querier(ctx, []string{types.QueryEvidence}, query)
	suite.Nil(err)
	suite.NotNil(bz)

	var e exported.Evidence
	suite.Nil(cdc.UnmarshalJSON(bz, &e))
	suite.Equal(evidence[0], e)
}

func (suite *KeeperTestSuite) TestQuerier_QueryEvidence_NonExisting() {
	ctx := suite.ctx.WithIsCheckTx(false)
	cdc, _ := creataapp.MakeCodecs()
	numEvidence := 100

	suite.populateEvidence(ctx, numEvidence)
	query := msm.RequestQuery{
		Path: strings.Join([]string{custom, types.QuerierRoute, types.QueryEvidence}, "/"),
		Data: cdc.MustMarshalJSON(types.NewQueryEvidenceRequest([]byte("0000000000000000000000000000000000000000000000000000000000000000"))),
	}

	bz, err := suite.querier(ctx, []string{types.QueryEvidence}, query)
	suite.NotNil(err)
	suite.Nil(bz)
}

func (suite *KeeperTestSuite) TestQuerier_QueryAllEvidence() {
	ctx := suite.ctx.WithIsCheckTx(false)
	_, cdc := creataapp.MakeCodecs()
	numEvidence := 100

	suite.populateEvidence(ctx, numEvidence)
	query := msm.RequestQuery{
		Path: strings.Join([]string{custom, types.QuerierRoute, types.QueryAllEvidence}, "/"),
		Data: cdc.MustMarshalJSON(types.NewQueryAllEvidenceParams(1, numEvidence)),
	}

	bz, err := suite.querier(ctx, []string{types.QueryAllEvidence}, query)
	suite.Nil(err)
	suite.NotNil(bz)

	var e []exported.Evidence
	suite.Nil(cdc.UnmarshalJSON(bz, &e))
	suite.Len(e, numEvidence)
}

func (suite *KeeperTestSuite) TestQuerier_QueryAllEvidence_InvalidPagination() {
	ctx := suite.ctx.WithIsCheckTx(false)
	_, cdc := creataapp.MakeCodecs()
	numEvidence := 100

	suite.populateEvidence(ctx, numEvidence)
	query := msm.RequestQuery{
		Path: strings.Join([]string{custom, types.QuerierRoute, types.QueryAllEvidence}, "/"),
		Data: cdc.MustMarshalJSON(types.NewQueryAllEvidenceParams(0, numEvidence)),
	}

	bz, err := suite.querier(ctx, []string{types.QueryAllEvidence}, query)
	suite.Nil(err)
	suite.NotNil(bz)

	var e []exported.Evidence
	suite.Nil(cdc.UnmarshalJSON(bz, &e))
	suite.Len(e, 0)
}
