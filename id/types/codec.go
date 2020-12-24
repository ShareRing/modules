package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// ModuleCdc is the codec for the module
var ModuleCdc = codec.New()

func init() {
	RegisterCodec(ModuleCdc)
}

// RegisterCodec registers concrete types on the Amino codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgCreateId{}, "id/MsgCreateId", nil)
	cdc.RegisterConcrete(MsgCreateIdBatch{}, "id/MsgCreateIdBatch", nil)
	cdc.RegisterConcrete(MsgUpdateId{}, "id/MsgUpdateId", nil)
	cdc.RegisterConcrete(MsgReplaceIdOwner{}, "id/MsgReplaceIdOwner", nil)
}
