package keeper

import (
	"github.com/comdex-official/comdex/x/auctionsV2/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	protobuftypes "github.com/gogo/protobuf/types"
)

func (k Keeper) SetAuctionID(ctx sdk.Context, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.AuctionIDKey
		value = k.cdc.MustMarshal(
			&protobuftypes.UInt64Value{
				Value: id,
			},
		)
	)
	store.Set(key, value)
}
func (k Keeper) GetAuctionID(ctx sdk.Context) uint64 {
	var (
		store = k.Store(ctx)
		key   = types.AuctionIDKey
		value = store.Get(key)
	)

	if value == nil {
		return 0
	}

	var id protobuftypes.UInt64Value
	k.cdc.MustUnmarshal(value, &id)

	return id.GetValue()
}

func (k Keeper) SetUserBidID(ctx sdk.Context, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.UserBidIDKey
		value = k.cdc.MustMarshal(
			&protobuftypes.UInt64Value{
				Value: id,
			},
		)
	)
	store.Set(key, value)
}
func (k Keeper) GetUserBidID(ctx sdk.Context) uint64 {
	var (
		store = k.Store(ctx)
		key   = types.UserBidIDKey
		value = store.Get(key)
	)

	if value == nil {
		return 0
	}

	var id protobuftypes.UInt64Value
	k.cdc.MustUnmarshal(value, &id)

	return id.GetValue()
}

func (k Keeper) SetAuction(ctx sdk.Context, auction types.Auction) error {

	var (
		store = k.Store(ctx)
		key   = types.AuctionKey(auction.AuctionId)
		value = k.cdc.MustMarshal(&auction)
	)

	store.Set(key, value)
	return nil
}
func (k Keeper) SetAuctionHistorical(ctx sdk.Context, auction types.AuctionHistorical) error {

	var (
		store = k.Store(ctx)
		key   = types.AuctionHistoricalKey(auction.AuctionId)
		value = k.cdc.MustMarshal(&auction)
	)

	store.Set(key, value)
	return nil
}

func (k Keeper) SetUserBid(ctx sdk.Context, userBid types.Bid) error {

	var (
		store = k.Store(ctx)
		key   = types.UserBidKey(userBid.BiddingId)
		value = k.cdc.MustMarshal(&userBid)
	)

	store.Set(key, value)
	return nil
}

// func (k Keeper) AddAuctionParams(ctx sdk.Context, liquidationData liquidationtypes.LockedVault,auctionID uint64) (auction types.Auctions, err error) {

// 	auctionData := types.Auctions{
// 		AuctionId:                   auctionID + 1,
// 		CollateralToken:             liquidationData.CollateralToken,
// 		DebtToken:                   liquidationData.TargetDebt,
// 		CollateralTokenAuctionPrice: CollateralTokenInitialPrice,
// 		CollateralTokenOraclePrice:  sdk.NewDecFromInt(sdk.NewInt(int64(twaDataCollateral.Twa))),
// 		DebtTokenOraclePrice:        sdk.NewDecFromInt(sdk.NewInt(int64(twaDataDebt.Twa))),
// 		LockedVaultId:               liquidationData.LockedVaultId,
// 		StartTime:                   ctx.BlockTime(),
// 		EndTime:                     ctx.BlockTime().Add(time.Second * time.Duration(auctionParams.AuctionDurationSeconds)),
// 		AppId:                       liquidationData.AppId,
// 		AuctionType:                 liquidationData.AuctionType,
// 	}

// 	err := k.SetAuction(ctx, auctionData)
// 	if err != nil {
// 		return auction, err
// 	}

// 	return auctionData, nil

// }

func (k Keeper) DeleteAuction(ctx sdk.Context, auction types.Auction) error {

	var (
		store = k.Store(ctx)
		key   = types.AuctionKey(auction.AuctionId)
	)
	store.Delete(key)
	return nil
}

func (k Keeper) GetUserBid(ctx sdk.Context, userBidId uint64) (userBid types.Bid, err error) {
	var (
		store = k.Store(ctx)
		key   = types.UserBidKey(userBidId)
		value = store.Get(key)
	)

	if value == nil {
		return userBid, sdkerrors.ErrNotFound
	}

	k.cdc.MustUnmarshal(value, &userBid)
	return userBid, nil
}

func (k Keeper) GetAuction(ctx sdk.Context, auctionID uint64) (auction types.Auction, err error) {
	var (
		store = k.Store(ctx)
		key   = types.AuctionKey(auctionID)
		value = store.Get(key)
	)

	if value == nil {
		return auction, sdkerrors.ErrNotFound
	}

	k.cdc.MustUnmarshal(value, &auction)
	return auction, nil
}

func (k Keeper) GetAuctionHistorical(ctx sdk.Context, auctionID uint64) (auction types.AuctionHistorical, err error) {
	var (
		store = k.Store(ctx)
		key   = types.AuctionHistoricalKey(auctionID)
		value = store.Get(key)
	)

	if value == nil {
		return auction, sdkerrors.ErrNotFound
	}

	k.cdc.MustUnmarshal(value, &auction)
	return auction, nil
}

func (k Keeper) GetAuctions(ctx sdk.Context) (auctions []types.Auction) {
	var (
		store = k.Store(ctx)
		iter  = sdk.KVStorePrefixIterator(store, types.AuctionKeyPrefix)
	)

	defer func(iter sdk.Iterator) {
		err := iter.Close()
		if err != nil {
			return
		}
	}(iter)

	for ; iter.Valid(); iter.Next() {
		var auction types.Auction
		k.cdc.MustUnmarshal(iter.Value(), &auction)
		auctions = append(auctions, auction)
	}

	return auctions
}
func (k Keeper) GetAuctionHistoricals(ctx sdk.Context) (auctions []types.AuctionHistorical) {
	var (
		store = k.Store(ctx)
		iter  = sdk.KVStorePrefixIterator(store, types.AuctionHistoricalKeyPrefix)
	)

	defer func(iter sdk.Iterator) {
		err := iter.Close()
		if err != nil {
			return
		}
	}(iter)

	for ; iter.Valid(); iter.Next() {
		var auction types.AuctionHistorical
		k.cdc.MustUnmarshal(iter.Value(), &auction)
		auctions = append(auctions, auction)
	}

	return auctions
}

func (k Keeper) GetUserBids(ctx sdk.Context) (userBids []types.Bid) {
	var (
		store = k.Store(ctx)
		iter  = sdk.KVStorePrefixIterator(store, types.UserBidKeyPrefix)
	)

	defer func(iter sdk.Iterator) {
		err := iter.Close()
		if err != nil {
			return
		}
	}(iter)

	for ; iter.Valid(); iter.Next() {
		var userBid types.Bid
		k.cdc.MustUnmarshal(iter.Value(), &userBid)
		userBids = append(userBids, userBid)
	}

	return userBids
}
