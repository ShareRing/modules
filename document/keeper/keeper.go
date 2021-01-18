package keeper

import (
	types "bitbucket.org/shareringvietnam/shareledger-modules/document/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Keeper struct {
	storeKey sdk.StoreKey
	cdc      *codec.Codec
}

func NewKeeper(cdc *codec.Codec, key sdk.StoreKey) Keeper {
	return Keeper{
		storeKey: key,
		cdc:      cdc,
	}
}

func (k Keeper) SetDoc(ctx sdk.Context, doc types.Doc) {
	store := ctx.KVStore(k.storeKey)

	// Doc detail
	store.Set(doc.GetKeyDetailState(), types.MustMarshalDocDetailState(k.cdc, doc.GetDetailState()))

	// Doc basic for easy query
	store.Set(doc.GetKeyBasicState(), types.MustMarshalDocBasicState(k.cdc, doc.GetBasicState()))
}

func (k Keeper) GetDocByProof(ctx sdk.Context, queryDoc types.Doc) types.Doc {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(queryDoc.GetKeyBasicState())
	if len(bz) == 0 {
		return types.Doc{}
	}

	dbs := types.MustUnmarshalDocBasicState(k.cdc, bz)
	queryDoc.Holder = dbs.Holder
	queryDoc.Issuer = dbs.Issuer

	return k.GetDoc(ctx, queryDoc)
}

func (k Keeper) GetDoc(ctx sdk.Context, queryDoc types.Doc) types.Doc {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(queryDoc.GetKeyDetailState())
	if len(bz) == 0 {
		return types.Doc{}
	}

	ds := types.MustUnmarshalDocDetailState(k.cdc, bz)

	queryDoc.Version = ds.Version
	queryDoc.Data = ds.Data
	return queryDoc
}

func (k Keeper) IterateAllDocsOfAHolder(ctx sdk.Context, holderId string, cb func(doc types.Doc) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	queryDoc := types.Doc{Holder: holderId}
	it := sdk.KVStorePrefixIterator(store, queryDoc.GetKeyDetailOfHolder())

	defer it.Close()
	for ; it.Valid(); it.Next() {

		doc := types.MustMarshalFromDetailRawState(k.cdc, it.Key(), it.Value())
		if cb(doc) {
			break
		}
	}
}
