package cli

import (
	"errors"
	"fmt"
	"smartfreight/x/smartfreight/types"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/spf13/cobra"
)

func GetTxCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	sfTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "smartfreight transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	sfTxCmd.AddCommand(client.PostCommands(
		GetCmdSetDelivery(cdc),
		GetCmdCompleteDelivery(cdc),
	)...)

	return sfTxCmd
}

// GetCmdBuyName is the CLI command for sending a BuyName transaction
func GetCmdSetDelivery(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "set-delivery [job] [amount]",
		Short: "set new delivery",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println(args)
			fmt.Println(args[0])
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)

			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			if err := cliCtx.EnsureAccountExists(); err != nil {
				return err
			}

			jobID, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				fmt.Println("Decoding jobID")
				return errors.New("Invalid jobID")
			}

			broker, err := sdk.AccAddressFromBech32(args[1])
			if err != nil {
				fmt.Println("Decoding BBroker:")
				return err
			}

			shipper := cliCtx.GetFromAddress()
			price, err := sdk.ParseCoins(args[2])
			if err != nil {
				fmt.Println("Decoding Coins:")
				return err
			}

			d := types.NewDelivery(broker, shipper, price)

			msg := types.NewMsgSetDelivery(jobID, d)
			err = msg.ValidateBasic()
			if err != nil {
				fmt.Println("Validating message:")
				return err
			}

			cliCtx.PrintResponse = true
			fmt.Println("tx.go:")
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

// GetCmdSetName is the CLI command for sending a SetName transaction
func GetCmdCompleteDelivery(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "complete-delivery [jobID]",
		Short: "complete delivery",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)

			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			if err := cliCtx.EnsureAccountExists(); err != nil {
				return err
			}

			broker := cliCtx.GetFromAddress()
			jobID, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				return errors.New("Invalid jobID")
			}

			msg := types.NewMsgCompleteDelivery(jobID, broker)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			cliCtx.PrintResponse = true

			// return utils.CompleteAndBroadcastTxCLI(txBldr, cliCtx, msgs)
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
