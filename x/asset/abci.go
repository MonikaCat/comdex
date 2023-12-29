package asset

import (
	abci "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	utils "github.com/MonikaCat/comdex/v5/types"
	"github.com/MonikaCat/comdex/v5/x/asset/keeper"
)

func BeginBlocker(ctx sdk.Context, _ abci.RequestBeginBlock, k keeper.Keeper) {
	_ = utils.ApplyFuncIfNoError(ctx, func(ctx sdk.Context) error {
		return nil
	})
}
