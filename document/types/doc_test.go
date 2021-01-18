package types

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestDoc_DetailKey(t *testing.T) {
	holder := "hodler 1"
	issuer := sdk.AccAddress([]byte("issuer"))
	proof := "proof-1"
	doc := Doc{Holder: holder, Issuer: issuer, Proof: proof}

	detailKey := doc.GetKeyDetailState()
	fmt.Println("detailKey " + string(detailKey))
}
