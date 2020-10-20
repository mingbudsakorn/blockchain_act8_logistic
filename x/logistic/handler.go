// Note: this function is mapping Msg request to the handler

package logistic

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/earth2378/logistic/x/logistic/keeper"
	"github.com/earth2378/logistic/x/logistic/types"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		// Step 7: register handler functions to main handler
		case types.MsgInitDeal:
			return handlerMsgInitDeal(ctx, k, msg)
		case types.MsgTransport:
			return handlerMsgTransport(ctx, k, msg)
		case types.MsgUpdateTemp:
			return handlerMsgUpdateTemp(ctx, k, msg)
		case types.MsgReceive:
			return handlerMsgReceive(ctx, k, msg)
		case types.MsgReject:
			return handlerMsgReject(ctx, k, msg)

		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}
