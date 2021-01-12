package rest

import (
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
)

// QueryIdByAddressRequestHandlerFn returns a REST handler that queries for all
// id information of and account or and id.
func QueryDocumentByProofRequestHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
