package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// ID message types
const (
	TypeMsgCreateID = "create_id"
)

type MsgCreateId struct {
	IssuerAddr sdk.AccAddress `json:"issuer_address"`
	BackupAddr sdk.AccAddress `json:"backup_address"`
	OwnerAddr  sdk.AccAddress `json:"owner_address"`
	ExtraData  string         `json:"extra_data"`
}

func NewMsgCreateId(issuerAddr, backupAddr, ownerAddr sdk.AccAddress, extraData string) MsgCreateId {
	return MsgCreateId{
		IssuerAddr: issuerAddr,
		BackupAddr: backupAddr,
		OwnerAddr:  ownerAddr,
		ExtraData:  extraData,
	}
}

func (msg MsgCreateId) Route() string {
	return ""
}

func (msg MsgCreateId) Type() string {
	return TypeMsgCreateID
}

// If sender is different with owner
// sender must be registered as IDSigner
func (msg MsgCreateId) ValidateBasic() error {
	return nil
}

func (msg MsgCreateId) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgCreateId) String() string {
	return fmt.Sprintf("id/MsgCreateId{IssuerAddr:%s,OwnerAddr:%s,BackupAddr:%s,ExtraData:%s}", msg.IssuerAddr, msg.OwnerAddr.String(), msg.BackupAddr.String(), msg.ExtraData)
}

func (msg MsgCreateId) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.IssuerAddr}
}
