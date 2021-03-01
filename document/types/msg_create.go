package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type MsgCreateDoc struct {
	Data   string         `json:"data"`
	Holder string         `json:"holder"`
	Issuer sdk.AccAddress `json:"issuer"`
	Proof  string         `json:"proof"`
}

func NewMsgCreateDoc(issuer sdk.AccAddress, holderId, proof, data string) MsgCreateDoc {
	return MsgCreateDoc{Issuer: issuer, Holder: holderId, Proof: proof, Data: data}
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
