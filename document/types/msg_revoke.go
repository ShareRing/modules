package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type MsgRevokeDoc struct {
	Holder string         `json:"holder"`
	Issuer sdk.AccAddress `json:"issuer"`
	Proof  string         `json:"proof"`
}

func NewMsgRevokeDoc(issuer sdk.AccAddress, holder, proof string) MsgRevokeDoc {
	return MsgRevokeDoc{Issuer: issuer, Holder: holder, Proof: proof}
}

func (msg MsgRevokeDoc) Route() string {
	return RouterKey
}

func (msg MsgRevokeDoc) Type() string {
	return TypeMsgRevokeDoc
}

func (msg MsgRevokeDoc) ValidateBasic() error {

	return nil
}

func (msg MsgRevokeDoc) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgRevokeDoc) String() string {
	return fmt.Sprintf("%s/%s{Holder:%s,Issuer:%s,Proof:%s}", ModuleName, TypeMsgRevokeDoc, msg.Holder, msg.Issuer.String(), msg.Proof)
}

func (msg MsgRevokeDoc) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Issuer}
}
