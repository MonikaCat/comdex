package vault_test

import (
	"testing"

	"github.com/MonikaCat/comdex/v5/app"
	"github.com/MonikaCat/comdex/v5/x/vault"
	"github.com/MonikaCat/comdex/v5/x/vault/types"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	comdexApp := app.Setup(t, false)
	ctx := comdexApp.BaseApp.NewContext(false, tmproto.Header{})

	genesisState := types.GenesisState{
		Vaults:                      nil,
		StableMintVault:             nil,
		AppExtendedPairVaultMapping: nil,
		UserVaultAssetMapping:       nil,
		LengthOfVaults:              0,
	}

	vault.InitGenesis(ctx, comdexApp.VaultKeeper, &genesisState)
	got := vault.ExportGenesis(ctx, comdexApp.VaultKeeper)
	require.NotNil(t, got)
}
