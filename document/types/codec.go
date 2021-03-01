package types

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
)

// ModuleCdc is the codec for the module
var ModuleCdc = codec.New()

func init() {
	RegisterCodec(ModuleCdc)
}

// RegisterCodec registers concrete types on the Amino codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgCreateDoc{}, fmt.Sprintf("%s/%s", ModuleName, TypeMsgCreateDoc), nil)
	cdc.RegisterConcrete(MsgCreateDocBatch{}, fmt.Sprintf("%s/%s", ModuleName, TypeMsgCreateDocInBatch), nil)
	cdc.RegisterConcrete(MsgUpdateDoc{}, fmt.Sprintf("%s/%s", ModuleName, TypeMsgUpdateDoc), nil)
	cdc.RegisterConcrete(MsgRevokeDoc{}, fmt.Sprintf("%s/%s", ModuleName, TypeMsgRevokeDoc), nil)
}
