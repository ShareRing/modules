package types

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestCreateMsgBatchValidateBasic_NoError(t *testing.T) {
	issuerAddr := sdk.AccAddress([]byte(`1`))
	backupAddrs := []sdk.AccAddress{sdk.AccAddress([]byte(`1`)), sdk.AccAddress([]byte(`2`)), sdk.AccAddress([]byte(`3`))}
	ownerAddrs := []sdk.AccAddress{sdk.AccAddress([]byte(`1`)), sdk.AccAddress([]byte(`2`)), sdk.AccAddress([]byte(`3`))}
	ids := []string{"1", "2", "3"}
	extras := []string{"1", "2", "3"}
	msg := NewMsgCreateIdBatch(issuerAddr, backupAddrs, ownerAddrs, ids, extras)
	err := msg.ValidateBasic()
	require.Nil(t, err)
}

func TestCreateMsgBatchValidateBasic_EmptyIssuer(t *testing.T) {
	issuerAddr := sdk.AccAddress([]byte(``))
	backupAddrs := []sdk.AccAddress{sdk.AccAddress([]byte(`1`)), sdk.AccAddress([]byte(`2`)), sdk.AccAddress([]byte(`3`))}
	ownerAddrs := []sdk.AccAddress{sdk.AccAddress([]byte(`1`)), sdk.AccAddress([]byte(`2`)), sdk.AccAddress([]byte(`3`))}
	ids := []string{"1", "2", "3"}
	extras := []string{"1", "2", "3"}

	msg := NewMsgCreateIdBatch(issuerAddr, backupAddrs, ownerAddrs, ids, extras)
	err := msg.ValidateBasic()
	require.NotNil(t, err)

	skdErr, _ := err.(*sdkerrors.Error)

	require.True(t, skdErr.Is(sdkerrors.ErrInvalidAddress))
}

func TestCreateMsgBatchValidateBasic_NoBackupAddr(t *testing.T) {
	issuerAddr := sdk.AccAddress([]byte(`1`))
	backupAddrs := []sdk.AccAddress{}
	ownerAddrs := []sdk.AccAddress{sdk.AccAddress([]byte(`1`)), sdk.AccAddress([]byte(`2`)), sdk.AccAddress([]byte(`3`))}
	ids := []string{"1", "2", "3"}
	extras := []string{"1", "2", "3"}

	msg := NewMsgCreateIdBatch(issuerAddr, backupAddrs, ownerAddrs, ids, extras)
	err := msg.ValidateBasic()
	require.NotNil(t, err)

	skdErr, _ := err.(*sdkerrors.Error)

	require.True(t, skdErr.Is(InvalidData))
}
func TestCreateMsgBatchValidateBasic_NoOwnerAddr(t *testing.T) {
	issuerAddr := sdk.AccAddress([]byte(`1`))
	backupAddrs := []sdk.AccAddress{sdk.AccAddress([]byte(`1`)), sdk.AccAddress([]byte(`2`)), sdk.AccAddress([]byte(`3`))}
	ownerAddrs := []sdk.AccAddress{}
	ids := []string{"1", "2", "3"}
	extras := []string{"1", "2", "3"}

	msg := NewMsgCreateIdBatch(issuerAddr, backupAddrs, ownerAddrs, ids, extras)
	err := msg.ValidateBasic()
	require.NotNil(t, err)

	skdErr, _ := err.(*sdkerrors.Error)

	require.True(t, skdErr.Is(InvalidData))
}
func TestCreateMsgBatchValidateBasic_NoId(t *testing.T) {
	issuerAddr := sdk.AccAddress([]byte(`1`))
	backupAddrs := []sdk.AccAddress{sdk.AccAddress([]byte(`1`)), sdk.AccAddress([]byte(`2`)), sdk.AccAddress([]byte(`3`))}
	ownerAddrs := []sdk.AccAddress{sdk.AccAddress([]byte(`1`)), sdk.AccAddress([]byte(`2`)), sdk.AccAddress([]byte(`3`))}
	ids := []string{}
	extras := []string{"1", "2", "3"}

	msg := NewMsgCreateIdBatch(issuerAddr, backupAddrs, ownerAddrs, ids, extras)
	err := msg.ValidateBasic()
	require.NotNil(t, err)

	skdErr, _ := err.(*sdkerrors.Error)

	require.True(t, skdErr.Is(InvalidData))
}
func TestCreateMsgBatchValidateBasic_NoExtradata(t *testing.T) {
	issuerAddr := sdk.AccAddress([]byte(`1`))
	backupAddrs := []sdk.AccAddress{sdk.AccAddress([]byte(`1`)), sdk.AccAddress([]byte(`2`)), sdk.AccAddress([]byte(`3`))}
	ownerAddrs := []sdk.AccAddress{sdk.AccAddress([]byte(`1`)), sdk.AccAddress([]byte(`2`)), sdk.AccAddress([]byte(`3`))}
	ids := []string{"1", "2", "3"}
	extras := []string{}

	msg := NewMsgCreateIdBatch(issuerAddr, backupAddrs, ownerAddrs, ids, extras)
	err := msg.ValidateBasic()
	require.NotNil(t, err)

	skdErr, _ := err.(*sdkerrors.Error)

	require.True(t, skdErr.Is(InvalidData))
}

