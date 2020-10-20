// Step 1
// Note: Messag for transporter to update current temp of product with orderid

package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// TODO: Describe your actions, these will implment the interface of `sdk.Msg`

// verify interface at compile time
var _ sdk.Msg = &MsgUpdateTemp{}

// MsgUpdateTemp - struct for unjailing jailed validator
type MsgUpdateTemp struct {
	Transporter sdk.AccAddress `json:"transporter" yaml:"transporter"`
	OrderID     string         `json:"orderid" yaml:"orderid"`
	Temp        int            `json:"temp" yaml:"temp"`
}

// NewMsgUpdateTemp creates a new MsgUpdateTemp instance
func NewMsgUpdateTemp(transporter sdk.AccAddress, orderid string, temp int) MsgUpdateTemp {
	return MsgUpdateTemp{
		Transporter: transporter,
		OrderID:     orderid,
		Temp:        temp,
	}
}

const UpdateTempConst = "UpdateTemp"

// nolint
func (msg MsgUpdateTemp) Route() string { return RouterKey }
func (msg MsgUpdateTemp) Type() string  { return UpdateTempConst }
func (msg MsgUpdateTemp) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Transporter)}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgUpdateTemp) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgUpdateTemp) ValidateBasic() error {
	if msg.Transporter.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing transporter address")
	}
	return nil
}
