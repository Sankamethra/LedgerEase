package keeper_test

import (
	"context"
	"strconv"
	"testing"

	keepertest "invoice/testutil/keeper"
	"invoice/testutil/nullify"
	"invoice/x/invoice/keeper"
	"invoice/x/invoice/types"

	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNInvoice(keeper keeper.Keeper, ctx context.Context, n int) []types.Invoice {
	items := make([]types.Invoice, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetInvoice(ctx, items[i])
	}
	return items
}

func TestInvoiceGet(t *testing.T) {
	keeper, ctx := keepertest.InvoiceKeeper(t)
	items := createNInvoice(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetInvoice(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestInvoiceRemove(t *testing.T) {
	keeper, ctx := keepertest.InvoiceKeeper(t)
	items := createNInvoice(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveInvoice(ctx,
			item.Index,
		)
		_, found := keeper.GetInvoice(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestInvoiceGetAll(t *testing.T) {
	keeper, ctx := keepertest.InvoiceKeeper(t)
	items := createNInvoice(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllInvoice(ctx)),
	)
}
