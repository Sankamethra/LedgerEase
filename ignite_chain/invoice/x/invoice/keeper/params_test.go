package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "invoice/testutil/keeper"
	"invoice/x/invoice/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.InvoiceKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
