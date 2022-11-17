package creataapp

import (
	creataappparams "github.com/creatachain/creata-sdk/creataapp/params"
	"github.com/creatachain/creata-sdk/std"
)

// MakeTestEncodingConfig creates an EncodingConfig for testing.
// This function should be used only internally (in the SDK).
// App user should'nt create new codecs - use the app.AppCodec instead.
// [DEPRECATED]
func MakeTestEncodingConfig() creataappparams.EncodingConfig {
	encodingConfig := creataappparams.MakeTestEncodingConfig()
	std.RegisterLegacyAminoCodec(encodingConfig.Amino)
	std.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	ModuleBasics.RegisterLegacyAminoCodec(encodingConfig.Amino)
	ModuleBasics.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	return encodingConfig
}
