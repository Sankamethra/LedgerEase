package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgStoreinvoice{}

func NewMsgStoreinvoice(creator string, Invoice_Number string, Customer_Name string, Invoice_Date string, Total_Amount string, Due_Date string) *MsgStoreinvoice {
	return &MsgStoreinvoice{
		Creator: creator,
		Invoice_Number: Invoice_Number,
        Customer_Name: Customer_Name,
        Invoice_Date: Invoice_Date,
        Total_Amount: Total_Amount,
		Due_Date: Due_Date,
	}
}

func (msg *MsgStoreinvoice) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
