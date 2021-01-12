package cli

import (
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"

	"bitbucket.org/shareringvietnam/shareledger-modules/document/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"
)

// GetQueryCmd returns the parent command for all document CLi query commands. The
// provided clientCtx should have, at a minimum, a verifier, Tendermint RPC client,
// and marshaler set.
func GetQueryCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the document module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		GetDocByProofCmd(types.QuerierRoute, cdc),
	)

	return cmd
}

func GetDocByProofCmd(queryRoute string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "by-proof <proof>",
		Short: "Query for doc information",
		Long: strings.TrimSpace(fmt.Sprintf(`
Query document information by the proof.
Example:
$ %s query %s by-proof 5wpluxhf4qru2ewy58kc3w4tkzm3v`, version.Name, types.ModuleName)),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			var bz []byte
			var err error
			bz, err = createDocByAddressParams(&cliCtx, args[1])

			if err != nil {
				return err
			}

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", queryRoute, "proof"), bz)
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

func createDocByAddressParams(cliCtx *context.CLIContext, proof string) ([]byte, error) {

	params := types.QueryDocByProofParams{Proof: []byte(proof)}
	bz, err := cliCtx.Codec.MarshalJSON(params)
	return bz, err
}
