package invoice

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"invoice/x/invoice/keeper"
	"invoice/x/invoice/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the invoice
	for _, elem := range genState.InvoiceList {
		k.SetInvoice(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	if err := k.SetParams(ctx, genState.Params); err != nil {
		panic(err)
	}
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.InvoiceList = k.GetAllInvoice(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
