package reflection_test

import (
	"context"
	"testing"

	tmproto "github.com/creatachain/augusteum/proto/augusteum/types"
	"github.com/stretchr/testify/suite"

	"github.com/creatachain/creata-sdk/baseapp"
	"github.com/creatachain/creata-sdk/client/grpc/reflection"
	"github.com/creatachain/creata-sdk/creataapp"
)

type IntegrationTestSuite struct {
	suite.Suite

	queryClient reflection.ReflectionServiceClient
}

func (s *IntegrationTestSuite) SetupSuite() {
	app := creataapp.Setup(false)

	sdkCtx := app.BaseApp.NewContext(false, tmproto.Header{})
	queryHelper := baseapp.NewQueryServerTestHelper(sdkCtx, app.InterfaceRegistry())
	queryClient := reflection.NewReflectionServiceClient(queryHelper)
	s.queryClient = queryClient
}

func (s IntegrationTestSuite) TestSimulateService() {
	// We will test the following interface for testing.
	var iface = "creata.evidence.v1beta1.Evidence"

	// Test that "creata.evidence.v1beta1.Evidence" is included in the
	// interfaces.
	resIface, err := s.queryClient.ListAllInterfaces(
		context.Background(),
		&reflection.ListAllInterfacesRequest{},
	)
	s.Require().NoError(err)
	s.Require().Contains(resIface.GetInterfaceNames(), iface)

	// Test that "creata.evidence.v1beta1.Evidence" has at least the
	// Equivocation implementations.
	resImpl, err := s.queryClient.ListImplementations(
		context.Background(),
		&reflection.ListImplementationsRequest{InterfaceName: iface},
	)
	s.Require().NoError(err)
	s.Require().Contains(resImpl.GetImplementationMessageNames(), "/creata.evidence.v1beta1.Equivocation")
}

func TestSimulateTestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}
