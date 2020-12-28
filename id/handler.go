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
		case types.MsgUpdateId:
			return handleMsgUpdateId(ctx, keeper, msg)
		case types.MsgCreateIdBatch:
			return handleMsgCreateIdBatch(ctx, keeper, msg)
		case types.MsgReplaceIdOwner:
			return handleMsgReplaceOwnerId(ctx, keeper, msg)
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

	// Check existing
	if keeper.IsExist(ctx, &id) {
		return nil, ErrIdExisted
	}

	keeper.SetID(ctx, &id)

	event := sdk.NewEvent(
		types.EventCreateID,
		sdk.NewAttribute(types.EventAttrIssuer, msg.IssuerAddr.String()),
		sdk.NewAttribute(types.EventAttrOwner, msg.OwnerAddr.String()),
		sdk.NewAttribute(types.EventAttrId, msg.Id),
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

func handleMsgCreateIdBatch(ctx sdk.Context, keeper k.Keeper, msg types.MsgCreateIdBatch) (*sdk.Result, error) {
	for i := 0; i < len(msg.Id); i++ {
		id := types.NewID(msg.Id[i], msg.IssuerAddr, msg.BackupAddr[i], msg.OwnerAddr[i], msg.ExtraData[i])
		// Check id existing
		if keeper.IsExist(ctx, &id) {
			return nil, ErrIdExisted
		}

		keeper.SetID(ctx, &id)
		event := sdk.NewEvent(
			types.EventCreateID,
			sdk.NewAttribute(types.EventAttrIssuer, msg.IssuerAddr.String()),
			sdk.NewAttribute(types.EventAttrOwner, msg.OwnerAddr[i].String()),
			sdk.NewAttribute(types.EventAttrId, msg.Id[i]),
		)
		ctx.EventManager().EmitEvent(event)
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.IssuerAddr.String()),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.EventCreateIDBatch),
		),
	)

	return &sdk.Result{
		Log:    msg.String(),
		Events: ctx.EventManager().Events(),
	}, nil
}

func handleMsgUpdateId(ctx sdk.Context, keeper k.Keeper, msg types.MsgUpdateId) (*sdk.Result, error) {
	id := keeper.GetIDByIdString(ctx, msg.Id)

	if id.IsEmpty() {
		return nil, ErrIdNotExisted
	}

	// Update extra data
	id.ExtraData = msg.ExtraData
	keeper.SetID(ctx, id)

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventUpdateID,
		sdk.NewAttribute(types.EventAttrIssuer, msg.IssuerAddr.String()),
		sdk.NewAttribute(types.EventAttrId, msg.Id),
	))

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.IssuerAddr.String()),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.EventUpdateID),
		),
	)
	return &sdk.Result{
		Log:    msg.String(),
		Events: ctx.EventManager().Events(),
	}, nil

}

func handleMsgReplaceOwnerId(ctx sdk.Context, keeper k.Keeper, msg types.MsgReplaceIdOwner) (*sdk.Result, error) {
	id := keeper.GetIDByIdString(ctx, msg.Id)

	// Check if the id is existed or not
	if id.IsEmpty() {
		return nil, ErrIdNotExisted
	}

	// Check if the new owner has id or not
	idOfNewOwner := keeper.GetIdOnlyByAddress(ctx, msg.OwnerAddr)
	if len(idOfNewOwner) > 0 {
		return nil, ErrIdExisted
	}

	// Update owner
	id.OwnerAddr = msg.OwnerAddr
	keeper.SetID(ctx, id)

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventReplaceIDOwner,
		sdk.NewAttribute(types.EventAttrOwner, msg.OwnerAddr.String()),
		sdk.NewAttribute(types.EventAttrId, msg.Id),
	))

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.BackupAddr.String()),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.EventReplaceIDOwner),
		),
	)

	return &sdk.Result{
		Log:    msg.String(),
		Events: ctx.EventManager().Events(),
	}, nil

}
