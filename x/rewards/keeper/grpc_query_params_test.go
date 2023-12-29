package keeper_test

import (
	"testing"

	"github.com/MonikaCat/comdex/v13/app"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"

	"github.com/MonikaCat/comdex/v13/x/rewards/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestParamsQuery(t *testing.T) {
	comdexApp := app.Setup(t, false)
	ctx := comdexApp.BaseApp.NewContext(false, tmproto.Header{})

	wctx := sdk.WrapSDKContext(ctx)
	params := types.DefaultParams()
	comdexApp.Rewardskeeper.SetParams(ctx, params)

	response, err := comdexApp.Rewardskeeper.Params(wctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
}
