package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	codectypes "github.com/creatachain/creata-sdk/codec/types"
	cryptocodec "github.com/creatachain/creata-sdk/crypto/codec"
	"github.com/creatachain/creata-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/creatachain/creata-sdk/crypto/types"
	"github.com/creatachain/creata-sdk/testutil/testdata"
	sdk "github.com/creatachain/creata-sdk/types"
	host "github.com/creatachain/creata-sdk/x/icp/core/24-host"
	"github.com/creatachain/creata-sdk/x/icp/core/exported"
	"github.com/creatachain/creata-sdk/x/icp/light-clients/06-solomachine/types"
	icptesting "github.com/creatachain/creata-sdk/x/icp/testing"
)

type SoloMachineTestSuite struct {
	suite.Suite

	solomachine      *icptesting.Solomachine // singlesig public key
	solomachineMulti *icptesting.Solomachine // multisig public key
	coordinator      *icptesting.Coordinator

	// testing chain used for convenience and readability
	chainA *icptesting.TestChain
	chainB *icptesting.TestChain

	store sdk.KVStore
}

func (suite *SoloMachineTestSuite) SetupTest() {
	suite.coordinator = icptesting.NewCoordinator(suite.T(), 2)
	suite.chainA = suite.coordinator.GetChain(icptesting.GetChainID(0))
	suite.chainB = suite.coordinator.GetChain(icptesting.GetChainID(1))

	suite.solomachine = icptesting.NewSolomachine(suite.T(), suite.chainA.Codec, "solomachinesingle", "testing", 1)
	suite.solomachineMulti = icptesting.NewSolomachine(suite.T(), suite.chainA.Codec, "solomachinemulti", "testing", 4)

	suite.store = suite.chainA.App.ICPKeeper.ClientKeeper.ClientStore(suite.chainA.GetContext(), exported.Solomachine)
}

func TestSoloMachineTestSuite(t *testing.T) {
	suite.Run(t, new(SoloMachineTestSuite))
}

func (suite *SoloMachineTestSuite) GetSequenceFromStore() uint64 {
	bz := suite.store.Get(host.ClientStateKey())
	suite.Require().NotNil(bz)

	var clientState exported.ClientState
	err := suite.chainA.Codec.UnmarshalInterface(bz, &clientState)
	suite.Require().NoError(err)
	return clientState.GetLatestHeight().GetRevisionHeight()
}

func (suite *SoloMachineTestSuite) GetInvalidProof() []byte {
	invalidProof, err := suite.chainA.Codec.MarshalBinaryBare(&types.TimestampedSignatureData{Timestamp: suite.solomachine.Time})
	suite.Require().NoError(err)

	return invalidProof
}

func TestUnpackInterfaces_Header(t *testing.T) {
	registry := testdata.NewTestInterfaceRegistry()
	cryptocodec.RegisterInterfaces(registry)

	pk := secp256k1.GenPrivKey().PubKey().(cryptotypes.PubKey)
	any, err := codectypes.NewAnyWithValue(pk)
	require.NoError(t, err)

	header := types.Header{
		NewPublicKey: any,
	}
	bz, err := header.Marshal()
	require.NoError(t, err)

	var header2 types.Header
	err = header2.Unmarshal(bz)
	require.NoError(t, err)

	err = codectypes.UnpackInterfaces(header2, registry)
	require.NoError(t, err)

	require.Equal(t, pk, header2.NewPublicKey.GetCachedValue())
}

func TestUnpackInterfaces_HeaderData(t *testing.T) {
	registry := testdata.NewTestInterfaceRegistry()
	cryptocodec.RegisterInterfaces(registry)

	pk := secp256k1.GenPrivKey().PubKey().(cryptotypes.PubKey)
	any, err := codectypes.NewAnyWithValue(pk)
	require.NoError(t, err)

	hd := types.HeaderData{
		NewPubKey: any,
	}
	bz, err := hd.Marshal()
	require.NoError(t, err)

	var hd2 types.HeaderData
	err = hd2.Unmarshal(bz)
	require.NoError(t, err)

	err = codectypes.UnpackInterfaces(hd2, registry)
	require.NoError(t, err)

	require.Equal(t, pk, hd2.NewPubKey.GetCachedValue())
}
