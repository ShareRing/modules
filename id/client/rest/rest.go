package rest

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"

	"github.com/gorilla/mux"
)

// RegisterRoutes - Central function to define routes that get registered by the main application
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router, storeKey string) {
	r.HandleFunc(fmt.Sprintf("/%s/info/{path}/{address}", storeKey), QueryIdInfoRequestHandlerFn(cliCtx)).Methods("GET")
}
