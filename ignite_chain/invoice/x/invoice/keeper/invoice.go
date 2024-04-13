package keeper

import (
	"context"

	"invoice/x/invoice/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetInvoice set a specific invoice in the store from its index
func (k Keeper) SetInvoice(ctx context.Context, invoice types.Invoice) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.InvoiceKeyPrefix))
	b := k.cdc.MustMarshal(&invoice)
	store.Set(types.InvoiceKey(
		invoice.Index,
		// invoice.InvoiceId,
		// invoice.Company,
		// invoice.DueDate,
		// invoice.Amount,
	), b)
}

// GetInvoice returns a invoice from its index
func (k Keeper) GetInvoice(
	ctx context.Context,
	creator string,
	index string,
	Invoice_Number string,
	Customer_Name string,
	Invoice_Date string,
	Total_Amount string,
	Due_Date string,


) (val types.Invoice, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.InvoiceKeyPrefix))

	b := store.Get(types.InvoiceKey(
		index,
		// InvoiceId,
		// Company,
		// DueDate,
		// Amount,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveInvoice removes a specific invoice from the store based on its index
func (k Keeper) RemoveInvoice(
	ctx context.Context,
	creator string,
	index string,
	Invoice_Number string,
	Customer_Name string,
	Invoice_Date string,
	Total_Amount string,
	Due_Date string,
) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.InvoiceKeyPrefix))
	key := types.InvoiceKey(index) // Assuming 'index' is the correct parameter to uniquely identify an invoice
	store.Delete(key)
}


// GetAllInvoice returns all invoice
func (k Keeper) GetAllInvoice(ctx context.Context) (list []types.Invoice) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.InvoiceKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Invoice
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
