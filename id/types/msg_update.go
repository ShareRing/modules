package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type MsgUpdateId struct {
	Id         string         `json:"id"`
	ExtraData  string         `json:"extra_data"`
	IssuerAddr sdk.AccAddress `json:"issuer_address"`
}

func NewMsgUpdateId(issuerAddr sdk.AccAddress, id, extraData string) MsgUpdateId {
	return MsgUpdateId{
		ExtraData:  extraData,
		Id:         id,
		IssuerAddr: issuerAddr,
	}
}

func (msg MsgUpdateId) Route() string {
	return RouterKey
}

func (msg MsgUpdateId) Type() string {
	return TypeMsgUpdateID
}

func (msg MsgUpdateId) ValidateBasic() error {
	return nil
}

func (msg MsgUpdateId) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgUpdateId) String() string {
	return fmt.Sprintf("id/MsgUpdateId{Id:%s,ExtraData:%s}", msg.Id, msg.ExtraData)
}

func (msg MsgUpdateId) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.IssuerAddr}
}
