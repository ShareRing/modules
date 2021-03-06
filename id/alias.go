package id

import (
	"github.com/ShareRing/modules/id/keeper"
	"github.com/ShareRing/modules/id/types"
)

type (
	Keeper            = keeper.Keeper
	MsgReplaceIdOwner = types.MsgReplaceIdOwner
	ID                = types.ID
	BaseID            = types.BaseID
)

const (
	StoreKey              = types.StoreKey
	ModuleName            = types.ModuleName
	QuerierRoute          = types.QuerierRoute
	RouterKey             = types.RouterKey
	TypeMsgCreateID       = types.TypeMsgCreateID
	TypeMsgCreateIDBatch  = types.TypeMsgCreateIDBatch
	TypeMsgReplaceIdOwner = types.TypeMsgReplaceIdOwner
	TypeMsgUpdateID       = types.TypeMsgUpdateID
)

var (
	NewKeeper = keeper.NewKeeper
	// NewQuerier    = keeper.NewQuerier
	ModuleCdc     = types.ModuleCdc
	RegisterCodec = types.RegisterCodec
)

var (
	ErrIdNotExisted    = types.ErrIdNotExisted
	ErrIdExisted       = types.ErrIdExisted
	ErrWrongBackupAddr = types.ErrWrongBackupAddr
)
