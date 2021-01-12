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
	cdc.RegisterConcrete(MsgCreateDoc{}, "id/MsgCreateDoc", nil)
	cdc.RegisterConcrete(MsgCreateDocBatch{}, "id/MsgCreateDocBatch", nil)
	cdc.RegisterConcrete(MsgUpdateDoc{}, "id/MsgUpdateDoc", nil)
	cdc.RegisterConcrete(MsgRevokeDoc{}, "id/MsgRevokeDoc", nil)
}
