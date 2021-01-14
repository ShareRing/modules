package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type MsgRevokeDoc struct {
	Holder ID
	Issuer sdk.AccAddress
	Proof  []byte
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
