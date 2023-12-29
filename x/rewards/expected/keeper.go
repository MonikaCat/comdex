package expected

import (
	lendtypes "github.com/MonikaCat/comdex/v5/x/lend/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	assettypes "github.com/MonikaCat/comdex/v5/x/asset/types"
	collectortypes "github.com/MonikaCat/comdex/v5/x/collector/types"
	esmtypes "github.com/MonikaCat/comdex/v5/x/esm/types"
	liquiditytypes "github.com/MonikaCat/comdex/v5/x/liquidity/types"
	lockertypes "github.com/MonikaCat/comdex/v5/x/locker/types"
	markettypes "github.com/MonikaCat/comdex/v5/x/market/types"
	"github.com/MonikaCat/comdex/v5/x/rewards/types"
	vaulttypes "github.com/MonikaCat/comdex/v5/x/vault/types"
)

// AccountKeeper defines the expected account keeper used for simulations (noalias).
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) authtypes.AccountI
	// Methods imported from account should be defined here
}

type LiquidityKeeper interface {
	GetPair(ctx sdk.Context, appID, id uint64) (pair liquiditytypes.Pair, found bool)
	GetPool(ctx sdk.Context, appID, id uint64) (pool liquiditytypes.Pool, found bool)
	GetFarmingRewardsData(ctx sdk.Context, appID uint64, coinToDistribute sdk.Coin, liquidityGaugeData types.LiquidtyGaugeMetaData) ([]types.RewardDistributionDataCollector, error)
	TransferFundsForSwapFeeDistribution(ctx sdk.Context, appID, poolID uint64) (sdk.Coin, error)
	GetActiveFarmer(ctx sdk.Context, appID, poolID uint64, farmer sdk.AccAddress) (activeFarmer liquiditytypes.ActiveFarmer, found bool)
	GetPoolTokenDesrializerKit(ctx sdk.Context, appID, poolID uint64) (liquiditytypes.PoolTokenDeserializerKit, error)
	CalculateXYFromPoolCoin(ctx sdk.Context, deserializerKit liquiditytypes.PoolTokenDeserializerKit, poolCoin sdk.Coin) (sdk.Int, sdk.Int, error)
	GetQueuedFarmer(ctx sdk.Context, appID, poolID uint64, farmer sdk.AccAddress) (queuedFarmer liquiditytypes.QueuedFarmer, found bool)
	GetAmountFarmedForAssetID(ctx sdk.Context, appID, assetID uint64, farmer sdk.AccAddress) (sdk.Int, error)
}

type AssetKeeper interface {
	GetPairsVault(ctx sdk.Context, id uint64) (pairs assettypes.ExtendedPairVault, found bool)
	HasAssetForDenom(ctx sdk.Context, denom string) bool
	GetAssetForDenom(ctx sdk.Context, denom string) (asset assettypes.Asset, found bool)
	GetAsset(ctx sdk.Context, id uint64) (assettypes.Asset, bool)
	GetApp(ctx sdk.Context, id uint64) (app assettypes.AppData, found bool)
	GetApps(ctx sdk.Context) (apps []assettypes.AppData, found bool)
	GetPair(ctx sdk.Context, id uint64) (pair assettypes.Pair, found bool)
	GetPairsVaults(ctx sdk.Context) (apps []assettypes.ExtendedPairVault, found bool)
}

type MarketKeeper interface {
	CalcAssetPrice(ctx sdk.Context, id uint64, amt sdk.Int) (price sdk.Dec, err error)
	GetTwa(ctx sdk.Context, id uint64) (twa markettypes.TimeWeightedAverage, found bool)
}

type LockerKeeper interface {
	GetLockerProductAssetMapping(ctx sdk.Context, appMappingID, assetID uint64) (lockerProductMapping lockertypes.LockerProductAssetMapping, found bool)
	GetLocker(ctx sdk.Context, lockerID uint64) (locker lockertypes.Locker, found bool)
	GetLockers(ctx sdk.Context) (locker []lockertypes.Locker)
	GetLockerLookupTableByApp(ctx sdk.Context, appID uint64) (lockerLookupData []lockertypes.LockerLookupTableData, found bool)
	GetLockerLookupTable(ctx sdk.Context, appID, assetID uint64) (lockerLookupData lockertypes.LockerLookupTableData, found bool)
	SetLocker(ctx sdk.Context, locker lockertypes.Locker)
	SetLockerTotalRewardsByAssetAppWise(ctx sdk.Context, lockerRewardsMapping lockertypes.LockerTotalRewardsByAssetAppWise) error
	GetLockerTotalRewardsByAssetAppWise(ctx sdk.Context, appID, assetID uint64) (lockerRewardsMapping lockertypes.LockerTotalRewardsByAssetAppWise, found bool)
	SetLockerLookupTable(ctx sdk.Context, lockerLookupData lockertypes.LockerLookupTableData)
	GetUserLockerAssetMapping(ctx sdk.Context, address string, appID, assetID uint64) (userLockerAssetData lockertypes.UserAppAssetLockerMapping, found bool)
}

