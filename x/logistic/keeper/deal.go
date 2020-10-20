// Step 3: implement get and set deal function in keeper
// Note: add deal getter and setter function
//		 SetDeal is use when initDeal or update deal status
//		 GetDeal is use when someone want to get deal information by prodiving orderid

package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/earth2378/logistic/x/logistic/types"
)

// function to set deal when there are some mutation on deal to KVStore (both create and update)
func (k Keeper) SetDeal(ctx sdk.Context, deal types.Deal) {
	store := ctx.KVStore(k.storeKey)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(deal)
	key := []byte(types.DealPrefix + deal.OrderID)
	store.Set(key, value)
}

// function to get deal that store on KVStore
func (k Keeper) GetDeal(ctx sdk.Context, orderid string) (types.Deal, error) {
	store := ctx.KVStore(k.storeKey)
	var deal types.Deal
	byteKey := []byte(types.DealPrefix + orderid)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &deal)
	if err != nil {
		return deal, err
	}
	return deal, nil
}

// function to get deal with defined orderid
func getDeal(ctx sdk.Context, path []string, k Keeper) (res []byte, sdkError error) {
	orderid := path[0]
	deal, err := k.GetDeal(ctx, orderid)
	if err != nil {
		return nil, err
	}
	res, err = codec.MarshalJSONIndent(k.cdc, deal)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}
