package transfer

import (
	sdk "github.com/creatachain/creata-sdk/types"
	sdkerrors "github.com/creatachain/creata-sdk/types/errors"
	"github.com/creatachain/creata-sdk/x/icp/applications/transfer/types"
)

// NewHandler returns sdk.Handler for ICP token transfer module messages
func NewHandler(k types.MsgServer) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case *types.MsgTransfer:
			res, err := k.Transfer(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)

		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized ICS-20 transfer message type: %T", msg)
		}
	}
}
