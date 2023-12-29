package expected

import (
	auctiontypes "github.com/MonikaCat/comdex/v13/x/auction/types"
	"github.com/MonikaCat/comdex/v13/x/liquidation/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	assettypes "github.com/MonikaCat/comdex/v13/x/asset/types"
	esmtypes "github.com/MonikaCat/comdex/v13/x/esm/types"
	markettypes "github.com/MonikaCat/comdex/v13/x/market/types"
)

type BankKeeper interface {
	BurnCoins(ctx sdk.Context, name string, coins sdk.Coins) error
	MintCoins(ctx sdk.Context, name string, coins sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, address sdk.AccAddress, name string, coins sdk.Coins) error
	SendCoinsFromModuleToAccount(ctx sdk.Context, name string, address sdk.AccAddress, coins sdk.Coins) error
	SpendableCoins(ctx sdk.Context, address sdk.AccAddress) sdk.Coins
	GetAllBalances(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	SendCoinsFromModuleToModule(
		ctx sdk.Context, senderModule, recipientModule string, amt sdk.Coins,
	) error
	SendCoins(ctx sdk.Context, fromAddr, toAddr sdk.AccAddress, amt sdk.Coins) error
	GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin
}

type AccountKeeper interface {
	GetModuleAddress(name string) sdk.AccAddress
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) authtypes.AccountI
}

type MarketKeeper interface {
	GetTwa(ctx sdk.Context, id uint64) (twa markettypes.TimeWeightedAverage, found bool)
	CalcAssetPrice(ctx sdk.Context, id uint64, amt sdk.Int) (price sdk.Dec, err error)
}

type BandOracleKeeper interface {
	GetOracleValidationResult(ctx sdk.Context) bool
}

type AssetKeeper interface {
	GetAsset(ctx sdk.Context, id uint64) (assettypes.Asset, bool)
	GetApp(ctx sdk.Context, id uint64) (assettypes.AppData, bool)
	SetApp(ctx sdk.Context, app assettypes.AppData)
	SetAppID(ctx sdk.Context, id uint64)
}

type EsmKeeper interface {
	GetKillSwitchData(ctx sdk.Context, appID uint64) (esmtypes.KillSwitchParams, bool)
}

type LiquidationKeeper interface {
	GetLockedVaultByApp(ctx sdk.Context, appID uint64) (lockedVault []types.LockedVault)
	GetLockedVault(ctx sdk.Context, appID, id uint64) (lockedVault types.LockedVault, found bool)
	DeleteLockedVault(ctx sdk.Context, appID, id uint64)
}

type AuctionKeeper interface {
	GetDutchLendAuctions(ctx sdk.Context, appID uint64) (auctions []auctiontypes.DutchAuction)
	SetHistoryDutchLendAuction(ctx sdk.Context, auction auctiontypes.DutchAuction) error
	DeleteDutchLendAuction(ctx sdk.Context, auction auctiontypes.DutchAuction) error
}
