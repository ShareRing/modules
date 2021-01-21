package document

import (
	"bitbucket.org/shareringvietnam/shareledger-modules/document/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type GenesisState struct {
	Documents []types.Doc
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
	for _, doc := range data.Documents {
		keeper.SetDoc(ctx, doc)
	}
}

func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	docs := []types.Doc{}

	cb := func(doc types.Doc) (stop bool) {
		docs = append(docs, doc)
		return false
	}

	k.IterateDocs(ctx, cb)

	return GenesisState{
		Documents: docs,
	}
}
