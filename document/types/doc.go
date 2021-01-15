package types

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Doc struct {
	Holder  string         `json:"holder"`
	Issuer  sdk.AccAddress `json:"issuer"`
	Proof   string         `json:"proof"`
	Data    string         `json:"data"`
	Version uint16         `json:"version"`
}

type DocDetailState struct {
	Data    string
	Version uint16
}

type DocBasicState struct {
	Holder string
	Issuer sdk.AccAddress
}

func (d Doc) GetDetailState() DocDetailState {
	ds := DocDetailState{d.Data, d.Version}
	return ds
}

// 0x1|<hodler>|<proof>|<issuer>
func (d Doc) GetKeyDetailState() []byte {
	key := []byte{}
	key = append(key, StateKeySep...)
	key = append(key, d.Holder...)

	key = append(key, StateKeySep...)
	key = append(key, d.Proof...)

	key = append(key, StateKeySep...)
	key = append(key, d.Issuer...)

	key = append(DocDetailPrefix, key...)

	return key
}

// Marshal doc state
func MustMarshalDocDetailState(cdc *codec.Codec, ds DocDetailState) []byte {
	return cdc.MustMarshalBinaryLengthPrefixed(ds)
}

func MustUnmarshalDocDetailState(cdc *codec.Codec, value []byte) DocDetailState {
	ds := DocDetailState{}

	err := cdc.UnmarshalBinaryLengthPrefixed(value, &ds)
	if err != nil {
		panic(err)
	}

	return ds
}

func (d Doc) GetBasicState() DocBasicState {
	ds := DocBasicState{Holder: d.Holder, Issuer: d.Issuer}
	return ds
}

// 0x2|<proof>
func (d Doc) GetKeyBasicState() []byte {
	key := []byte{}
	key = append(key, StateKeySep...)
	key = append(key, d.Proof...)

	key = append(DocBasicPrefix, key...)
	return key
}

// Marshal doc state
func MustMarshalDocBasicState(cdc *codec.Codec, bs DocBasicState) []byte {
	return cdc.MustMarshalBinaryLengthPrefixed(bs)
}

func MustUnmarshalDocBasicState(cdc *codec.Codec, value []byte) DocBasicState {
	ds := DocBasicState{}

	err := cdc.UnmarshalBinaryLengthPrefixed(value, &ds)
	if err != nil {
		panic(err)
	}

	return ds
}

func (d Doc) String() string {
	s := fmt.Sprintf("Hodler %v, issuer %v, Proof: %v, Data: %v, Ver: %d", d.Holder, d.Issuer, d.Proof, d.Data, d.Version)
	return s
}
