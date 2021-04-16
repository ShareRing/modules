package rest

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/gorilla/mux"
)

// RegisterRoutes - Central function to define routes that get registered by the main application
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router, storeKey string) {
	r.HandleFunc(fmt.Sprintf("/%s/proof/{proof}", storeKey), QueryDocumentByProofRequestHandlerFn(cliCtx)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/holderid/{id}", storeKey), QueryDocumentByHolderIDRequestHandlerFn(cliCtx)).Methods("GET")
}
