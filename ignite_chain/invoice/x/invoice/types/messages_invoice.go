package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateInvoice{}

func NewMsgCreateInvoice(
	creator string,
	index string,
	Invoice_Number string,
	Customer_Name string,
	Invoice_Date string,
	Total_Amount string,
	Due_Date string,

) *MsgCreateInvoice {
	return &MsgCreateInvoice{
		Creator: creator,
		Index:   index,
		Invoice_Number: Invoice_Number,
        Customer_Name:   Customer_Name,
        Invoice_Date:   Invoice_Date,
        Total_Amount:    Total_Amount,
		Due_Date: Due_Date,
	}
}

func (msg *MsgCreateInvoice) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateInvoice{}

func NewMsgUpdateInvoice(
	creator string,
	index string,
	Invoice_Number string,
	Customer_Name string,
	Invoice_Date string,
	Total_Amount string,
	Due_Date string,

) *MsgUpdateInvoice {
	return &MsgUpdateInvoice{
		Creator: creator,
		Index:   index,
		Invoice_Number: Invoice_Number,
        Customer_Name:  Customer_Name,
        Invoice_Date:  Invoice_Date,
        Total_Amount:  Total_Amount,
		Due_Date: Due_Date,
	}
}

func (msg *MsgUpdateInvoice) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteInvoice{}

func NewMsgDeleteInvoice(
	creator string,
	index string,
	Invoice_Number string,
	Customer_Name string,
	Invoice_Date string,
	Total_Amount string,
	Due_Date string,

) *MsgDeleteInvoice {
	return &MsgDeleteInvoice{
		Creator: creator,
		Index:   index,
		Invoice_Number: Invoice_Number,
        Customer_Name:  Customer_Name,
        Invoice_Date:  Invoice_Date,
        Total_Amount:  Total_Amount,
		Due_Date: Due_Date,
	}
}

func (msg *MsgDeleteInvoice) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
