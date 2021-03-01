package keeper

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authexported "github.com/cosmos/cosmos-sdk/x/auth/exported"
	"github.com/cosmos/cosmos-sdk/x/bank"

	"github.com/ShareRing/modules/id/types"
	"github.com/cosmos/cosmos-sdk/x/supply"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"
)

// Copy from x/staking/keeper/test_common.go
// create a codec used only for testing
func MakeTestCodec() *codec.Codec {
	var cdc = codec.New()

	// Register Msgs
	cdc.RegisterInterface((*sdk.Msg)(nil), nil)
	cdc.RegisterConcrete(bank.MsgSend{}, "test/staking/Send", nil)
	cdc.RegisterConcrete(types.MsgCreateId{}, "test/id/MsgCreateId", nil)

	// Register AppAccount
	cdc.RegisterInterface((*authexported.Account)(nil), nil)
	cdc.RegisterConcrete(&auth.BaseAccount{}, "test/staking/BaseAccount", nil)
	supply.RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)

	return cdc
}

// Hogpodge of all sorts of input required for testing.
func CreateTestInput(t *testing.T, isCheckTx bool) (sdk.Context, Keeper) {

	keyId := sdk.NewKVStoreKey(types.StoreKey)

	db := dbm.NewMemDB()
	ms := store.NewCommitMultiStore(db)

	ms.MountStoreWithDB(keyId, sdk.StoreTypeIAVL, db)

	err := ms.LoadLatestVersion()
	require.Nil(t, err)

	ctx := sdk.NewContext(ms, abci.Header{ChainID: "foochainid"}, isCheckTx, log.NewNopLogger())
	// ctx = ctx.WithConsensusParams(
	// 	&abci.ConsensusParams{
	// 		Validator: &abci.ValidatorParams{
	// 			PubKeyTypes: []string{tmtypes.ABCIPubKeyTypeEd25519},
	// 		},
	// 	},
	// )
	cdc := MakeTestCodec()

	// pk := params.NewKeeper(cdc, keyParams, tkeyParams)

	keeper := NewKeeper(cdc, keyId, nil)

	return ctx, keeper
}
