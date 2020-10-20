package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/earth2378/logistic/x/logistic/types"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd(cdc *codec.Codec) *cobra.Command {
	logisticTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	logisticTxCmd.AddCommand(flags.PostCommands(
		// this line is used by starport scaffolding
		// Step 11: add cli command for tx
		GetCmdInitDeal(cdc),
		GetCmdTransport(cdc),
		GetCmdUpdateTemp(cdc),
		GetCmdReceive(cdc),
	)...)

	return logisticTxCmd
}
