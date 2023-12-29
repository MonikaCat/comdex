package liquidationsV2

import (
	"fmt"

	"github.com/MonikaCat/comdex/v5/x/liquidationsV2/keeper"
	"github.com/MonikaCat/comdex/v5/x/liquidationsV2/types"
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, ctx.BlockTime(), telemetry.MetricKeyBeginBlocker)

	err := k.Liquidate(ctx)
	if err != nil {
		ctx.Logger().Error("error in Liquidate function")
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeLiquidateErr,
				sdk.NewAttribute(types.Error, fmt.Sprintf("%s", err)),
			),
		)
	}
}
