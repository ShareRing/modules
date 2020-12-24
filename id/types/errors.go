package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrIdNotExisted = sdkerrors.Register(ModuleName, 2, "Id does not exist")
	ErrIdExisted    = sdkerrors.Register(ModuleName, 3, "Id existed")
)
