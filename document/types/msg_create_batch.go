package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type MsgCreateDocBatch struct {
	Issuer sdk.AccAddress
	Holder []ID
	Proof  [][]byte
	Data   [][]byte
}

func (msg MsgCreateDocBatch) Route() string {
	return RouterKey
}

func (msg MsgCreateDocBatch) Type() string {
	return TypeMsgCreateDocInBatch
}

func (msg MsgCreateDocBatch) ValidateBasic() error {

	return nil
}

func (msg MsgCreateDocBatch) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgCreateDocBatch) String() string {
	return fmt.Sprintf("%s/%s{Issuer:%s,Len:%d}", ModuleName, TypeMsgCreateDocInBatch, msg.Issuer.String(), len(msg.Holder))
}

func (msg MsgCreateDocBatch) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Issuer}
}
