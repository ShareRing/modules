package rest

import (
	"fmt"
	"net/http"

	"github.com/ShareRing/modules/document/types"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
)

// QueryDocumentByProofRequestHandlerFn returns a REST handler that queries for all
// document by proof
func QueryDocumentByProofRequestHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)

		var bz []byte

		proof := vars["proof"]

		if len(proof) == 0 || len(proof) > 64 {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "Invalid proof")
			return
		}

		params := types.NewQueryDocByProofParams(proof)
		bz, err := cliCtx.Codec.MarshalJSON(params)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}

		res, height, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/proof", types.QuerierRoute, proof), bz)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		cliCtx = cliCtx.WithHeight(height)

		// the query will return empty if there is no data for this account
		if len(res) == 0 {
			rest.PostProcessResponse(w, cliCtx, types.Doc{})
			return
		}

		rest.PostProcessResponse(w, cliCtx, res)
	}
}

func QueryDocumentByHolderIDRequestHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)

		var bz []byte

		holderID := vars["id"]

		if len(holderID) == 0 || len(holderID) > 64 {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "Invalid id")
			return
		}

		params := types.NewQueryDocByHolderParams(holderID)
		bz, err := cliCtx.Codec.MarshalJSON(params)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}

		res, height, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/holder", types.QuerierRoute), bz)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		cliCtx = cliCtx.WithHeight(height)

		if len(res) == 0 {
			rest.PostProcessResponse(w, cliCtx, []types.Doc{})
			return
		}

		rest.PostProcessResponse(w, cliCtx, res)
	}
}
