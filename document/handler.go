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
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, fmt.Sprintf("Unrecognized %s Msg type: %v", types.ModuleName, msg.Type()))
		}
	}
}

func handleMsgCreateDoc(ctx sdk.Context, k Keeper, msg types.MsgCreateDoc) (*sdk.Result, error) {
	doc := types.Doc{Issuer: msg.Issuer, Holder: msg.Holder, Proof: msg.Proof, Data: msg.Data, Version: 0}
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
