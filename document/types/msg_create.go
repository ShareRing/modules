package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type ID []byte

type MsgCreateDoc struct {
	Holder ID
	Issuer sdk.AccAddress
	Proof  []byte
	Data   []byte
}

func (msg MsgCreateDoc) Route() string {
	return RouterKey
}

func (msg MsgCreateDoc) Type() string {
	return TypeMsgCreateDoc
}

func (msg MsgCreateDoc) ValidateBasic() error {

	return nil
}

func (msg MsgCreateDoc) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgCreateDoc) String() string {
	return fmt.Sprintf("%s/%s{Holder:%s,Issuer:%s,Proof:%s,Data:%s}", ModuleName, TypeMsgCreateDoc, msg.Holder, msg.Issuer.String(), msg.Proof, msg.Data)
}

func (msg MsgCreateDoc) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Issuer}
}
