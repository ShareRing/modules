package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	MAX_DOC_DATA_LEN = 64
	MAX_BATCH_LENGH  = 20
)

type MsgCreateDocBatch struct {
	Data   []string       `json:"data"`
	Issuer sdk.AccAddress `json:"issuer"`
	Holder []string       `json:"holder"`
	Proof  []string       `json:"proof"`
}

func NewMsgCreateDocBatch(issuer sdk.AccAddress, holderId, proof, data []string) MsgCreateDocBatch {
	return MsgCreateDocBatch{Issuer: issuer, Holder: holderId, Proof: proof, Data: data}
}

func (msg MsgCreateDocBatch) Route() string {
	return RouterKey
}

func (msg MsgCreateDocBatch) Type() string {
	return TypeMsgCreateDocInBatch
}

func (msg MsgCreateDocBatch) ValidateBasic() error {
	if msg.Issuer.Empty() {
		return sdkerrors.ErrInvalidAddress
	}

	// Check len
	if len(msg.Holder) == 0 || len(msg.Proof) == 0 || len(msg.Data) == 0 {
		return ErrDocInvalidData
	}

	maxLen := len(msg.Holder)
	if maxLen > MAX_BATCH_LENGH || len(msg.Proof) != maxLen || len(msg.Data) != maxLen {
		return ErrDocInvalidData
	}

	for i := 0; i < maxLen; i++ {
		// Check len
		if len(msg.Holder[i]) > MAX_DOC_DATA_LEN || len(msg.Holder[i]) == 0 || len(msg.Data[i]) > MAX_DOC_DATA_LEN {
			return ErrDocInvalidData
		}

		// Check duplicate
		for j := i + 1; j < maxLen; j++ {
			if msg.Holder[j] == msg.Holder[i] || msg.Proof[j] == msg.Proof[i] {
				return ErrDocInvalidData
			}
		}
	}
	return nil
}

func (msg MsgCreateDocBatch) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgCreateDocBatch) String() string {
	return fmt.Sprintf("%s/%s{Issuer:%s,Len:%d}", ModuleName, TypeMsgCreateDocInBatch, msg.Issuer.String(), len(msg.Holder))
}

func (msg MsgCreateDocBatch) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Issuer}
}
