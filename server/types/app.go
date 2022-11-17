package types

import (
	"encoding/json"
	"io"

	"github.com/creatachain/augusteum/libs/log"
	msm "github.com/creatachain/augusteum/msm/types"
	tmtypes "github.com/creatachain/augusteum/types"
	dbm "github.com/creatachain/tm-db"
	"github.com/gogo/protobuf/grpc"
	"github.com/spf13/cobra"

	"github.com/creatachain/creata-sdk/client"
	"github.com/creatachain/creata-sdk/server/api"
	"github.com/creatachain/creata-sdk/server/config"
)

type (
	// AppOptions defines an interface that is passed into an application
	// constructor, typically used to set BaseApp options that are either supplied
	// via config file or through CLI arguments/flags. The underlying implementation
	// is defined by the server package and is typically implemented via a Viper
	// literal defined on the server Context. Note, casting Get calls may not yield
	// the expected types and could result in type assertion errors. It is recommend
	// to either use the cast package or perform manual conversion for safety.
	AppOptions interface {
		Get(string) interface{}
	}

	// Application defines an application interface that wraps msm.Application.
	// The interface defines the necessary contracts to be implemented in order
	// to fully bootstrap and start an application.
	Application interface {
		msm.Application

		RegisterAPIRoutes(*api.Server, config.APIConfig)

		// RegisterGRPCServer registers gRPC services directly with the gRPC
		// server.
		RegisterGRPCServer(client.Context, grpc.Server)

		// RegisterTxService registers the gRPC Query service for tx (such as tx
		// simulation, fetching txs by hash...).
		RegisterTxService(clientCtx client.Context)

		// RegisterAugusteumService registers the gRPC Query service for augusteum queries.
		RegisterAugusteumService(clientCtx client.Context)
	}

	// AppCreator is a function that allows us to lazily initialize an
	// application using various configurations.
	AppCreator func(log.Logger, dbm.DB, io.Writer, AppOptions) Application

	// ModuleInitFlags takes a start command and adds modules specific init flags.
	ModuleInitFlags func(startCmd *cobra.Command)

	// ExportedApp represents an exported app state, along with
	// validators, consensus params and latest app height.
	ExportedApp struct {
		// AppState is the application state as JSON.
		AppState json.RawMessage
		// Validators is the exported validator set.
		Validators []tmtypes.GenesisValidator
		// Height is the app's latest block height.
		Height int64
		// ConsensusParams are the exported consensus params for MSM.
		ConsensusParams *msm.ConsensusParams
	}

	// AppExporter is a function that dumps all app state to
	// JSON-serializable structure and returns the current validator set.
	AppExporter func(log.Logger, dbm.DB, io.Writer, int64, bool, []string, AppOptions) (ExportedApp, error)
)
