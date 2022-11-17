package client

import (
	sdk "github.com/creatachain/creata-sdk/types"
	sdkerrors "github.com/creatachain/creata-sdk/types/errors"
	govtypes "github.com/creatachain/creata-sdk/x/gov/types"
	"github.com/creatachain/creata-sdk/x/icp/core/02-client/keeper"
	"github.com/creatachain/creata-sdk/x/icp/core/02-client/types"
)

// NewClientUpdateProposalHandler defines the client update proposal handler
func NewClientUpdateProposalHandler(k keeper.Keeper) govtypes.Handler {
	return func(ctx sdk.Context, content govtypes.Content) error {
		switch c := content.(type) {
		case *types.ClientUpdateProposal:
			return k.ClientUpdateProposal(ctx, c)

		default:
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized icp proposal content type: %T", c)
		}
	}
}
