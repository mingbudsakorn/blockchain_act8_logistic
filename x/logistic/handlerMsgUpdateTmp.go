// Step 6: create handler for MsgUpdateTemp
// Note: this file is use to initialize deal when receive MsgUpdateTemp

package logistic

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/earth2378/logistic/x/logistic/keeper"
	"github.com/earth2378/logistic/x/logistic/types"
)

func handlerMsgUpdateTemp(ctx sdk.Context, k keeper.Keeper, msg types.MsgUpdateTemp) (*sdk.Result, error) {
	// get deal with orderid
	currentDeal, err := k.GetDeal(ctx, msg.OrderID)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Deal does not exists")
	}
	// check if caller is transporter
	if msg.Transporter.String() != currentDeal.Transporter.String() {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Fail to update temp, invalid transporter")
	}
	// check if state is valid
	if currentDeal.State != types.InTransit {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Invalid state")
	}
	// if temp not in range, deal id cancelable and update deal
	if msg.Temp > currentDeal.MaxTemp || msg.Temp < currentDeal.MinTemp {
		currentDeal.Cancelable = true
		k.SetDeal(ctx, currentDeal)
	}

	// set event (for logging transaction)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.EventTypeUpdateTemp),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Transporter.String()),
			sdk.NewAttribute(types.AttributeTransporter, msg.Transporter.String()),
			sdk.NewAttribute(types.AttributeUpdateTemp, string(msg.Temp)),
			sdk.NewAttribute(types.AttributeOrderID, msg.OrderID),
		),
	)
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
