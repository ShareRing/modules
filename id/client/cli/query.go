package cli

import (
	"errors"
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/ShareRing/modules/id/types"
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"
)

// GetQueryCmd returns the parent command for all x/bank CLi query commands. The
// provided clientCtx should have, at a minimum, a verifier, Tendermint RPC client,
// and marshaler set.
func GetQueryCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the id module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		GetIdByAddressCmd(types.QuerierRoute, cdc),
	)

	return cmd
}

func GetIdByAddressCmd(queryRoute string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "info <address>|<id> [address,[id]]",
		Short: "Query for id information",
		Long: strings.TrimSpace(fmt.Sprintf(`
Query id information of an account by owner address or the id.
Example:
$ %s query %s info address shareledger1s432u6zv95wpluxhf4qru2ewy58kc3w4tkzm3v
$ %s query %s info id 123e4567-e89b-12d3-a456-426655440000`, version.Name, types.ModuleName, version.Name, types.ModuleName)),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			var bz []byte
			var err error
			if args[0] == types.QueryPathAddress {
				bz, err = createGetIdByAddress(&cliCtx, args[1])
			} else if args[0] == types.QueryPathId {
				bz, err = createGetIdById(&cliCtx, args[1])
			} else {
				return errors.New("unknow command: " + args[0])
			}

			if err != nil {
				return err
			}

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/info/%s", queryRoute, args[0]), bz)
			if err != nil {
				return err
			}

			if len(res) == 0 {
				return fmt.Errorf("id not found")
			}

			var out types.ID
			cdc.MustUnmarshalJSON(res, &out)

			return cliCtx.PrintOutput(out)
		},
	}
	return flags.GetCommands(cmd)[0]
}

func createGetIdByAddress(cliCtx *context.CLIContext, bench32Addr string) ([]byte, error) {
	addr, addrErr := sdk.AccAddressFromBech32(bench32Addr)
	if addrErr != nil {
		return nil, addrErr
	}
	params := types.QueryIdByAddressParams{Address: addr}
	bz, err := cliCtx.Codec.MarshalJSON(params)
	return bz, err
}

func createGetIdById(cliCtx *context.CLIContext, id string) ([]byte, error) {
	params := types.QueryIdByIdParams{Id: id}
	bz, err := cliCtx.Codec.MarshalJSON(params)
	return bz, err
}
