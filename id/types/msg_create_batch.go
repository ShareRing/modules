package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
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

// Check len
func (msg MsgCreateIdBatch) ValidateBasic() error {
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
