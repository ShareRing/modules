package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type MsgReplaceIdOwner struct {
	Id         string         `json:"id"`
	BackupAddr sdk.AccAddress `json:"backup_address"`
	OwnerAddr  sdk.AccAddress `json:"owner_address"`
}

func NewMsgReplaceIdOwner(id string, ownerAddr, backupAddr sdk.AccAddress) MsgReplaceIdOwner {
	return MsgReplaceIdOwner{
		Id:         id,
		BackupAddr: backupAddr,
		OwnerAddr:  ownerAddr,
	}
}

func (msg MsgReplaceIdOwner) Route() string {
	return RouterKey
}

func (msg MsgReplaceIdOwner) Type() string {
	return TypeMsgReplaceIdOwner
}

func (msg MsgReplaceIdOwner) ValidateBasic() error {
	// TODO
	return nil
}

func (msg MsgReplaceIdOwner) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgReplaceIdOwner) String() string {
	return fmt.Sprintf("id/MsgReplaceIdOwner{Id:%s,OwnerAddr:%s}", msg.Id, msg.OwnerAddr)
}

func (msg MsgReplaceIdOwner) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.BackupAddr}
}
