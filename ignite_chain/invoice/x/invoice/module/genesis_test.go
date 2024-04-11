package invoice_test

import (
	"testing"

	keepertest "invoice/testutil/keeper"
	"invoice/testutil/nullify"
	invoice "invoice/x/invoice/module"
	"invoice/x/invoice/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		InvoiceList: []types.Invoice{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.InvoiceKeeper(t)
	invoice.InitGenesis(ctx, k, genesisState)
	got := invoice.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.InvoiceList, got.InvoiceList)
	// this line is used by starport scaffolding # genesis/test/assert
}
