package id

import (
	"encoding/json"

	"bitbucket.org/shareringvietnam/shareledger-modules/id/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type GenesisState struct {
	IDs []types.ID `json:"IDs" yaml:"IDs"`
}

func NewGenesisState() GenesisState {
	return GenesisState{}
}

func ValidateGenesis(data GenesisState) error {
	return nil
}

func DefaultGenesisState() GenesisState {
	return GenesisState{}
}

func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) {
	for _, id := range data.IDs {
		keeper.SetID(ctx, &id)
	}
}

func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	ids := []types.ID{}

	cb := func(id types.ID) (stop bool) {
		ids = append(ids, id)
		return false
	}

	k.IterateID(ctx, cb)

	return GenesisState{
		IDs: ids,
	}
}

func GetGenesisStateFromAppState(cdc *codec.Codec, appState map[string]json.RawMessage) GenesisState {
	var genesisState GenesisState
	if appState[ModuleName] != nil {
		cdc.MustUnmarshalJSON(appState[ModuleName], &genesisState)
	}

	return genesisState
}
