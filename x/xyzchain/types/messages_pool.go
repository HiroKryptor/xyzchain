package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreatePool{}

func NewMsgCreatePool(creator string, token1 string, token2 string, total1 int32, total2 int32) *MsgCreatePool {
  return &MsgCreatePool{
		Creator: creator,
    Token1: token1,
    Token2: token2,
    Total1: total1,
    Total2: total2,
	}
}

func (msg *MsgCreatePool) Route() string {
  return RouterKey
}

func (msg *MsgCreatePool) Type() string {
  return "CreatePool"
}

func (msg *MsgCreatePool) GetSigners() []sdk.AccAddress {
  creator, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    panic(err)
  }
  return []sdk.AccAddress{creator}
}

func (msg *MsgCreatePool) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgCreatePool) ValidateBasic() error {
  _, err := sdk.AccAddressFromBech32(msg.Creator)
  	if err != nil {
  		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
  	}
  return nil
}

var _ sdk.Msg = &MsgUpdatePool{}

func NewMsgUpdatePool(creator string, id string, token1 string, token2 string, total1 int32, total2 int32) *MsgUpdatePool {
  return &MsgUpdatePool{
        Id: id,
		Creator: creator,
    Token1: token1,
    Token2: token2,
    Total1: total1,
    Total2: total2,
	}
}

func (msg *MsgUpdatePool) Route() string {
  return RouterKey
}

func (msg *MsgUpdatePool) Type() string {
  return "UpdatePool"
}

func (msg *MsgUpdatePool) GetSigners() []sdk.AccAddress {
  creator, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    panic(err)
  }
  return []sdk.AccAddress{creator}
}

func (msg *MsgUpdatePool) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdatePool) ValidateBasic() error {
  _, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
  }
   return nil
}

var _ sdk.Msg = &MsgCreatePool{}

func NewMsgDeletePool(creator string, id string) *MsgDeletePool {
  return &MsgDeletePool{
        Id: id,
		Creator: creator,
	}
} 
func (msg *MsgDeletePool) Route() string {
  return RouterKey
}

func (msg *MsgDeletePool) Type() string {
  return "DeletePool"
}

func (msg *MsgDeletePool) GetSigners() []sdk.AccAddress {
  creator, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    panic(err)
  }
  return []sdk.AccAddress{creator}
}

func (msg *MsgDeletePool) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgDeletePool) ValidateBasic() error {
  _, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
  }
  return nil
}
