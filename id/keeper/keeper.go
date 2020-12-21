package keeper

import (
	i "bitbucket.org/shareringvietnam/shareledger-modules/id/interfaces"
	"bitbucket.org/shareringvietnam/shareledger-modules/id/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Keeper struct {
	storeKey sdk.StoreKey
	cdc      *codec.Codec
	gmKeeper i.IGentlemintKeeper
}

func NewKeeper(cdc *codec.Codec, key sdk.StoreKey, gmKeeper i.IGentlemintKeeper) Keeper {
	return Keeper{
		storeKey: key,
		cdc:      cdc,
		gmKeeper: gmKeeper,
	}
}

func (k Keeper) GetIdByAddress(ctx sdk.Context, ownerAddr sdk.AccAddress) *types.ID {
	store := ctx.KVStore(k.storeKey)

	id := store.Get(ownerAddr)

	if len(id) == 0 {
		// TODO
		return nil
	}

	ids := k.GetBaseID(ctx, id)
	rs := types.NewIDFromStorage(string(id), ids)
	return &rs
}

func (k Keeper) GetIDByIdString(ctx sdk.Context, id string) *types.ID {
	ids := k.GetBaseID(ctx, []byte(id))
	rs := types.NewIDFromStorage(id, ids)
	return &rs
}

func (k Keeper) GetBaseID(ctx sdk.Context, id []byte) types.BaseID {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(id)

	if len(bz) == 0 {
		// TODO
	}

	ids := types.MustUnmarshalBaseID(k.cdc, bz)
	return ids
}

func (k Keeper) SetID(ctx sdk.Context, id types.ID) {
	store := ctx.KVStore(k.storeKey)

	// address -> id
	store.Set(id.OwnerAddr, []byte(id.Id))

	// id -> {ID}
	store.Set([]byte(id.Id), types.MustMarshalBaseID(k.cdc, id.ToStorage()))
}
