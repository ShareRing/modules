package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrDocNotExisted  = sdkerrors.Register(ModuleName, 2, "Doc does not exist")
	ErrDocExisted     = sdkerrors.Register(ModuleName, 3, "Doc existed")
	ErrDocInvalidData = sdkerrors.Register(ModuleName, 4, "Invalid data")
	ErrorNotIssuer    = sdkerrors.Register(ModuleName, 5, "Not doc issuer")
)
