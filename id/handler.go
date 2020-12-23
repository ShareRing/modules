package id

import (
	"fmt"

	k "bitbucket.org/shareringvietnam/shareledger-modules/id/keeper"
	"bitbucket.org/shareringvietnam/shareledger-modules/id/types"
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
			fmt.Printf("Unrecognized Id Msg type: %v", msg.Type())
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, fmt.Sprintf("Unrecognized identity Msg type: %v", msg.Type()))
		}
	}
}

func handleMsgCreateId(ctx sdk.Context, keeper k.Keeper, msg types.MsgCreateId) (*sdk.Result, error) {
	id := msg.ToID()
	keeper.SetID(ctx, id)

	event := sdk.NewEvent(
		types.EventCreateID,
		sdk.NewAttribute(types.EventAttrIssuer, msg.IssuerAddr.String()),
		sdk.NewAttribute(types.EventAttrOwner, msg.OwnerAddr.String()),
		sdk.NewAttribute(types.EventAttrId, msg.Id),
		sdk.NewAttribute(types.EventAttrOwner, msg.OwnerAddr.String()),
	)
	ctx.EventManager().EmitEvent(event)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.IssuerAddr.String()),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.EventCreateID),
		),
	)
	return &sdk.Result{
		Log:    msg.String(),
		Events: ctx.EventManager().Events(),
	}, nil
}
