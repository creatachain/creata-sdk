package types_test

import (
	"github.com/creatachain/creata-sdk/creataapp"
)

var (
	app         = creataapp.Setup(false)
	appCodec, _ = creataapp.MakeCodecs()
)
