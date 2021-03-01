package id

import (
	"testing"

	keep "github.com/ShareRing/modules/id/keeper"
	"github.com/ShareRing/modules/id/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/tendermint/crypto/secp256k1"
)

// nolint:deadcode,unused,varcheck
var (
	issuerPriv = secp256k1.GenPrivKey()
	issuerAddr = sdk.AccAddress(issuerPriv.PubKey().Address())

	backupPriv = secp256k1.GenPrivKey()
	backupAddr = sdk.AccAddress(issuerPriv.PubKey().Address())

	ownerPriv = secp256k1.GenPrivKey()
	ownerAddr = sdk.AccAddress(issuerPriv.PubKey().Address())

	priv1 = secp256k1.GenPrivKey()
	addr1 = sdk.AccAddress(issuerPriv.PubKey().Address())

	extraData = "extraData"
)

func Test_MsgCreateID_Dupdicate_ID(t *testing.T) {
	ctx, keeper := keep.CreateTestInput(t, false)
	// Pre insert id
	id1 := "id1"
	msgCreateId1 := types.NewMsgCreateId(issuerAddr, backupAddr, ownerAddr, id1, extraData)
	res, err := handleMsgCreateId(ctx, keeper, msgCreateId1)

	// Can not create the same id
	msgCreateId2 := types.NewMsgCreateId(issuerAddr, backupAddr, addr1, id1, extraData)

	res, err = handleMsgCreateId(ctx, keeper, msgCreateId2)
	require.Error(t, err)
	require.Nil(t, res)
	require.Equal(t, err.Error(), ErrIdExisted.Error())
}
func Test_MsgCreateID_Dupdicate_Owner(t *testing.T) {
	ctx, keeper := keep.CreateTestInput(t, false)

	// Pre insert id
	id1 := "id1"
	msgCreateId1 := types.NewMsgCreateId(issuerAddr, backupAddr, ownerAddr, id1, extraData)
	res, err := handleMsgCreateId(ctx, keeper, msgCreateId1)

	// Can not create id if owner has id
	id2 := "id2"
	msgCreateId3 := types.NewMsgCreateId(issuerAddr, backupAddr, ownerAddr, id2, extraData)

	res, err = handleMsgCreateId(ctx, keeper, msgCreateId3)
	require.Error(t, err)
	require.Nil(t, res)
	require.Equal(t, err.Error(), ErrIdExisted.Error())
}

func Test_CreateIDBatch_Succeed(t *testing.T) {
	ctx, keeper := keep.CreateTestInput(t, false)

	listId := []string{"id-batch-1", "id-batch-2", "id-batch-3"}
	listExtraData := []string{"id-batch-1", "id-batch-2", "id-batch-3"}
	_, listBackupAddr := createRandomAddr(3)
	_, listOwnerAddr := createRandomAddr(3)

	msgCreateIdBath1 := types.MsgCreateIdBatch{
		IssuerAddr: issuerAddr,
		Id:         listId,
		BackupAddr: listBackupAddr,
		OwnerAddr:  listOwnerAddr,
		ExtraData:  listExtraData,
	}

	res, err := handleMsgCreateIdBatch(ctx, keeper, msgCreateIdBath1)

	require.Nil(t, err)
	require.NotNil(t, res)

	//TODO: Check events
}

func TestCreateIDBatchDuplicateOwner(t *testing.T) {
	ctx, keeper := keep.CreateTestInput(t, false)

	dupOwner := sdk.AccAddress(secp256k1.GenPrivKey().PubKey().Address())
	listId := []string{"id-batch-1", "id-batch-2", "id-batch-3"}
	listExtraData := []string{"id-batch-1", "id-batch-2", "id-batch-3"}
	_, listBackupAddr := createRandomAddr(3)
	_, listOwnerAddr := createRandomAddr(1)

	listOwnerAddr = append(listOwnerAddr, dupOwner, dupOwner)

	msgCreateIdBath1 := types.MsgCreateIdBatch{
		IssuerAddr: issuerAddr,
		Id:         listId,
		BackupAddr: listBackupAddr,
		OwnerAddr:  listOwnerAddr,
		ExtraData:  listExtraData,
	}

	res, err := handleMsgCreateIdBatch(ctx, keeper, msgCreateIdBath1)
	require.Error(t, err)
	require.Nil(t, res)
	require.Equal(t, err.Error(), ErrIdExisted.Error())
}

