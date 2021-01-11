package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type MsgCreateIdBatch struct {
	IssuerAddr sdk.AccAddress   `json:"issuer_address"`
	Id         []string         `json:"id"`
	BackupAddr []sdk.AccAddress `json:"backup_address"`
	OwnerAddr  []sdk.AccAddress `json:"owner_address"`
	ExtraData  []string         `json:"extra_data"`
}

func NewMsgCreateIdBatch(issuerAddr sdk.AccAddress, backupAddr, ownerAddr []sdk.AccAddress, id, extraData []string) MsgCreateIdBatch {
	return MsgCreateIdBatch{
		IssuerAddr: issuerAddr,
		BackupAddr: backupAddr,
		OwnerAddr:  ownerAddr,
		ExtraData:  extraData,
		Id:         id,
	}
}

func (msg MsgCreateIdBatch) Route() string {
	return RouterKey
}

func (msg MsgCreateIdBatch) Type() string {
	return TypeMsgCreateIDBatch
}

func (msg MsgCreateIdBatch) ValidateBasic() error {
	if msg.IssuerAddr.Empty() {
		return sdkerrors.ErrInvalidAddress
	}

	// Check len
	if len(msg.Id) == 0 || len(msg.BackupAddr) == 0 || len(msg.OwnerAddr) == 0 || len(msg.ExtraData) == 0 {
		return InvalidData
	}

	maxLen := len(msg.Id)
	if len(msg.BackupAddr) != maxLen || len(msg.OwnerAddr) != maxLen || len(msg.ExtraData) != maxLen {
		return InvalidData
	}

	for i := 0; i < maxLen; i++ {
		// Check len
		if len(msg.Id[i]) > MAX_ID_LEN || len(msg.Id[i]) == 0 || len(msg.ExtraData[i]) > MAX_ID_LEN {
			return InvalidData
		}

		// Check address
		if msg.OwnerAddr[i].Empty() || msg.BackupAddr[i].Empty() {
			return InvalidData
		}

		// Check duplicate
		for j := i + 1; j < maxLen; j++ {
			if msg.Id[j] == msg.Id[i] || msg.OwnerAddr[j].Equals(msg.OwnerAddr[i]) {
				return InvalidData
			}
		}
	}

	return nil
}

func (msg MsgCreateIdBatch) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgCreateIdBatch) String() string {
	return fmt.Sprintf("id/MsgCreateIdBatch{Issuer:%s,Total:%d}", msg.IssuerAddr, len(msg.OwnerAddr))
}

func (msg MsgCreateIdBatch) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.IssuerAddr}
}
