package document

import (
	"github.com/ShareRing/modules/document/keeper"
	"github.com/ShareRing/modules/document/types"
)

type (
	Keeper       = keeper.Keeper
	MsgCreateDoc = types.MsgCreateDoc
)

const (
	StoreKey     = types.StoreKey
	ModuleName   = types.ModuleName
	QuerierRoute = types.QuerierRoute
	RouterKey    = types.RouterKey

	TypeMsgCreateDoc        = types.TypeMsgCreateDoc
	TypeMsgCreateDocInBatch = types.TypeMsgCreateDocInBatch
	TypeMsgUpdateDoc        = types.TypeMsgUpdateDoc
	TypeMsgRevokeDoc        = types.TypeMsgRevokeDoc
)

var (
	NewKeeper = keeper.NewKeeper
	// NewQuerier    = keeper.NewQuerier
	ModuleCdc     = types.ModuleCdc
	RegisterCodec = types.RegisterCodec
)

var (
	ErrDocNotExisted  = types.ErrDocNotExisted
	ErrDocExisted     = types.ErrDocExisted
	ErrDocInvalidData = types.ErrDocInvalidData
)

var (
	DocRevokeFlag = types.DocRevokeFlag
)
