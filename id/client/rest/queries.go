package rest

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/ShareRing/modules/id/types"
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
)

// QueryIdByAddressRequestHandlerFn returns a REST handler that queries for all
// id information of and account or and id.
func QueryIdInfoRequestHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)

		path := vars["path"]
		var bz []byte

		if path == types.QueryPathAddress {
			bech32addr := vars["address"]
			addr, err := sdk.AccAddressFromBech32(bech32addr)
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
				return
			}
			params := types.NewQueryIdByAddressParams(addr)
			bz, err = cliCtx.Codec.MarshalJSON(params)
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
				return
			}
		} else if path == types.QueryPathId {
			params := types.NewQueryIdByIdParams(vars["address"])
			var err error
			bz, err = cliCtx.Codec.MarshalJSON(params)
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
				return
			}
		} else {
			rest.WriteErrorResponse(w, http.StatusBadRequest, errors.New("unknow endpoint").Error())
			return
		}

		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}

		res, height, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/info/%s", types.QuerierRoute, path), bz)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		cliCtx = cliCtx.WithHeight(height)

		// the query will return empty if there is no data for this account
		if len(res) == 0 {
			rest.PostProcessResponse(w, cliCtx, types.ID{})
			return
		}

		rest.PostProcessResponse(w, cliCtx, res)
	}
}
