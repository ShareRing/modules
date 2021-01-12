package document

import (
	"testing"

	keep "bitbucket.org/shareringvietnam/shareledger-modules/document/keeper"
	"bitbucket.org/shareringvietnam/shareledger-modules/document/types"
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

	extraData = []byte("extraData")
	proof     = []byte("proof-top-not-roof-top")
)

func Test_MsgCreateDocument_OK(t *testing.T) {
	ctx, keeper := keep.CreateTestInput(t, false)
	holder := []byte("id-123")

	msgCreateId1 := types.MsgCreateDoc{Holder: holder, Issuer: issuerAddr, Proof: proof, Data: extraData}
	res, err := handleMsgCreateDoc(ctx, keeper, msgCreateId1)

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
