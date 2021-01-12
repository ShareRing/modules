package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type MsgCreateDocBatch struct {
	holder []ID
	issuer sdk.AccAddress
	proof  []byte
	extra  [][]byte
}

type MsgUpdateDoc struct {
	holder ID
	issuer sdk.AccAddress
	proof  []byte
	extra  []byte
}

type MsgRevokeDoc struct {
	holder ID
	issuer sdk.AccAddress
	proof  []byte
}
