package wasm

import (
	"github.com/CosmWasm/wasmd/x/wasm"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"

	assetkeeper "github.com/MonikaCat/comdex/v5/x/asset/keeper"
	auctionKeeper "github.com/MonikaCat/comdex/v5/x/auction/keeper"
	collectorKeeper "github.com/MonikaCat/comdex/v5/x/collector/keeper"
	esmKeeper "github.com/MonikaCat/comdex/v5/x/esm/keeper"
	lendKeeper "github.com/MonikaCat/comdex/v5/x/lend/keeper"
	liquidationKeeper "github.com/MonikaCat/comdex/v5/x/liquidation/keeper"
	liquidityKeeper "github.com/MonikaCat/comdex/v5/x/liquidity/keeper"
	lockerkeeper "github.com/MonikaCat/comdex/v5/x/locker/keeper"
	marketKeeper "github.com/MonikaCat/comdex/v5/x/market/keeper"
	rewardsKeeper "github.com/MonikaCat/comdex/v5/x/rewards/keeper"
	tokenMintkeeper "github.com/MonikaCat/comdex/v5/x/tokenmint/keeper"
	vaultKeeper "github.com/MonikaCat/comdex/v5/x/vault/keeper"
)

func RegisterCustomPlugins(
	locker *lockerkeeper.Keeper,
	tokenMint *tokenMintkeeper.Keeper,
	asset *assetkeeper.Keeper,
	rewards *rewardsKeeper.Keeper,
	collector *collectorKeeper.Keeper,
	liquidation *liquidationKeeper.Keeper,
	auction *auctionKeeper.Keeper,
	esm *esmKeeper.Keeper,
	vault *vaultKeeper.Keeper,
	lend *lendKeeper.Keeper,
	liquidity *liquidityKeeper.Keeper,
	market *marketKeeper.Keeper,
) []wasmkeeper.Option {
	comdexQueryPlugin := NewQueryPlugin(asset, locker, tokenMint, rewards, collector, liquidation, esm, vault, lend, liquidity, market)

	appDataQueryPluginOpt := wasmkeeper.WithQueryPlugins(&wasmkeeper.QueryPlugins{
		Custom: CustomQuerier(comdexQueryPlugin),
	})
	messengerDecoratorOpt := wasmkeeper.WithMessageHandlerDecorator(
		CustomMessageDecorator(*locker, *rewards, *asset, *collector, *liquidation, *auction, *tokenMint, *esm, *vault, *liquidity),
	)

	return []wasm.Option{
		appDataQueryPluginOpt,
		messengerDecoratorOpt,
	}
}
