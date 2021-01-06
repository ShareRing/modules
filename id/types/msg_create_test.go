package types

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/crypto/secp256k1"
)

func TestCreateMsgValidateBasic(t *testing.T) {
	issuerPriv := secp256k1.GenPrivKey()
	issuerAddr := sdk.AccAddress(issuerPriv.PubKey().Address())
	backupAddr := sdk.AccAddress(issuerPriv.PubKey().Address())
	ownerAddr := sdk.AccAddress(issuerPriv.PubKey().Address())
	id := "1"
	extra := "2"
	msg := NewMsgCreateId(issuerAddr, backupAddr, ownerAddr, id, extra)

	require.Nil(t, msg.ValidateBasic())
}
func TestCreateMsgValidateBasic_EmptyIssuer(t *testing.T) {
	issuerPriv := secp256k1.GenPrivKey()
	issuerAddr := sdk.AccAddress{}
	backupAddr := sdk.AccAddress(issuerPriv.PubKey().Address())
	ownerAddr := sdk.AccAddress(issuerPriv.PubKey().Address())
	id := "1"
	extra := "2"
	msg := NewMsgCreateId(issuerAddr, backupAddr, ownerAddr, id, extra)
	err := msg.ValidateBasic()
	require.NotNil(t, err)

	skdErr, _ := err.(*sdkerrors.Error)

	require.True(t, skdErr.Is(sdkerrors.ErrInvalidAddress))
}

func TestCreateMsgValidateBasic_EmptyBackupAddr(t *testing.T) {
	issuerPriv := secp256k1.GenPrivKey()
	issuerAddr := sdk.AccAddress(issuerPriv.PubKey().Address())
	backupAddr := sdk.AccAddress{}
	ownerAddr := sdk.AccAddress(issuerPriv.PubKey().Address())
	id := "1"
	extra := "2"
	msg := NewMsgCreateId(issuerAddr, backupAddr, ownerAddr, id, extra)
	err := msg.ValidateBasic()
	require.NotNil(t, err)

	skdErr, _ := err.(*sdkerrors.Error)

	require.True(t, skdErr.Is(sdkerrors.ErrInvalidAddress))
}

func TestCreateMsgValidateBasic_EmptyOwnerAddr(t *testing.T) {
	issuerPriv := secp256k1.GenPrivKey()
	issuerAddr := sdk.AccAddress(issuerPriv.PubKey().Address())
	backupAddr := sdk.AccAddress(issuerPriv.PubKey().Address())
	ownerAddr := sdk.AccAddress{}
	id := "1"
	extra := "2"
	msg := NewMsgCreateId(issuerAddr, backupAddr, ownerAddr, id, extra)
	err := msg.ValidateBasic()
	require.NotNil(t, err)

	skdErr, _ := err.(*sdkerrors.Error)

	require.True(t, skdErr.Is(sdkerrors.ErrInvalidAddress))
}
func TestCreateMsgValidateBasic_EmptyId(t *testing.T) {
	issuerPriv := secp256k1.GenPrivKey()
	issuerAddr := sdk.AccAddress(issuerPriv.PubKey().Address())
	backupAddr := sdk.AccAddress(issuerPriv.PubKey().Address())
	ownerAddr := sdk.AccAddress(issuerPriv.PubKey().Address())
	id := ""
	extra := "2"
	msg := NewMsgCreateId(issuerAddr, backupAddr, ownerAddr, id, extra)
	err := msg.ValidateBasic()
	require.NotNil(t, err)

	skdErr, _ := err.(*sdkerrors.Error)

	require.True(t, skdErr.Is(InvalidData))
}

func TestCreateMsgValidateBasic_TooLongId(t *testing.T) {
	issuerPriv := secp256k1.GenPrivKey()
	issuerAddr := sdk.AccAddress(issuerPriv.PubKey().Address())
	backupAddr := sdk.AccAddress(issuerPriv.PubKey().Address())
	ownerAddr := sdk.AccAddress(issuerPriv.PubKey().Address())
	id := "c89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc61"
	extra := "2"
	msg := NewMsgCreateId(issuerAddr, backupAddr, ownerAddr, id, extra)
	err := msg.ValidateBasic()
	require.NotNil(t, err)

	skdErr, _ := err.(*sdkerrors.Error)

	require.True(t, skdErr.Is(InvalidData))
}

func TestCreateMsgValidateBasic_TooLongExtradata(t *testing.T) {
	issuerPriv := secp256k1.GenPrivKey()
	issuerAddr := sdk.AccAddress(issuerPriv.PubKey().Address())
	backupAddr := sdk.AccAddress(issuerPriv.PubKey().Address())
	ownerAddr := sdk.AccAddress(issuerPriv.PubKey().Address())
	id := "1"
	extra := "c89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc61"
	msg := NewMsgCreateId(issuerAddr, backupAddr, ownerAddr, id, extra)
	err := msg.ValidateBasic()
	require.NotNil(t, err)

	skdErr, _ := err.(*sdkerrors.Error)

	require.True(t, skdErr.Is(InvalidData))
}
