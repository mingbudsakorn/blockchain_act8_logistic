// Step 6: create handler for MsgInitDeal
// Note: this file is use to initialize deal when receive MsgInitdeal

package logistic

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/earth2378/logistic/x/logistic/keeper"
	"github.com/earth2378/logistic/x/logistic/types"
)

func handlerMsgInitDeal(ctx sdk.Context, k keeper.Keeper, msg types.MsgInitDeal) (*sdk.Result, error) {
	// add msg to deal struct
	var deal = types.Deal{
		Owner:      msg.Owner,
		Customer:   msg.Customer,
		Price:      msg.Price,
		OrderID:    msg.OrderID,
		MaxTemp:    msg.MaxTemp,
		MinTemp:    msg.MinTemp,
		State:      types.Created,
		Cancelable: false,
	}

	// check if deal with provided orderid is exist or not
	_, err := k.GetDeal(ctx, msg.OrderID)
	if err == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Deal of this owner already exists")
	}

	// create deak
	k.SetDeal(ctx, deal)

	// set event (for logging transaction)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.EventTypeInitDeal),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Owner.String()),
			sdk.NewAttribute(types.AttributeOwner, msg.Owner.String()),
			sdk.NewAttribute(types.AttributeCustomer, msg.Customer.String()),
			sdk.NewAttribute(types.AttributePrice, msg.Price.String()),
			sdk.NewAttribute(types.AttributeOrderID, msg.OrderID),
			sdk.NewAttribute(types.AttributeMaxTemp, string(msg.MaxTemp)),
			sdk.NewAttribute(types.AttributeMinTemp, string(msg.MinTemp)),
		),
	)
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
