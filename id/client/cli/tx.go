package cli

import (
	"bufio"
	"strings"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"

	"github.com/ShareRing/modules/id/types"
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
		Short:                      "Id transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	txCmd.AddCommand(
		NewIdTxCmd(cdc),
		NewIdBatchTxCmd(cdc),
		UpdateIdTxCmd(cdc),
		UpdateReplaceIdownerTxCmd(cdc),
	)

	return txCmd
}

// NewIdTxCmd returns a CLI command handler for creating a MsgCreateId transaction.
func NewIdTxCmd(cdc *codec.Codec) *cobra.Command {
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

			id := args[0]

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
			msg := types.NewMsgCreateId(cliCtx.GetFromAddress(), backupAddr, ownerAddr, id, extraData)

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	cmd = flags.PostCommands(cmd)[0]

	return cmd
}

func NewIdBatchTxCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-batch [id] [backup_address] [owner_address] [extra_data]",
		Short: `Create id by a batch`,
		Long: `Create multiple ids by a bath.
eg:
$ create-batch id id1,id2 shareledger1yyc7hnyjqwlxqy9rt8scwuttwa3nt4yn2wl0z0,shareledger1m3tfsfc5hkj78c9kk6k500e3p35tngnl3xymg4 shareledger17kejntc3f8njkermmpgtff0zg44ejlw5aw9jka,shareledger1e37j0radceaa5qp0akffe0j5kxwn9jxfy7ve9k shareledger17kejntc3f8njkermmpgtff0zg44ejlw5aw9jka,shareledger1e37j0radceaa5qp0akffe0j5kxwn9jxfy7ve9k manager,staff
		`,
		Args: cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)

			seperator := ","
			ids := strings.Split(args[0], seperator)
			backups := strings.Split(args[1], seperator)
			owners := strings.Split(args[2], seperator)
			extras := strings.Split(args[3], seperator)

			backupAddrs := make([]sdk.AccAddress, 0, len(ids))
			ownerAddrs := make([]sdk.AccAddress, 0, len(ids))

			for i := 0; i < len(ids); i++ {
				backupAddr, err := sdk.AccAddressFromBech32(backups[i])
				if err != nil {
					return err
				}

				ownerAddr, err := sdk.AccAddressFromBech32(owners[i])
				if err != nil {
					return err
				}

				backupAddrs = append(backupAddrs, backupAddr)
				ownerAddrs = append(ownerAddrs, ownerAddr)
			}

			// build and sign the transaction, then broadcast to Tendermint
			msg := types.NewMsgCreateIdBatch(cliCtx.GetFromAddress(), backupAddrs, ownerAddrs, ids, extras)

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	cmd = flags.PostCommands(cmd)[0]

	return cmd
}

// NewIdTxCmd returns a CLI command handler for creating a MsgUpdateId transaction.
func UpdateIdTxCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update [id] [extra_data]",
		Short: `Update Id`,
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)

			id := args[0]

			extraData := args[1]

			// build and sign the transaction, then broadcast to Tendermint
			msg := types.NewMsgUpdateId(cliCtx.GetFromAddress(), id, extraData)

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	cmd = flags.PostCommands(cmd)[0]

	return cmd
}

// NewIdTxCmd returns a CLI command handler for creating a MsgCreateId transaction.
func UpdateReplaceIdownerTxCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "replace [id] [new_owner_address]",
		Short: `Replace owner of an Id`,
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)

			id := args[0]

			newOwner, err := sdk.AccAddressFromBech32(args[1])
			if err != nil {
				return err
			}

			// build and sign the transaction, then broadcast to Tendermint
			msg := types.NewMsgReplaceIdOwner(id, newOwner, cliCtx.GetFromAddress())

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	cmd = flags.PostCommands(cmd)[0]

	return cmd
}
