package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type MsgUpdateDoc struct {
	Holder string
	Issuer sdk.AccAddress
	Proof  string
	Data   string
}

func NewMsgUpdateDoc(issuer sdk.AccAddress, holderId, proof, data string) MsgUpdateDoc {
	return MsgUpdateDoc{Issuer: issuer, Holder: holderId, Proof: proof, Data: data}
}

func (msg MsgUpdateDoc) Route() string {
	return RouterKey
}

func (msg MsgUpdateDoc) Type() string {
	return TypeMsgCreateDocInBatch
}

func (msg MsgUpdateDoc) ValidateBasic() error {

	return nil
}

func (msg MsgUpdateDoc) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgUpdateDoc) String() string {
	return fmt.Sprintf("%s/%s{Holder:%s,Issuer:%s,Proof:%s,Data:%s}", ModuleName, TypeMsgUpdateDoc, msg.Holder, msg.Issuer.String(), msg.Proof, msg.Data)
}

func (msg MsgUpdateDoc) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Issuer}
}
