// Step 6: create handler for MsgReceive
// Note: this file is use to initialize deal when receive MsgReceive

package logistic

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/earth2378/logistic/x/logistic/keeper"
	"github.com/earth2378/logistic/x/logistic/types"
)

func handlerMsgReceive(ctx sdk.Context, k keeper.Keeper, msg types.MsgReceive) (*sdk.Result, error) {
	// get deal with orderid
	currentDeal, err := k.GetDeal(ctx, msg.OrderID)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Deal does not exists")
	}
	// check if caller is customer
	if msg.Customer.String() != currentDeal.Customer.String() {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Fail receive, invalid customer")
	}
	// check if state is valid
	if currentDeal.State != types.InTransit {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Invalid state")
	}

	// send coin to owner (make a payment)
	sdkError := k.CoinKeeper.SendCoins(ctx, currentDeal.Customer, currentDeal.Owner, currentDeal.Price)
	if sdkError != nil {
		return nil, sdkError
	}

	// set state to complete and update deal
	currentDeal.State = types.Complete
	k.SetDeal(ctx, currentDeal)

	// set event (for logging transaction)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.EventTypeReceive),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Customer.String()),
			sdk.NewAttribute(types.AttributeOrderID, msg.OrderID),
		),
	)
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
