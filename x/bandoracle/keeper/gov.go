package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/MonikaCat/comdex/v5/x/bandoracle/types"
)

func (k Keeper) HandleProposalFetchPrice(ctx sdk.Context, p *types.FetchPriceProposal) error {
	return k.AddFetchPriceRecords(ctx, p.FetchPrice)
}
