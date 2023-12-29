package collector_test

import (
	"testing"

	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"

	app "github.com/MonikaCat/comdex/v13/app"
	"github.com/MonikaCat/comdex/v13/x/collector"
	"github.com/MonikaCat/comdex/v13/x/collector/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	comdexApp := app.Setup(t, false)
	ctx := comdexApp.BaseApp.NewContext(false, tmproto.Header{})
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
	}

	k := comdexApp.CollectorKeeper
	collector.InitGenesis(ctx, k, &genesisState)
	got := collector.ExportGenesis(ctx, k)
	require.NotNil(t, got)
}
