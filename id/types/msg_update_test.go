package types

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestUpdateMsgValidateBasic(t *testing.T) {
	issuerAddr := sdk.AccAddress([]byte(`1`))
	id := "1"
	extra := "2"
	msg := NewMsgUpdateId(issuerAddr, id, extra)

	require.Nil(t, msg.ValidateBasic())
}

func TestUpdateMsgValidateBasic_EmptyIssuer(t *testing.T) {
	issuerAddr := sdk.AccAddress([]byte(``))
	id := "1"
	extra := "2"
	msg := NewMsgUpdateId(issuerAddr, id, extra)

	err := msg.ValidateBasic()
	require.NotNil(t, err)

	skdErr, _ := err.(*sdkerrors.Error)
	require.True(t, skdErr.Is(sdkerrors.ErrInvalidAddress))
}
func TestUpdateMsgValidateBasic_TooLongExtradata(t *testing.T) {
	issuerAddr := sdk.AccAddress([]byte(``))
	id := "1"
	extra := "c89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc61"
	msg := NewMsgUpdateId(issuerAddr, id, extra)

	err := msg.ValidateBasic()
	require.NotNil(t, err)

	skdErr, _ := err.(*sdkerrors.Error)
	require.True(t, skdErr.Is(InvalidData))
}
