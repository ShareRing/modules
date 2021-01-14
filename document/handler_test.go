package document

import (
	"testing"

	keep "bitbucket.org/shareringvietnam/shareledger-modules/document/keeper"
	"bitbucket.org/shareringvietnam/shareledger-modules/document/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// nolint:deadcode,unused,varcheck
var (
	issuerAddr = sdk.AccAddress("issuer")

	extraData = []byte("extraData")
	proof     = []byte("proof-top-not-roof-top")
)

func Test_MsgCreateDocument_OK(t *testing.T) {
	ctx, keeper := keep.CreateTestInput(t, false)
	holder := []byte("id-123")

	msgCreateDoc1 := types.MsgCreateDoc{Holder: holder, Issuer: issuerAddr, Proof: proof, Data: extraData}
	res, err := handleMsgCreateDoc(ctx, keeper, msgCreateDoc1)

	require.Nil(t, err)
	require.NotNil(t, res)

	require.Equal(t, types.EventTypeCreateDoc, res.Events[0].Type)

	queryDoc := types.Doc{Proof: proof}
	doc := keeper.GetDocByProof(ctx, queryDoc)

	require.ElementsMatch(t, proof, doc.Proof)
	require.ElementsMatch(t, issuerAddr, doc.Issuer)
	require.ElementsMatch(t, holder, doc.Holder)
	require.ElementsMatch(t, extraData, doc.Data)

}
func Test_MsgCreateDocument_Duplicate(t *testing.T) {
	ctx, keeper := keep.CreateTestInput(t, false)
	holder := []byte("id-123")

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

	require.ElementsMatch(t, proof, doc.Proof)
	require.ElementsMatch(t, issuerAddr, doc.Issuer)
	require.ElementsMatch(t, holder, doc.Holder)
	require.ElementsMatch(t, extraData, doc.Data)
}

func Test_MsgCreateDocumentInBatch_OK(t *testing.T) {
	ctx, keeper := keep.CreateTestInput(t, false)
	holders := []types.ID{types.ID("id-1"), types.ID("id-2"), types.ID("id-3")}
	proofs := [][]byte{[]byte("proof-1"), []byte("proof-2"), []byte("proof-3")}
	datas := [][]byte{[]byte("data-1"), []byte("data-2"), []byte("data-3")}

	msgCreateDocInBatch := types.MsgCreateDocBatch{Holder: holders, Issuer: issuerAddr, Proof: proofs, Data: datas}
	res, err := handleMsgCreateDocInBatch(ctx, keeper, msgCreateDocInBatch)

	require.Nil(t, err)
	require.NotNil(t, res)

	require.Equal(t, types.EventTypeCreateDoc, res.Events[0].Type)
	require.Equal(t, 4, len(res.Events))

	for i := 0; i < len(holders); i++ {
		queryDoc := types.Doc{Proof: proofs[i]}
		doc := keeper.GetDocByProof(ctx, queryDoc)

		require.ElementsMatch(t, issuerAddr, doc.Issuer)
		require.ElementsMatch(t, proofs[i], doc.Proof)
		require.ElementsMatch(t, holders[i], doc.Holder)
		require.ElementsMatch(t, datas[i], doc.Data)
	}
}

func Test_MsgUpdateDocument_OK(t *testing.T) {
	ctx, keeper := keep.CreateTestInput(t, false)
	holder := []byte("id-123")

	msgCreateDoc1 := types.MsgCreateDoc{Holder: holder, Issuer: issuerAddr, Proof: proof, Data: extraData}
	res, err := handleMsgCreateDoc(ctx, keeper, msgCreateDoc1)

	require.Nil(t, err)
	require.NotNil(t, res)

	require.Equal(t, types.EventTypeCreateDoc, res.Events[0].Type)

	queryDoc := types.Doc{Proof: proof}
	doc := keeper.GetDocByProof(ctx, queryDoc)

	require.ElementsMatch(t, proof, doc.Proof)
	require.ElementsMatch(t, issuerAddr, doc.Issuer)
	require.ElementsMatch(t, holder, doc.Holder)
	require.ElementsMatch(t, extraData, doc.Data)

	newData := []byte("new data")
	msgUpdate := types.MsgUpdateDoc{Holder: holder, Issuer: issuerAddr, Proof: proof, Data: newData}
	res1, err1 := handleMsgUpdateDoc(ctx, keeper, msgUpdate)

	updatedDoc := keeper.GetDocByProof(ctx, queryDoc)
	require.Nil(t, err1)
	require.NotNil(t, res1)

	// TODO Check event
	// require.Equal(t, types.EventTypeUpdateDoc, res1.Events[0].Type)

	require.ElementsMatch(t, proof, updatedDoc.Proof)
	require.ElementsMatch(t, issuerAddr, updatedDoc.Issuer)
	require.ElementsMatch(t, holder, updatedDoc.Holder)
	require.ElementsMatch(t, newData, updatedDoc.Data)

	// Verion must be increased after updated
	newVer := doc.Version + uint16(1)
	require.Equal(t, newVer, updatedDoc.Version)

}

func Test_MsgUpdateDocument_DoesNotExist(t *testing.T) {
	ctx, keeper := keep.CreateTestInput(t, false)
	holder := []byte("id-123")

	newData := []byte("new data")
	msgUpdate := types.MsgUpdateDoc{Holder: holder, Issuer: issuerAddr, Proof: proof, Data: newData}
	res, err := handleMsgUpdateDoc(ctx, keeper, msgUpdate)

	require.Nil(t, res)
	require.Error(t, err)
	require.True(t, types.ErrDocNotExisted.Is(err))
}

func Test_MsgRevokeDocument_OK(t *testing.T) {
	ctx, keeper := keep.CreateTestInput(t, false)
	holder := []byte("id-123")

	msgCreateDoc1 := types.MsgCreateDoc{Holder: holder, Issuer: issuerAddr, Proof: proof, Data: extraData}
	res, err := handleMsgCreateDoc(ctx, keeper, msgCreateDoc1)

	require.Nil(t, err)
	require.NotNil(t, res)

	require.Equal(t, types.EventTypeCreateDoc, res.Events[0].Type)

	queryDoc := types.Doc{Proof: proof}
	doc := keeper.GetDocByProof(ctx, queryDoc)

	require.ElementsMatch(t, proof, doc.Proof)
	require.ElementsMatch(t, issuerAddr, doc.Issuer)
	require.ElementsMatch(t, holder, doc.Holder)
	require.ElementsMatch(t, extraData, doc.Data)

	msgRevokeDoc := types.MsgRevokeDoc{Issuer: issuerAddr, Proof: proof, Holder: holder}

	res2, err1 := handleMsgRevokeDoc(ctx, keeper, msgRevokeDoc)
	require.Nil(t, err1)
	require.NotNil(t, res2)

	// TODO Check event more
	// require.Equal(t, types.EventTypeRevokeDoc, res2.Events[0].Type)

	revokedDoc := keeper.GetDocByProof(ctx, queryDoc)

	require.ElementsMatch(t, proof, revokedDoc.Proof)
	require.ElementsMatch(t, issuerAddr, revokedDoc.Issuer)
	require.ElementsMatch(t, holder, revokedDoc.Holder)
	require.ElementsMatch(t, extraData, revokedDoc.Data)
	require.Equal(t, types.DocRevokeFlag, revokedDoc.Version)
}