type CollectorKeeper interface {
	GetAppidToAssetCollectorMapping(ctx sdk.Context, appID, assetID uint64) (appAssetCollectorData collectortypes.AppToAssetIdCollectorMapping, found bool)
	GetCollectorLookupTable(ctx sdk.Context, appID, assetID uint64) (collectorLookup collectortypes.CollectorLookupTableData, found bool)
	GetAppToDenomsMapping(ctx sdk.Context, appID uint64) (appToDenom collectortypes.AppToDenomsMapping, found bool)
	GetNetFeeCollectedData(ctx sdk.Context, appID, assetID uint64) (netFeeData collectortypes.AppAssetIdToFeeCollectedData, found bool)
	SetNetFeeCollectedData(ctx sdk.Context, appID, assetID uint64, fee sdk.Int) error
	DecreaseNetFeeCollectedData(ctx sdk.Context, appID, assetID uint64, amount sdk.Int) error
}

type VaultKeeper interface {
	GetAppMappingData(ctx sdk.Context, appMappingID uint64) (appExtendedPairVaultData []vaulttypes.AppExtendedPairVaultMappingData, found bool)
	CalculateCollateralizationRatio(ctx sdk.Context, extendedPairVaultID uint64, amountIn sdk.Int, amountOut sdk.Int) (sdk.Dec, error)
	GetVault(ctx sdk.Context, id uint64) (vault vaulttypes.Vault, found bool)
	DeleteVault(ctx sdk.Context, id uint64)
	UpdateAppExtendedPairVaultMappingDataOnMsgCreate(ctx sdk.Context, vaultData vaulttypes.Vault)
	UpdateCollateralLockedAmountLockerMapping(ctx sdk.Context, appMappingID uint64, extendedPairID uint64, amount sdk.Int, changeType bool)
	UpdateTokenMintedAmountLockerMapping(ctx sdk.Context, appMappingID uint64, extendedPairID uint64, amount sdk.Int, changeType bool)
	DeleteUserVaultExtendedPairMapping(ctx sdk.Context, address string, appID uint64, pairVaultID uint64)
	DeleteAddressFromAppExtendedPairVaultMapping(ctx sdk.Context, extendedPairID uint64, userVaultID uint64, appMappingID uint64)
	SetVault(ctx sdk.Context, vault vaulttypes.Vault)
	GetAppExtendedPairVaultMappingData(ctx sdk.Context, appMappingID uint64, pairVaultID uint64) (appExtendedPairVaultData vaulttypes.AppExtendedPairVaultMappingData, found bool)
	GetStableMintVaultUserRewards(ctx sdk.Context, appID uint64, user string) (mappingData []vaulttypes.StableMintVaultRewards, found bool)
	DeleteStableMintVaultRewards(ctx sdk.Context, stableMintVaultRewards vaulttypes.StableMintVaultRewards)
	SetStableMintVaultRewards(ctx sdk.Context, stableMintVaultRewards vaulttypes.StableMintVaultRewards)
	GetStableMintVaultRewards(ctx sdk.Context, stableMintVaultRewards vaulttypes.StableMintVaultRewards) (mappingData vaulttypes.StableMintVaultRewards, found bool)
	GetStableMintVaultRewardsByApp(ctx sdk.Context, appID uint64) (mappingData []vaulttypes.StableMintVaultRewards, found bool)
	GetStableMintVaultRewardsOfAllApps(ctx sdk.Context) (mappingData []vaulttypes.StableMintVaultRewards)
}

type BankKeeper interface {
	BurnCoins(ctx sdk.Context, name string, coins sdk.Coins) error
	MintCoins(ctx sdk.Context, name string, coins sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, address sdk.AccAddress, name string, coins sdk.Coins) error
	SendCoinsFromModuleToAccount(ctx sdk.Context, name string, address sdk.AccAddress, coins sdk.Coins) error

	SendCoinsFromModuleToModule(
		ctx sdk.Context, senderModule, recipientModule string, amt sdk.Coins,
	) error

	SpendableCoins(ctx sdk.Context, address sdk.AccAddress) sdk.Coins
	GetSupply(ctx sdk.Context, denom string) sdk.Coin
	GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin
}

type EsmKeeper interface {
	GetKillSwitchData(ctx sdk.Context, appID uint64) (esmtypes.KillSwitchParams, bool)
	GetESMStatus(ctx sdk.Context, id uint64) (esmStatus esmtypes.ESMStatus, found bool)
}

type LendKeeper interface {
	GetBorrow(ctx sdk.Context, id uint64) (borrow lendtypes.BorrowAsset, found bool)
	GetLend(ctx sdk.Context, id uint64) (lend lendtypes.LendAsset, found bool)
	GetAssetStatsByPoolIDAndAssetID(ctx sdk.Context, poolID, assetID uint64) (PoolAssetLBMapping lendtypes.PoolAssetLBMapping, found bool)
	GetLendPair(ctx sdk.Context, id uint64) (pair lendtypes.Extended_Pair, found bool)
	GetPool(ctx sdk.Context, id uint64) (pool lendtypes.Pool, found bool)
	UserAssetLends(ctx sdk.Context, addr string, assetID uint64) (sdk.Int, bool)
}
