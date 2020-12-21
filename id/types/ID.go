package types

import (
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
	Id string `json:"id"`
	BaseID
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

func (id ID) ToStorage() BaseID {
	return NewBaseID(id.IssuerAddr, id.BackupAddr, id.OwnerAddr, id.ExtraData)
}

func NewIDFromStorage(id string, ids BaseID) ID {
	return ID{id, ids}
}
