package keeper

import (
	"context"

	"invoice/x/invoice/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateInvoice(goCtx context.Context, msg *types.MsgCreateInvoice) (*types.MsgCreateInvoiceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetInvoice(
		ctx,
		msg.Creator,
		msg.Index,
		msg.Invoice_Number,
		msg.Customer_Name,
		msg.Invoice_Date,
		msg.Total_Amount,
		msg.Due_Date,
	)
	if isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var invoice = types.Invoice{
		Creator       : msg.Creator,
		Invoice_Number: msg.Invoice_Number,
        Customer_Name : msg.Customer_Name,
        Invoice_Date  : msg.Invoice_Date,
        Total_Amount  : msg.Total_Amount,
		Due_Date      : msg.Due_Date,
	}

	k.SetInvoice(
		ctx,
		invoice,
	)
	return &types.MsgCreateInvoiceResponse{}, nil
}

func (k msgServer) UpdateInvoice(goCtx context.Context, msg *types.MsgUpdateInvoice) (*types.MsgUpdateInvoiceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetInvoice(
		ctx,
		msg.Creator,
		msg.Index,
		msg.Invoice_Number,
		msg.Customer_Name,
		msg.Invoice_Date,
		msg.Total_Amount,
		msg.Due_Date,
	)
	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var invoice = types.Invoice{
		Creator: msg.Creator,
		Index:   msg.Index,
		Invoice_Number: msg.Invoice_Number,
        Customer_Name : msg.Customer_Name,
        Invoice_Date  : msg.Invoice_Date,
        Total_Amount  : msg.Total_Amount,
		Due_Date      : msg.Due_Date,
	}

	k.SetInvoice(ctx, invoice)

	return &types.MsgUpdateInvoiceResponse{}, nil
}

func (k msgServer) DeleteInvoice(goCtx context.Context, msg *types.MsgDeleteInvoice) (*types.MsgDeleteInvoiceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetInvoice(
		ctx,
		msg.Creator,
		msg.Index,
		msg.Invoice_Number,
		msg.Customer_Name,
		msg.Invoice_Date,
		msg.Total_Amount,
		msg.Due_Date,
	)
	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveInvoice(
		ctx,
		msg.Creator,
		msg.Index,
		msg.Invoice_Number,
		msg.Customer_Name,
		msg.Invoice_Date,
		msg.Total_Amount,
		msg.Due_Date,
	)

	return &types.MsgDeleteInvoiceResponse{}, nil
}
