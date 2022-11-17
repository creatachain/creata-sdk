package types

import (
	msm "github.com/creatachain/augusteum/msm/types"
)

// Querier defines a function type that a module querier must implement to handle
// custom client queries.
type Querier = func(ctx Context, path []string, req msm.RequestQuery) ([]byte, error)