func TestCreateIDBatchExistingId(t *testing.T) {
	ctx, keeper := keep.CreateTestInput(t, false)

	// Pre insert 1 id
	id1 := "id1"
	msgCreateId1 := types.NewMsgCreateId(issuerAddr, backupAddr, ownerAddr, id1, extraData)
	_, err := handleMsgCreateId(ctx, keeper, msgCreateId1)
	require.Nil(t, err)

	// Try to create new id with the same id
	listId := []string{"id-batch-1", "id1", "id-batch-3"}
	listExtraData := []string{"id-batch-1", "id-batch-2", "id-batch-3"}
	_, listBackupAddr := createRandomAddr(3)
	_, listOwnerAddr := createRandomAddr(3)

	msgCreateIdBath1 := types.MsgCreateIdBatch{
		IssuerAddr: issuerAddr,
		Id:         listId,
		BackupAddr: listBackupAddr,
		OwnerAddr:  listOwnerAddr,
		ExtraData:  listExtraData,
	}

	{
		res, err := handleMsgCreateIdBatch(ctx, keeper, msgCreateIdBath1)
		require.Error(t, err)
		require.Nil(t, res)
		require.Equal(t, err.Error(), ErrIdExisted.Error())
	}

	// Try to create new id with the same owner
	{
		listId := []string{"id-batch-1", "id-batch-2", "id-batch-3"}
		listExtraData := []string{"id-batch-1", "id-batch-2", "id-batch-3"}
		_, listBackupAddr := createRandomAddr(3)
		_, listOwnerAddr := createRandomAddr(2)
		listOwnerAddr = append(listOwnerAddr, ownerAddr)

		msgCreateIdBath1 := types.MsgCreateIdBatch{
			IssuerAddr: issuerAddr,
			Id:         listId,
			BackupAddr: listBackupAddr,
			OwnerAddr:  listOwnerAddr,
			ExtraData:  listExtraData,
		}

		{
			res, err := handleMsgCreateIdBatch(ctx, keeper, msgCreateIdBath1)
			require.Error(t, err)
			require.Nil(t, res)
			require.Equal(t, err.Error(), ErrIdExisted.Error())
		}
	}
}
func TestCreateIDBatchExistingOwner(t *testing.T) {
	ctx, keeper := keep.CreateTestInput(t, false)

	// Pre insert 1 id
	id1 := "id1"
	msgCreateId1 := types.NewMsgCreateId(issuerAddr, backupAddr, ownerAddr, id1, extraData)
	_, err := handleMsgCreateId(ctx, keeper, msgCreateId1)
	require.Nil(t, err)

	// Try to create new id with the same owner
	{
		listId := []string{"id-batch-1", "id-batch-2", "id-batch-3"}
		listExtraData := []string{"id-batch-1", "id-batch-2", "id-batch-3"}
		_, listBackupAddr := createRandomAddr(3)
		_, listOwnerAddr := createRandomAddr(2)
		listOwnerAddr = append(listOwnerAddr, ownerAddr)

		msgCreateIdBath1 := types.MsgCreateIdBatch{
			IssuerAddr: issuerAddr,
			Id:         listId,
			BackupAddr: listBackupAddr,
			OwnerAddr:  listOwnerAddr,
			ExtraData:  listExtraData,
		}

		{
			res, err := handleMsgCreateIdBatch(ctx, keeper, msgCreateIdBath1)
			require.Error(t, err)
			require.Nil(t, res)
			require.Equal(t, err.Error(), ErrIdExisted.Error())
		}
	}
}

func TestCreateIDBatchDuplicateId(t *testing.T) {
	ctx, keeper := keep.CreateTestInput(t, false)

	listId := []string{"id-batch-1", "id-batch-1", "id-batch-3"}
	listExtraData := []string{"id-batch-1", "id-batch-1", "id-batch-3"}
	_, listBackupAddr := createRandomAddr(3)
	_, listOwnerAddr := createRandomAddr(3)

	msgCreateIdBath1 := types.MsgCreateIdBatch{
		IssuerAddr: issuerAddr,
		Id:         listId,
		BackupAddr: listBackupAddr,
		OwnerAddr:  listOwnerAddr,
		ExtraData:  listExtraData,
	}

	res, err := handleMsgCreateIdBatch(ctx, keeper, msgCreateIdBath1)
	require.Error(t, err)
	require.Nil(t, res)
	require.Equal(t, err.Error(), ErrIdExisted.Error())
}

