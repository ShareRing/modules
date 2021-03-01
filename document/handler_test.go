package document

import (
	"testing"

	keep "github.com/ShareRing/modules/document/keeper"
	"github.com/ShareRing/modules/document/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// nolint:deadcode,unused,varcheck
var (
	issuerAddr = sdk.AccAddress("issuer")

	extraData = ("extraData")
	proof     = ("proof-top-not-roof-top")
)

func Test_MsgCreateDocument_OK(t *testing.T) {
	ctx, keeper := keep.CreateTestInput(t, false)
	holder := ("id-123")

	msgCreateDoc1 := types.MsgCreateDoc{Holder: holder, Issuer: issuerAddr, Proof: proof, Data: extraData}
	res, err := handleMsgCreateDoc(ctx, keeper, msgCreateDoc1)

	require.Nil(t, err)
	require.NotNil(t, res)

	require.Equal(t, types.EventTypeCreateDoc, res.Events[0].Type)

	queryDoc := types.Doc{Proof: proof}
	doc := keeper.GetDocByProof(ctx, queryDoc)

	require.Equal(t, proof, doc.Proof)
	require.Equal(t, issuerAddr, doc.Issuer)
	require.Equal(t, holder, doc.Holder)
	require.Equal(t, extraData, doc.Data)

}
func Test_MsgCreateDocument_Duplicate(t *testing.T) {
	ctx, keeper := keep.CreateTestInput(t, false)
	holder := ("id-123")

	msgCreateDoc1 := types.MsgCreateDoc{Holder: holder, Issuer: issuerAddr, Proof: proof, Data: extraData}
	res, err := handleMsgCreateDoc(ctx, keeper, msgCreateDoc1)

	require.Nil(t, err)
	require.NotNil(t, res)

	msgCreateDoc2 := types.MsgCreateDoc{Holder: holder, Issuer: issuerAddr, Proof: proof, Data: extraData}
	res2, err2 := handleMsgCreateDoc(ctx, keeper, msgCreateDoc2)

	require.Nil(t, res2)
	require.Error(t, err2)
	require.Equal(t, types.ErrDocExisted, err2)

	require.Equal(t, types.EventTypeCreateDoc, res.Events[0].Type)

	queryDoc := types.Doc{Proof: proof}
	doc := keeper.GetDocByProof(ctx, queryDoc)

	require.Equal(t, proof, doc.Proof)
	require.Equal(t, issuerAddr, doc.Issuer)
	require.Equal(t, holder, doc.Holder)
	require.Equal(t, extraData, doc.Data)
}

func Test_MsgCreateDocumentInBatch_OK(t *testing.T) {
	ctx, keeper := keep.CreateTestInput(t, false)
	holders := []string{("id-1"), ("id-2"), ("id-3")}
	proofs := []string{("proof-1"), ("proof-2"), ("proof-3")}
	datas := []string{("data-1"), ("data-2"), ("data-3")}

	msgCreateDocInBatch := types.MsgCreateDocBatch{Holder: holders, Issuer: issuerAddr, Proof: proofs, Data: datas}
	res, err := handleMsgCreateDocInBatch(ctx, keeper, msgCreateDocInBatch)

	require.Nil(t, err)
	require.NotNil(t, res)

	require.Equal(t, types.EventTypeCreateDoc, res.Events[0].Type)
	require.Equal(t, 4, len(res.Events))

	for i := 0; i < len(holders); i++ {
		queryDoc := types.Doc{Proof: proofs[i]}
		doc := keeper.GetDocByProof(ctx, queryDoc)

		require.Equal(t, issuerAddr, doc.Issuer)
		require.Equal(t, proofs[i], doc.Proof)
		require.Equal(t, holders[i], doc.Holder)
		require.Equal(t, datas[i], doc.Data)
	}
}

func Test_MsgUpdateDocument_OK(t *testing.T) {
	ctx, keeper := keep.CreateTestInput(t, false)
	holder := ("id-123")

	msgCreateDoc1 := types.MsgCreateDoc{Holder: holder, Issuer: issuerAddr, Proof: proof, Data: extraData}
	res, err := handleMsgCreateDoc(ctx, keeper, msgCreateDoc1)

	require.Nil(t, err)
	require.NotNil(t, res)

	require.Equal(t, types.EventTypeCreateDoc, res.Events[0].Type)

	queryDoc := types.Doc{Proof: proof}
	doc := keeper.GetDocByProof(ctx, queryDoc)

	require.Equal(t, proof, doc.Proof)
	require.Equal(t, issuerAddr, doc.Issuer)
	require.Equal(t, holder, doc.Holder)
	require.Equal(t, extraData, doc.Data)

	newData := ("new data")
	msgUpdate := types.MsgUpdateDoc{Holder: holder, Issuer: issuerAddr, Proof: proof, Data: newData}
	res1, err1 := handleMsgUpdateDoc(ctx, keeper, msgUpdate)

	updatedDoc := keeper.GetDocByProof(ctx, queryDoc)
	require.Nil(t, err1)
	require.NotNil(t, res1)

	// TODO Check event
	// require.Equal(t, types.EventTypeUpdateDoc, res1.Events[0].Type)

	require.Equal(t, proof, updatedDoc.Proof)
	require.Equal(t, issuerAddr, updatedDoc.Issuer)
	require.Equal(t, holder, updatedDoc.Holder)
	require.Equal(t, newData, updatedDoc.Data)

	// Verion must be increased after updated
	newVer := doc.Version + uint16(1)
	require.Equal(t, newVer, updatedDoc.Version)

}

func Test_MsgUpdateDocument_DoesNotExist(t *testing.T) {
	ctx, keeper := keep.CreateTestInput(t, false)
	holder := ("id-123")

	newData := ("new data")
	msgUpdate := types.MsgUpdateDoc{Holder: holder, Issuer: issuerAddr, Proof: proof, Data: newData}
	res, err := handleMsgUpdateDoc(ctx, keeper, msgUpdate)

	require.Nil(t, res)
	require.Error(t, err)
	require.True(t, types.ErrDocNotExisted.Is(err))
}

func Test_MsgRevokeDocument_OK(t *testing.T) {
	ctx, keeper := keep.CreateTestInput(t, false)
	holder := ("id-123")

	msgCreateDoc1 := types.MsgCreateDoc{Holder: holder, Issuer: issuerAddr, Proof: proof, Data: extraData}
	res, err := handleMsgCreateDoc(ctx, keeper, msgCreateDoc1)

	require.Nil(t, err)
	require.NotNil(t, res)

	require.Equal(t, types.EventTypeCreateDoc, res.Events[0].Type)

	queryDoc := types.Doc{Proof: proof}
	doc := keeper.GetDocByProof(ctx, queryDoc)

	require.Equal(t, proof, doc.Proof)
	require.Equal(t, issuerAddr, doc.Issuer)
	require.Equal(t, holder, doc.Holder)
	require.Equal(t, extraData, doc.Data)

	msgRevokeDoc := types.MsgRevokeDoc{Issuer: issuerAddr, Proof: proof, Holder: holder}

	res2, err1 := handleMsgRevokeDoc(ctx, keeper, msgRevokeDoc)
	require.Nil(t, err1)
	require.NotNil(t, res2)

	// TODO Check event more
	// require.Equal(t, types.EventTypeRevokeDoc, res2.Events[0].Type)

	revokedDoc := keeper.GetDocByProof(ctx, queryDoc)

	require.Equal(t, proof, revokedDoc.Proof)
	require.Equal(t, issuerAddr, revokedDoc.Issuer)
	require.Equal(t, holder, revokedDoc.Holder)
	require.Equal(t, extraData, revokedDoc.Data)
	require.Equal(t, uint16(types.DocRevokeFlag), revokedDoc.Version)
}

func Test_MsgRevokeDocument_DoesNotExist(t *testing.T) {
	ctx, keeper := keep.CreateTestInput(t, false)
	holder := ("id-123")

	msgCreateDoc1 := types.MsgCreateDoc{Holder: holder, Issuer: issuerAddr, Proof: proof, Data: extraData}
	res, err := handleMsgCreateDoc(ctx, keeper, msgCreateDoc1)

	require.Nil(t, err)
	require.NotNil(t, res)

	require.Equal(t, types.EventTypeCreateDoc, res.Events[0].Type)

	queryDoc := types.Doc{Proof: proof}
	doc := keeper.GetDocByProof(ctx, queryDoc)

	require.Equal(t, proof, doc.Proof)
	require.Equal(t, issuerAddr, doc.Issuer)
	require.Equal(t, holder, doc.Holder)
	require.Equal(t, extraData, doc.Data)

	msgRevokeDoc := types.MsgRevokeDoc{Issuer: issuerAddr, Proof: proof + "1", Holder: holder}

	res2, err2 := handleMsgRevokeDoc(ctx, keeper, msgRevokeDoc)

	require.Nil(t, res2)
	require.Error(t, err2)
	require.True(t, types.ErrDocNotExisted.Is(err2))

	// TODO Check event more
	// require.Equal(t, types.EventTypeRevokeDoc, res2.Events[0].Type)

	revokedDoc := keeper.GetDocByProof(ctx, queryDoc)

	require.Equal(t, proof, revokedDoc.Proof)
	require.Equal(t, issuerAddr, revokedDoc.Issuer)
	require.Equal(t, holder, revokedDoc.Holder)
	require.Equal(t, extraData, revokedDoc.Data)
	require.Equal(t, uint16(0x0), revokedDoc.Version)
}
