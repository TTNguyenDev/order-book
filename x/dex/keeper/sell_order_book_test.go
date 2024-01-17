package keeper_test

import (
	"context"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "orderbook-interchange/testutil/keeper"
	"orderbook-interchange/testutil/nullify"
	"orderbook-interchange/x/dex/keeper"
	"orderbook-interchange/x/dex/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNSellOrderBook(keeper keeper.Keeper, ctx context.Context, n int) []types.SellOrderBook {
	items := make([]types.SellOrderBook, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetSellOrderBook(ctx, items[i])
	}
	return items
}

func TestSellOrderBookGet(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNSellOrderBook(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetSellOrderBook(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestSellOrderBookRemove(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNSellOrderBook(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveSellOrderBook(ctx,
			item.Index,
		)
		_, found := keeper.GetSellOrderBook(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestSellOrderBookGetAll(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNSellOrderBook(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllSellOrderBook(ctx)),
	)
}
