// Step 1
// Note: message for owner to assign orderid to transporter

package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// TODO: Describe your actions, these will implment the interface of `sdk.Msg`

// verify interface at compile time
var _ sdk.Msg = &MsgTransport{}

// MsgTransport - struct for unjailing jailed validator
type MsgTransport struct {
	Owner       sdk.AccAddress `json:"owner" yaml:"owner"`
	Transporter sdk.AccAddress `json:"transporter" yaml:"transporter"`
	OrderID     string         `json:"orderid" yaml:"orderid"`
}

// NewMsgTransport creates a new MsgTransport instance
func NewMsgTransport(owner sdk.AccAddress, transporter sdk.AccAddress, orderid string) MsgTransport {
	return MsgTransport{
		Owner:       owner,
		OrderID:     orderid,
		Transporter: transporter,
	}
}

const TransportConst = "Transport"

// nolint
func (msg MsgTransport) Route() string { return RouterKey }
func (msg MsgTransport) Type() string  { return TransportConst }
func (msg MsgTransport) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Owner)}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgTransport) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgTransport) ValidateBasic() error {
	if msg.Owner.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing owner address")
	}
	if msg.Transporter.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing transporter address")
	}
	return nil
}
