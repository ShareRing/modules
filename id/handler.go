package id

import (
	"fmt"

	k "bitbucket.org/sharering/shareledger-modules/id/keeper"
	"bitbucket.org/sharering/shareledger-modules/id/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func NewHandler(keeper k.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		switch msg := msg.(type) {
		case types.MsgCreateId:
			return handleMsgCreateId(ctx, keeper, msg)
			// case MsgUpdateId:
			// 	return handleMsgUpdateId(ctx, keeper, msg)
			// case MsgDeleteId:
			// 	return handleMsgDeleteId(ctx, keeper, msg)
			// case MsgEnrollIDSigners:
			// 	return handleMsgEnrollIdSigners(ctx, keeper, msg)
			// case MsgRevokeIDSigners:
			// return handleMsgRevokeIdSigners(ctx, keeper, msg)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, fmt.Sprintf("Unrecognized identity Msg type: %v", msg.Type()))
		}
	}
}

func handleMsgCreateId(ctx sdk.Context, keeper k.Keeper, msg types.MsgCreateId) (*sdk.Result, error) {

	return &sdk.Result{}, nil
}
