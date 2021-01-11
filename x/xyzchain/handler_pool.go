package xyzchain

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
    sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/hirokryptor/xyzchain/x/xyzchain/types"
	"github.com/hirokryptor/xyzchain/x/xyzchain/keeper"
)

func handleMsgCreatePool(ctx sdk.Context, k keeper.Keeper, msg *types.MsgCreatePool) (*sdk.Result, error) {
	k.CreatePool(ctx, *msg)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgUpdatePool(ctx sdk.Context, k keeper.Keeper, msg *types.MsgUpdatePool) (*sdk.Result, error) {
	var pool = types.Pool{
		Creator: msg.Creator,
		Id:      msg.Id,
    	Token1: msg.Token1,
    	Token2: msg.Token2,
    	Total1: msg.Total1,
    	Total2: msg.Total2,
	}

    if msg.Creator != k.GetPoolOwner(ctx, msg.Id) { // Checks if the the msg sender is the same as the current owner
        return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect owner") // If not, throw an error                                                                                             
    }          

	k.UpdatePool(ctx, pool)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgDeletePool(ctx sdk.Context, k keeper.Keeper, msg *types.MsgDeletePool) (*sdk.Result, error) {
    if !k.HasPool(ctx, msg.Id) {                                                                                                                                                                    
        return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, msg.Id)                                                                                                                                
    }                                                                                                                                                                                                  
    if msg.Creator != k.GetPoolOwner(ctx, msg.Id) {
        return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect owner")                                                                                                                       
    } 

	k.DeletePool(ctx, msg.Id)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
