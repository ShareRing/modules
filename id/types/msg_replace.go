package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type MsgReplaceIdOwner struct {
	BackupAddr sdk.AccAddress `json:"backup_address"`
	Id         string         `json:"id"`
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
	if len(msg.Id) > MAX_ID_LEN || len(msg.Id) == 0 {
		return InvalidData
	}
	if msg.BackupAddr.Empty() || msg.OwnerAddr.Empty() {
		return sdkerrors.ErrInvalidAddress
	}
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
