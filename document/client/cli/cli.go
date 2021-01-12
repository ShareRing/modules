package cli

import (
	"bufio"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"

	"bitbucket.org/shareringvietnam/shareledger-modules/document/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
)

// NewTxCmd returns a root CLI command handler for all id transaction commands.
func NewTxCmd(cdc *codec.Codec) *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Document transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	txCmd.AddCommand(
		NewDocumentTxCmd(cdc),
	)

	return txCmd
}

// NewDocumentTxCmd returns a CLI command handler for creating a MsgCreateDoc transaction.
func NewDocumentTxCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create [id] [backup_address] [owner_address] [extra_data]",
		Short: `Create new Id`,
		Long: `Create a new Id by given information.
eg: 
$ create uid-159654 shareledger1s432u6zv95wpluxhf4qru2ewy58kc3w4tkzm3v shareledger1s432u6zv95wpluxhf4qru2ewy58kc3w4tkzm3v shareledger1s432u6zv95wpluxhf4qru2ewy58kc3w4tkzm3v http://sharering.network
		`,
		Args: cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)

			// build and sign the transaction, then broadcast to Tendermint
			msg := types.MsgCreateDoc{}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	cmd = flags.PostCommands(cmd)[0]

	return cmd
}
