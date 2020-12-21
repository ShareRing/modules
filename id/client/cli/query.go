package cli

import (
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"

	"bitbucket.org/shareringvietnam/shareledger-modules/id/types"
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
		Use:   "id [address]",
		Short: "Query for id by owner address",
		Long:  strings.TrimSpace(fmt.Sprintf(`Query id information of an account by owner address.Example:$ %s query %s id [address]`, version.Name, types.ModuleName)),
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			addr, err := sdk.ValAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", queryRoute, "info"), nil)
			if err != nil {
				return err
			}

			if len(res) == 0 {
				return fmt.Errorf("no id found with address %s", addr)
			}

			var out types.MsgCreateId
			cdc.MustUnmarshalJSON(res, &out)

			return cliCtx.PrintOutput(out)
		},
	}
	return cmd
}
