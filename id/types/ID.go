package types

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type BaseID struct {
	IssuerAddr sdk.AccAddress `json:"issuer_address"`
	BackupAddr sdk.AccAddress `json:"backup_address"`
	OwnerAddr  sdk.AccAddress `json:"owner_address"`
	ExtraData  string         `json:"extra_data"`
}

type ID struct {
	Id     string `json:"id"`
	BaseID `json:"data"`
}

func NewBaseID(issuerAddr, backupAddr, ownerAddr sdk.AccAddress, extraData string) BaseID {
	return BaseID{issuerAddr, backupAddr, ownerAddr, extraData}
}

// Marshal ID
func MustMarshalBaseID(cdc *codec.Codec, id BaseID) []byte {
	return cdc.MustMarshalBinaryLengthPrefixed(id)
}

// Unmarshal ID from store value
func UnmarshalBaseID(cdc *codec.Codec, value []byte) (id BaseID, err error) {
	err = cdc.UnmarshalBinaryLengthPrefixed(value, &id)
	return id, err
}

// Unmarshal ID from store value. Throw exception when there is error
func MustUnmarshalBaseID(cdc *codec.Codec, value []byte) BaseID {
	id, err := UnmarshalBaseID(cdc, value)
	if err != nil {
		panic(err)
	}
	return id
}

func (bId BaseID) String() string {
	return fmt.Sprintf("{IssuerAddr:%s,OwnerAddr:%s,BackupAddr:%s,ExtraData:%s}", bId.IssuerAddr, bId.OwnerAddr.String(), bId.BackupAddr.String(), bId.ExtraData)
}

func (id ID) ToBaseID() BaseID {
	return NewBaseID(id.IssuerAddr, id.BackupAddr, id.OwnerAddr, id.ExtraData)
}

func NewIDFromBaseID(id string, ids BaseID) ID {
	return ID{id, ids}
}

func NewID(id string, issuerAddr, backupAddr, ownerAddr sdk.AccAddress, extraData string) ID {
	return ID{id, BaseID{issuerAddr, backupAddr, ownerAddr, extraData}}
}

func (id *ID) IsEmpty() bool {
	if id == nil {
		return true
	}

	if len(id.IssuerAddr) == 0 {
		return true
	}

	return false
}
