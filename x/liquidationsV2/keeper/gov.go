package keeper

import (
	"github.com/MonikaCat/comdex/v5/x/liquidationsV2/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) HandleWhitelistLiquidationProposal(ctx sdk.Context, p *types.WhitelistLiquidationProposal) error {
	return k.WhitelistLiquidation(ctx, p.Whitelisting)
}