func TestCreateMsgBatchValidateBasic_EmptyBackupAddress(t *testing.T) {
	issuerAddr := sdk.AccAddress([]byte(`1`))
	backupAddrs := []sdk.AccAddress{sdk.AccAddress([]byte(`1`)), sdk.AccAddress([]byte(``)), sdk.AccAddress([]byte(`3`))}
	ownerAddrs := []sdk.AccAddress{sdk.AccAddress([]byte(`1`)), sdk.AccAddress([]byte(`2`)), sdk.AccAddress([]byte(`3`))}
	ids := []string{"1", "2", "3"}
	extras := []string{"1", "2", "3"}

	msg := NewMsgCreateIdBatch(issuerAddr, backupAddrs, ownerAddrs, ids, extras)
	err := msg.ValidateBasic()
	require.NotNil(t, err)

	skdErr, _ := err.(*sdkerrors.Error)

	require.True(t, skdErr.Is(InvalidData))
}
func TestCreateMsgBatchValidateBasic_DuplicateBackupAddress(t *testing.T) {
	issuerAddr := sdk.AccAddress([]byte(`1`))
	backupAddrs := []sdk.AccAddress{sdk.AccAddress([]byte(`1`)), sdk.AccAddress([]byte(`1`)), sdk.AccAddress([]byte(`3`))}
	ownerAddrs := []sdk.AccAddress{sdk.AccAddress([]byte(`1`)), sdk.AccAddress([]byte(`2`)), sdk.AccAddress([]byte(`3`))}
	ids := []string{"1", "2", "3"}
	extras := []string{"1", "2", "3"}

	msg := NewMsgCreateIdBatch(issuerAddr, backupAddrs, ownerAddrs, ids, extras)
	err := msg.ValidateBasic()
	require.Nil(t, err)
}
func TestCreateMsgBatchValidateBasic_EmptyOwnerAddress(t *testing.T) {
	issuerAddr := sdk.AccAddress([]byte(`1`))
	backupAddrs := []sdk.AccAddress{sdk.AccAddress([]byte(`1`)), sdk.AccAddress([]byte(`2`)), sdk.AccAddress([]byte(`3`))}
	ownerAddrs := []sdk.AccAddress{sdk.AccAddress([]byte(`1`)), sdk.AccAddress([]byte(`1`)), sdk.AccAddress([]byte(`3`))}
	ids := []string{"1", "2", "3"}
	extras := []string{"1", "2", "3"}

	msg := NewMsgCreateIdBatch(issuerAddr, backupAddrs, ownerAddrs, ids, extras)
	err := msg.ValidateBasic()
	require.NotNil(t, err)

	skdErr, _ := err.(*sdkerrors.Error)

	require.True(t, skdErr.Is(InvalidData))
}
func TestCreateMsgBatchValidateBasic_DuplicateOwnerAddress(t *testing.T) {
	issuerAddr := sdk.AccAddress([]byte(`1`))
	backupAddrs := []sdk.AccAddress{sdk.AccAddress([]byte(`1`)), sdk.AccAddress([]byte(`2`)), sdk.AccAddress([]byte(`3`))}
	ownerAddrs := []sdk.AccAddress{sdk.AccAddress([]byte(`1`)), sdk.AccAddress([]byte(``)), sdk.AccAddress([]byte(`3`))}
	ids := []string{"1", "2", "3"}
	extras := []string{"1", "2", "3"}

	msg := NewMsgCreateIdBatch(issuerAddr, backupAddrs, ownerAddrs, ids, extras)
	err := msg.ValidateBasic()
	require.NotNil(t, err)

	skdErr, _ := err.(*sdkerrors.Error)

	require.True(t, skdErr.Is(InvalidData))
}
func TestCreateMsgBatchValidateBasic_EmptyId(t *testing.T) {
	issuerAddr := sdk.AccAddress([]byte(`1`))
	backupAddrs := []sdk.AccAddress{sdk.AccAddress([]byte(`1`)), sdk.AccAddress([]byte(`2`)), sdk.AccAddress([]byte(`3`))}
	ownerAddrs := []sdk.AccAddress{sdk.AccAddress([]byte(`1`)), sdk.AccAddress([]byte(`2`)), sdk.AccAddress([]byte(`3`))}
	ids := []string{"1", "", "3"}
	extras := []string{"1", "2", "3"}

	msg := NewMsgCreateIdBatch(issuerAddr, backupAddrs, ownerAddrs, ids, extras)
	err := msg.ValidateBasic()
	require.NotNil(t, err)

	skdErr, _ := err.(*sdkerrors.Error)

	require.True(t, skdErr.Is(InvalidData))
}
func TestCreateMsgBatchValidateBasic_DuplicateId(t *testing.T) {
	issuerAddr := sdk.AccAddress([]byte(`1`))
	backupAddrs := []sdk.AccAddress{sdk.AccAddress([]byte(`1`)), sdk.AccAddress([]byte(`2`)), sdk.AccAddress([]byte(`3`))}
	ownerAddrs := []sdk.AccAddress{sdk.AccAddress([]byte(`1`)), sdk.AccAddress([]byte(`2`)), sdk.AccAddress([]byte(`3`))}
	ids := []string{"1", "2", "1"}
	extras := []string{"1", "2", "3"}

	msg := NewMsgCreateIdBatch(issuerAddr, backupAddrs, ownerAddrs, ids, extras)
	err := msg.ValidateBasic()
	require.NotNil(t, err)

	skdErr, _ := err.(*sdkerrors.Error)

	require.True(t, skdErr.Is(InvalidData))
}
func TestCreateMsgBatchValidateBasic_TooLongId(t *testing.T) {
	issuerAddr := sdk.AccAddress([]byte(`1`))
	backupAddrs := []sdk.AccAddress{sdk.AccAddress([]byte(`1`)), sdk.AccAddress([]byte(`2`)), sdk.AccAddress([]byte(`3`))}
	ownerAddrs := []sdk.AccAddress{sdk.AccAddress([]byte(`1`)), sdk.AccAddress([]byte(`2`)), sdk.AccAddress([]byte(`3`))}
	ids := []string{"1", "c89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc61", "3"}
	extras := []string{"1", "2", "3"}

	msg := NewMsgCreateIdBatch(issuerAddr, backupAddrs, ownerAddrs, ids, extras)
	err := msg.ValidateBasic()
	require.NotNil(t, err)

	skdErr, _ := err.(*sdkerrors.Error)

	require.True(t, skdErr.Is(InvalidData))
}
func TestCreateMsgBatchValidateBasic_TooLongExtra(t *testing.T) {
	issuerAddr := sdk.AccAddress([]byte(`1`))
	backupAddrs := []sdk.AccAddress{sdk.AccAddress([]byte(`1`)), sdk.AccAddress([]byte(`2`)), sdk.AccAddress([]byte(`3`))}
	ownerAddrs := []sdk.AccAddress{sdk.AccAddress([]byte(`1`)), sdk.AccAddress([]byte(`2`)), sdk.AccAddress([]byte(`3`))}
	ids := []string{"1", "2", "3"}
	extras := []string{"1", "c89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc61", "3"}

	msg := NewMsgCreateIdBatch(issuerAddr, backupAddrs, ownerAddrs, ids, extras)
	err := msg.ValidateBasic()
	require.NotNil(t, err)

	skdErr, _ := err.(*sdkerrors.Error)

	require.True(t, skdErr.Is(InvalidData))
}
