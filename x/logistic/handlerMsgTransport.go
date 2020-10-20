// Step 6: create handler for MsgTransport
// Note: this file is use to initialize deal when receive MsgTransport

package logistic

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/earth2378/logistic/x/logistic/keeper"
	"github.com/earth2378/logistic/x/logistic/types"
)

func handlerMsgTransport(ctx sdk.Context, k keeper.Keeper, msg types.MsgTransport) (*sdk.Result, error) {
	// get deal with orderid
	currentDeal, err := k.GetDeal(ctx, msg.OrderID)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Deal does not exists")
	}
	// check if caller is creater
	if msg.Owner.String() != currentDeal.Owner.String() {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Fail assign transporter, invalid owner")
	}
	// check if state is valid
	if currentDeal.State != types.Created {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Invalid state")
	}

	// set new transporter and state
	currentDeal.Transporter = msg.Transporter
	currentDeal.State = types.InTransit

	// update deal
	k.SetDeal(ctx, currentDeal)

	// set event (for logging transaction)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.EventTypeTransport),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Owner.String()),
			sdk.NewAttribute(types.AttributeOwner, msg.Owner.String()),
			sdk.NewAttribute(types.AttributeTransporter, msg.Transporter.String()),
			sdk.NewAttribute(types.AttributeOrderID, msg.OrderID),
		),
	)
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
