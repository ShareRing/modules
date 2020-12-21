package cli

import (
	"bufio"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"

	"bitbucket.org/shareringvietnam/shareledger-modules/id/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
)

// NewTxCmd returns a root CLI command handler for all x/bank transaction commands.
func NewTxCmd(cdc *codec.Codec) *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Id transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	txCmd.AddCommand(NewIdTxCmd(cdc))

	return txCmd
}

// NewSendTxCmd returns a CLI command handler for creating a MsgSend transaction.
func NewIdTxCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create [backup_address] [owner_address] [extra_data]",
		Short: `Create new Id`,
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContextWithInputAndFrom(inBuf, args[0]).WithCodec(cdc)

			backupAddr, err := sdk.AccAddressFromBech32(args[1])
			if err != nil {
				return err
			}

			ownerAddr, err := sdk.AccAddressFromBech32(args[2])
			if err != nil {
				return err
			}

			extraData := args[3]

			// build and sign the transaction, then broadcast to Tendermint
			msg := types.NewMsgCreateId(cliCtx.GetFromAddress(), backupAddr, ownerAddr, extraData)

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	cmd = flags.PostCommands(cmd)[0]

	return cmd
}
