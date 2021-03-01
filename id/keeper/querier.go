package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/ShareRing/modules/id/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		case types.QueryInfo:
			return queryIdInfo(ctx, path[1:], req, k)

		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unknown query path: %s", path[0])
		}
	}
}

// Get Id by address or id
func queryIdInfo(ctx sdk.Context, path []string, req abci.RequestQuery, k Keeper) ([]byte, error) {
	id := &types.ID{}

	if path[0] == types.QueryPathAddress {
		// Get id by owner's address
		var params types.QueryIdByAddressParams
		err := types.ModuleCdc.UnmarshalJSON(req.Data, &params)

		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
		}

		id = k.GetIdByAddress(ctx, params.Address)
	} else if path[0] == types.QueryPathId {
		// Get id by id
		var params types.QueryIdByIdParams
		err := types.ModuleCdc.UnmarshalJSON(req.Data, &params)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
		}

		id = k.GetIDByIdString(ctx, params.Id)
	}

	// Return empty id if the id does not exist
	if id == nil {
		id = &types.ID{}
	}

	bz, err := codec.MarshalJSONIndent(types.ModuleCdc, id)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}
