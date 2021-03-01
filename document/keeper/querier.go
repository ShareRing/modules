package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/ShareRing/modules/document/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		case types.QueryByProof:
			return queryDocByProof(ctx, req, k)
		case types.QueryByHolder:
			return queryAllDocsOfAHolder(ctx, req, k)
		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unknown query path: %s", path[0])
		}
	}
}

// Get Doc by proof
func queryDocByProof(ctx sdk.Context, req abci.RequestQuery, k Keeper) ([]byte, error) {
	var params types.QueryDocByProofParams

	err := types.ModuleCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	// Return empty doc if the doc does not exist
	queryDoc := types.Doc{Proof: params.Proof}
	doc := k.GetDocByProof(ctx, queryDoc)

	bz, err := codec.MarshalJSONIndent(types.ModuleCdc, doc)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}

func queryAllDocsOfAHolder(ctx sdk.Context, req abci.RequestQuery, k Keeper) ([]byte, error) {
	var params types.QueryDocByHolderParams

	err := types.ModuleCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	// Return empty doc if the doc does not exist
	docs := make([]types.Doc, 0)

	cb := func(doc types.Doc) bool {
		docs = append(docs, doc)
		return false
	}

	k.IterateAllDocsOfAHolder(ctx, params.Holder, cb)

	// docsType := types.Docs(docs)

	bz, err := codec.MarshalJSONIndent(types.ModuleCdc, docs)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil

}
