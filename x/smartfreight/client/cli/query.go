package cli

import (
	"fmt"

	"smartfreight/x/smartfreight/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
)

func GetQueryCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	sfQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the SmartFreight module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	sfQueryCmd.AddCommand(client.GetCommands(
		GetCmdResolve(storeKey, cdc),
		GetCmdNames(storeKey, cdc),
	)...)
	return sfQueryCmd
}

// GetCmdResolve queries information about a name
func GetCmdResolve(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "resolve [jobId]",
		Short: "resolve jobId",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			jobID := args[0]

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/resolve/%s", queryRoute, jobID), nil)
			if err != nil {
				fmt.Printf("could not resolve jobID - %s \n", jobID)
				return nil
			}
			var out types.Delivery
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

// GetCmdNames queries a list of all names
func GetCmdNames(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "keys",
		Short: "keys",
		// Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/keys", queryRoute), nil)
			if err != nil {
				fmt.Printf("could not get query keys\n")
				return nil
			}

			var out types.QueryResKeys
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}
