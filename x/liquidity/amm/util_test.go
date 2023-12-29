package amm_test

import (
	"testing"

	sdkmath "cosmossdk.io/math"
	"github.com/stretchr/testify/require"

	utils "github.com/MonikaCat/comdex/v5/types"
	"github.com/MonikaCat/comdex/v5/x/liquidity/amm"
)

func TestMatchableAmount(t *testing.T) {
	order1 := newOrder(amm.Buy, utils.ParseDec("1.0"), sdkmath.NewInt(10000))
	for _, tc := range []struct {
		order    amm.Order
		price    sdkmath.LegacyDec
		expected sdkmath.Int
	}{
		{order1, utils.ParseDec("1"), sdkmath.NewInt(10000)},
		{order1, utils.ParseDec("0.01"), sdkmath.NewInt(10000)},
		{order1, utils.ParseDec("100"), sdkmath.NewInt(100)},
		{order1, utils.ParseDec("100.1"), sdkmath.NewInt(99)},
		{order1, utils.ParseDec("9999"), sdkmath.NewInt(1)},
		{order1, utils.ParseDec("10001"), sdkmath.NewInt(0)},
	} {
		t.Run("", func(t *testing.T) {
			require.True(sdkmath.IntEq(t, tc.expected, amm.MatchableAmount(tc.order, tc.price)))
		})
	}
}

type batchIdOrderer struct {
	batchId uint64
}

func (orderer *batchIdOrderer) Order(dir amm.OrderDirection, price sdkmath.LegacyDec, amt sdkmath.Int) amm.Order {
	return &batchIdOrder{newOrder(dir, price, amt), orderer.batchId}
}

type batchIdOrder struct {
	amm.Order
	batchId uint64
}

func (order *batchIdOrder) GetBatchID() uint64 {
	return order.batchId
}

func TestGroupOrdersByBatchId(t *testing.T) {
	price := utils.ParseDec("1.0")
	newOrder := func(amt sdkmath.Int, batchId uint64) amm.Order {
		return (&batchIdOrderer{batchId}).Order(amm.Buy, price, amt)
	}
	orders := []amm.Order{
		newOrder(sdkmath.NewInt(32000), 0),
		newOrder(sdkmath.NewInt(8000), 4),
		newOrder(sdkmath.NewInt(1000), 1),
		newOrder(sdkmath.NewInt(16000), 4),
		newOrder(sdkmath.NewInt(4000), 2),
		newOrder(sdkmath.NewInt(2000), 1),
		newOrder(sdkmath.NewInt(64000), 0),
	}
	groups := amm.GroupOrdersByBatchID(orders)
	require.EqualValues(t, 1, groups[0].BatchID)
	require.True(sdkmath.IntEq(t, sdkmath.NewInt(3000), amm.TotalAmount(groups[0].Orders)))
	require.EqualValues(t, 2, groups[1].BatchID)
	require.True(sdkmath.IntEq(t, sdkmath.NewInt(4000), amm.TotalAmount(groups[1].Orders)))
	require.EqualValues(t, 4, groups[2].BatchID)
	require.True(sdkmath.IntEq(t, sdkmath.NewInt(24000), amm.TotalAmount(groups[2].Orders)))
	require.EqualValues(t, 0, groups[3].BatchID)
	require.True(sdkmath.IntEq(t, sdkmath.NewInt(96000), amm.TotalAmount(groups[3].Orders)))
}
