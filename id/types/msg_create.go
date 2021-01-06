package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	MAX_ID_LEN = 64
)

type MsgCreateId struct {
	Id         string         `json:"id"`
	IssuerAddr sdk.AccAddress `json:"issuer_address"`
	BackupAddr sdk.AccAddress `json:"backup_address"`
	OwnerAddr  sdk.AccAddress `json:"owner_address"`
	ExtraData  string         `json:"extra_data"`
}

func NewMsgCreateId(issuerAddr, backupAddr, ownerAddr sdk.AccAddress, id, extraData string) MsgCreateId {
	return MsgCreateId{
		IssuerAddr: issuerAddr,
		BackupAddr: backupAddr,
		OwnerAddr:  ownerAddr,
		ExtraData:  extraData,
		Id:         id,
	}
}

func (msg MsgCreateId) Route() string {
	return RouterKey
}

func (msg MsgCreateId) Type() string {
	return TypeMsgCreateID
}

func (msg MsgCreateId) ValidateBasic() error {
	if len(msg.Id) > MAX_ID_LEN || len(msg.Id) == 0 || len(msg.ExtraData) > MAX_ID_LEN {
		return InvalidData
	}
	if msg.BackupAddr.Empty() || msg.IssuerAddr.Empty() || msg.OwnerAddr.Empty() {
		return sdkerrors.ErrInvalidAddress
	}
	return nil
}

func (msg MsgCreateId) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgCreateId) String() string {
	return fmt.Sprintf("id/MsgCreateId{Id:%s,IssuerAddr:%s,OwnerAddr:%s,BackupAddr:%s,ExtraData:%s}", msg.Id, msg.IssuerAddr, msg.OwnerAddr.String(), msg.BackupAddr.String(), msg.ExtraData)
}

func (msg MsgCreateId) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.IssuerAddr}
}

func (msg MsgCreateId) ToBaseID() BaseID {
	baseId := NewBaseID(msg.IssuerAddr, msg.BackupAddr, msg.OwnerAddr, msg.ExtraData)
	return baseId
}

func (msg MsgCreateId) ToID() ID {
	baseId := msg.ToBaseID()
	id := NewIDFromBaseID(msg.Id, baseId)
	return id
}
