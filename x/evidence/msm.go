package evidence

import (
	"fmt"
	"time"

	msm "github.com/creatachain/augusteum/msm/types"

	"github.com/creatachain/creata-sdk/telemetry"
	sdk "github.com/creatachain/creata-sdk/types"
	"github.com/creatachain/creata-sdk/x/evidence/keeper"
	"github.com/creatachain/creata-sdk/x/evidence/types"
)

// BeginBlocker iterates through and handles any newly discovered evidence of
// misbehavior submitted by Augusteum. Currently, only equivocation is handled.
func BeginBlocker(ctx sdk.Context, req msm.RequestBeginBlock, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)

	for _, tmEvidence := range req.ByzantineValidators {
		switch tmEvidence.Type {
		// It's still ongoing discussion how should we treat and slash attacks with
		// premeditation. So for now we agree to treat them in the same way.
		case msm.EvidenceType_DUPLICATE_VOTE, msm.EvidenceType_LIGHT_CLIENT_ATTACK:
			evidence := types.FromMSMEvidence(tmEvidence)
			k.HandleEquivocationEvidence(ctx, evidence.(*types.Equivocation))

		default:
			k.Logger(ctx).Error(fmt.Sprintf("ignored unknown evidence type: %s", tmEvidence.Type))
		}
	}
}
