package keeper

import (
	"context"

	"invoice/x/invoice/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) InvoiceAll(ctx context.Context, req *types.QueryAllInvoiceRequest) (*types.QueryAllInvoiceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var invoices []types.Invoice

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	invoiceStore := prefix.NewStore(store, types.KeyPrefix(types.InvoiceKeyPrefix))

	pageRes, err := query.Paginate(invoiceStore, req.Pagination, func(key []byte, value []byte) error {
		var invoice types.Invoice
		if err := k.cdc.Unmarshal(value, &invoice); err != nil {
			return err
		}

		invoices = append(invoices, invoice)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllInvoiceResponse{Invoice: invoices, Pagination: pageRes}, nil
}

func (k Keeper) Invoice(ctx context.Context, req *types.QueryGetInvoiceRequest) (*types.QueryGetInvoiceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, found := k.GetInvoice(
		ctx,
		req.Creator,
		req.Index,
		req.Invoice_Number,
		req.Customer_Name,
		req.Invoice_Date,
		req.Total_Amount,
		req.Due_Date,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetInvoiceResponse{Invoice: val}, nil
}
