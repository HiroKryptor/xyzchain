package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
    cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
    sdk "github.com/cosmos/cosmos-sdk/types"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
    // this line is used by starport scaffolding # 2
cdc.RegisterConcrete(&MsgCreatePool{}, "xyzchain/CreatePool", nil)
cdc.RegisterConcrete(&MsgUpdatePool{}, "xyzchain/UpdatePool", nil)
cdc.RegisterConcrete(&MsgDeletePool{}, "xyzchain/DeletePool", nil)

} 

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
    // this line is used by starport scaffolding # 3
registry.RegisterImplementations((*sdk.Msg)(nil),
	&MsgCreatePool{},
	&MsgUpdatePool{},
	&MsgDeletePool{},
)
}

var (
	amino = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)
