package baseapp

import (
	"testing"

	"github.com/stretchr/testify/require"

	msm "github.com/creatachain/augusteum/msm/types"

	sdk "github.com/creatachain/creata-sdk/types"
)

var testQuerier = func(_ sdk.Context, _ []string, _ msm.RequestQuery) ([]byte, error) {
	return nil, nil
}

func TestQueryRouter(t *testing.T) {
	qr := NewQueryRouter()

	// require panic on invalid route
	require.Panics(t, func() {
		qr.AddRoute("*", testQuerier)
	})

	qr.AddRoute("testRoute", testQuerier)
	q := qr.Route("testRoute")
	require.NotNil(t, q)

	// require panic on duplicate route
	require.Panics(t, func() {
		qr.AddRoute("testRoute", testQuerier)
	})
}