func Test_MsgUpdateID_NotExist(t *testing.T) {
	ctx, keeper := keep.CreateTestInput(t, false)

	id1 := "id1"
	msgUpdateId2 := types.NewMsgUpdateId(issuerAddr, id1, extraData)

	res, err := handleMsgUpdateId(ctx, keeper, msgUpdateId2)
	require.Error(t, err)
	require.Nil(t, res)
	require.Equal(t, err.Error(), ErrIdNotExisted.Error())
}

func Test_MsgUpdateID_Succeed(t *testing.T) {
	ctx, keeper := keep.CreateTestInput(t, false)

	// Pre insert 1 id
	id1 := "id1"
	extraData2 := "extraData2"
	msgCreateId1 := types.NewMsgCreateId(issuerAddr, backupAddr, ownerAddr, id1, extraData)
	_, err := handleMsgCreateId(ctx, keeper, msgCreateId1)
	require.Nil(t, err)

	msgUpdateId2 := types.NewMsgUpdateId(issuerAddr, id1, extraData2)

	res, err := handleMsgUpdateId(ctx, keeper, msgUpdateId2)

	require.Nil(t, err)
	require.NotNil(t, res)

	id := keeper.GetIDByIdString(ctx, id1)

	require.Equal(t, extraData2, id.ExtraData)
}
func Test_MsgReplaceOwner_NotExist(t *testing.T) {
	ctx, keeper := keep.CreateTestInput(t, false)

	id1 := "id1"
	// Can not create the same id
	msgReplaceId := types.NewMsgReplaceIdOwner(id1, ownerAddr, backupAddr)

	res, err := handleMsgReplaceOwnerId(ctx, keeper, msgReplaceId)
	require.Error(t, err)
	require.Nil(t, res)
	require.Equal(t, err.Error(), ErrIdNotExisted.Error())
}

func Test_MsgReplaceOwner_NewOwner_HasId(t *testing.T) {
	ctx, keeper := keep.CreateTestInput(t, false)

	_, ownerAddrs := createRandomAddr(2)
	// Pre insert id
	id1 := "id1"
	msgCreateId1 := types.NewMsgCreateId(issuerAddr, backupAddr, ownerAddrs[0], id1, extraData)
	res, err := handleMsgCreateId(ctx, keeper, msgCreateId1)

	require.Nil(t, err)
	require.NotNil(t, res)

	id2 := "id2"
	msgCreateId2 := types.NewMsgCreateId(issuerAddr, backupAddr, ownerAddrs[1], id2, extraData)
	res, err = handleMsgCreateId(ctx, keeper, msgCreateId2)
	require.Nil(t, err)
	require.NotNil(t, res)

	// Replace owner of id1: ownerAddrs[0]-->ownerAddrs[1]
	msgReplaceId := types.NewMsgReplaceIdOwner(id1, ownerAddrs[1], backupAddr)

	res, err = handleMsgReplaceOwnerId(ctx, keeper, msgReplaceId)
	require.Error(t, err)
	require.Nil(t, res)
	require.Equal(t, ErrIdExisted.Error(), err.Error())
}

func Test_MsgReplaceOwner_Succeed(t *testing.T) {
	ctx, keeper := keep.CreateTestInput(t, false)

	_, ownerAddrs := createRandomAddr(2)
	// Pre insert id
	id1 := "id1"
	msgCreateId1 := types.NewMsgCreateId(issuerAddr, backupAddr, ownerAddrs[0], id1, extraData)
	res, err := handleMsgCreateId(ctx, keeper, msgCreateId1)

	require.Nil(t, err)
	require.NotNil(t, res)

	// Replace owner of id1: ownerAddrs[0]-->ownerAddrs[1]
	msgReplaceId := types.NewMsgReplaceIdOwner(id1, ownerAddrs[1], backupAddr)

	res, err = handleMsgReplaceOwnerId(ctx, keeper, msgReplaceId)
	require.Nil(t, err)
	require.NotNil(t, res)
}

func createRandomAddr(amount int) (prvs []secp256k1.PrivKeySecp256k1, addrs []sdk.AccAddress) {
	for i := 0; i < amount; i++ {
		prv := secp256k1.GenPrivKey()
		addr := sdk.AccAddress(prv.PubKey().Address())

		addrs = append(addrs, addr)
		prvs = append(prvs, prv)
	}
	return
}
