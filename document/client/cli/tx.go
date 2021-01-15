package cli

import (
	"bufio"
	"strings"

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
		NewDocumentInBatchTxCmd(cdc),
		UpdateDocCmd(cdc),
		RevokeDocCmd(cdc),
	)

	return txCmd
}

// NewDocumentTxCmd returns a CLI command handler for creating a MsgCreateDoc transaction.
func NewDocumentTxCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create [holder id] [proof] [extra_data]",
		Short: `Create new docoment`,
		Long: `Create a new document by given information.
eg: 
$ create uid-159654 c89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc6 https://sharering.network/id/463"
		`,
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)

			issuer := cliCtx.GetFromAddress()
			holderId := args[0]
			proof := args[1]
			data := args[2]

			// build and sign the transaction, then broadcast to Tendermint
			msg := types.NewMsgCreateDoc(issuer, holderId, proof, data)

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	cmd = flags.PostCommands(cmd)[0]

	return cmd
}

func NewDocumentInBatchTxCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-batch [holder id] [proof] [extra_data]",
		Short: `Create new docoment`,
		Long: `Create a new document by given information.
eg: 
$ create uid-159654 c89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc6 https://sharering.network/id/463"
		`,
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)

			issuer := cliCtx.GetFromAddress()

			sep := ","
			holderId := strings.Split(args[0], sep)
			proof := strings.Split(args[1], sep)
			data := strings.Split(args[2], sep)

			// build and sign the transaction, then broadcast to Tendermint
			msg := types.NewMsgCreateDocBatch(issuer, holderId, proof, data)

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	cmd = flags.PostCommands(cmd)[0]

	return cmd
}

func UpdateDocCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update [hodler id] [proof] [data]",
		Short: `Update the data of a document`,
		Long: `Update the data of a document.
eg: 
$ update uid-159654 c89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc6 https://sharering.network/id/463"
		`,
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)

			issuer := cliCtx.GetFromAddress()

			holderId := args[0]
			proof := args[1]
			data := args[2]

			// build and sign the transaction, then broadcast to Tendermint
			msg := types.NewMsgUpdateDoc(issuer, holderId, proof, data)

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	cmd = flags.PostCommands(cmd)[0]

	return cmd
}

func RevokeDocCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "revoke [holder id] [proof]",
		Short: `Revoke a document.`,
		Long: `Revoke a document.
eg: 
$ revoke uid-159654 c89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc6"
		`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)

			issuer := cliCtx.GetFromAddress()

			holderId := args[0]
			proof := args[1]

			// build and sign the transaction, then broadcast to Tendermint
			msg := types.NewMsgRevokeDoc(issuer, holderId, proof)

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	cmd = flags.PostCommands(cmd)[0]

	return cmd
}
