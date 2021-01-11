package xyzchain

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hirokryptor/xyzchain/x/xyzchain/keeper"
	"github.com/hirokryptor/xyzchain/x/xyzchain/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
        // this line is used by starport scaffolding # 1
case *types.MsgCreatePool:
	return handleMsgCreatePool(ctx, k, msg)

case *types.MsgUpdatePool:
	return handleMsgUpdatePool(ctx, k, msg)

case *types.MsgDeletePool:
	return handleMsgDeletePool(ctx, k, msg)

		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}
