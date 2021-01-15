package document

import (
	"fmt"

	types "bitbucket.org/shareringvietnam/shareledger-modules/document/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		switch msg := msg.(type) {
		case types.MsgCreateDoc:
			return handleMsgCreateDoc(ctx, keeper, msg)
		case types.MsgCreateDocBatch:
			return handleMsgCreateDocInBatch(ctx, keeper, msg)
		case types.MsgUpdateDoc:
			return handleMsgUpdateDoc(ctx, keeper, msg)
		case types.MsgRevokeDoc:
			return handleMsgRevokeDoc(ctx, keeper, msg)

		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, fmt.Sprintf("Unrecognized %s Msg type: %v", types.ModuleName, msg.Type()))
		}
	}
}

func handleMsgCreateDoc(ctx sdk.Context, k Keeper, msg types.MsgCreateDoc) (*sdk.Result, error) {
	doc := types.Doc{Issuer: msg.Issuer, Holder: msg.Holder, Proof: msg.Proof, Data: msg.Data, Version: 0}

	// Check doc is existed
	existingDoc := k.GetDoc(ctx, doc)
	if len(existingDoc.Proof) > 0 {
		return nil, types.ErrDocExisted
	}

	k.SetDoc(ctx, doc)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeCreateDoc,
			sdk.NewAttribute(types.EventAttrIssuer, msg.Issuer.String()),
			sdk.NewAttribute(types.EventAttrHolder, string(msg.Holder)),
			sdk.NewAttribute(types.EventAttrProof, string(msg.Proof)),
		),
	)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Issuer.String()),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.EventTypeCreateDoc),
		),
	)

	return &sdk.Result{
		Log:    msg.String(),
		Events: ctx.EventManager().Events(),
	}, nil
}

func handleMsgCreateDocInBatch(ctx sdk.Context, k Keeper, msg types.MsgCreateDocBatch) (*sdk.Result, error) {

	for i := 0; i < len(msg.Holder); i++ {
		doc := types.Doc{Issuer: msg.Issuer, Holder: msg.Holder[i], Proof: msg.Proof[i], Data: msg.Data[i], Version: 0}

		// Check doc is existed
		existingDoc := k.GetDoc(ctx, doc)
		if len(existingDoc.Proof) > 0 {
			return nil, types.ErrDocExisted
		}

		k.SetDoc(ctx, doc)

		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeCreateDoc,
				sdk.NewAttribute(types.EventAttrIssuer, msg.Issuer.String()),
				sdk.NewAttribute(types.EventAttrHolder, string(msg.Holder[i])),
				sdk.NewAttribute(types.EventAttrProof, string(msg.Proof[i])),
			),
		)
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Issuer.String()),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.EventTypeCreateDoc),
		),
	)

	return &sdk.Result{
		Log:    msg.String(),
		Events: ctx.EventManager().Events(),
	}, nil
}

func handleMsgUpdateDoc(ctx sdk.Context, k Keeper, msg types.MsgUpdateDoc) (*sdk.Result, error) {
	queryDoc := types.Doc{Issuer: msg.Issuer, Holder: msg.Holder, Proof: msg.Proof, Data: msg.Data, Version: 0}

	// Check doc is existed
	existingDoc := k.GetDoc(ctx, queryDoc)
	if len(existingDoc.Proof) == 0 {
		return nil, types.ErrDocNotExisted
	}

	existingDoc.Data = msg.Data
	existingDoc.Version = existingDoc.Version + 1

	k.SetDoc(ctx, existingDoc)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeUpdateDoc,
			sdk.NewAttribute(types.EventAttrIssuer, msg.Issuer.String()),
			sdk.NewAttribute(types.EventAttrHolder, string(msg.Holder)),
			sdk.NewAttribute(types.EventAttrProof, string(msg.Proof)),
			sdk.NewAttribute(types.EventAttrData, string(msg.Data)),
		),
	)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Issuer.String()),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.EventTypeUpdateDoc),
		),
	)

	return &sdk.Result{
		Log:    msg.String(),
		Events: ctx.EventManager().Events(),
	}, nil
}

func handleMsgRevokeDoc(ctx sdk.Context, k Keeper, msg types.MsgRevokeDoc) (*sdk.Result, error) {
	queryDoc := types.Doc{Issuer: msg.Issuer, Holder: msg.Holder, Proof: msg.Proof}

	// Check doc is existed
	existingDoc := k.GetDoc(ctx, queryDoc)
	if len(existingDoc.Proof) == 0 {
		return nil, types.ErrDocNotExisted
	}

	existingDoc.Version = uint16(types.DocRevokeFlag)

	k.SetDoc(ctx, existingDoc)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeRevokeDoc,
			sdk.NewAttribute(types.EventAttrIssuer, msg.Issuer.String()),
			sdk.NewAttribute(types.EventAttrHolder, string(msg.Holder)),
			sdk.NewAttribute(types.EventAttrProof, string(msg.Proof)),
		),
	)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Issuer.String()),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.EventTypeRevokeDoc),
		),
	)

	return &sdk.Result{
		Log:    msg.String(),
		Events: ctx.EventManager().Events(),
	}, nil
}
