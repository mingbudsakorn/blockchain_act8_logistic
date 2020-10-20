// Step 12: add tx cmd which are initDeal, transport, updateTemp, receive, reject
package cli

import (
	"bufio"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/earth2378/logistic/x/logistic/types"
)

// GetCmdInitDeal receive orderid, price, customer, maxTemp and minTemp
// then create a new deal
func GetCmdInitDeal(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "init-deal [orderid] [price] [customer] [maxTemp] [minTemp]",
		Short: "Init a new deal",
		Args:  cobra.ExactArgs(5), // Does your request require arguments
		RunE: func(cmd *cobra.Command, args []string) error {

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			orderid := args[0]

			price, err := sdk.ParseCoins(args[1])
			if err != nil {
				return err
			}

			customer, err := sdk.AccAddressFromBech32(args[2])
			if err != nil {
				return err
			}

			maxTemp, err := strconv.Atoi(args[3])
			if err != nil {
				return err
			}

			minTemp, err := strconv.Atoi(args[4])
			if err != nil {
				return err
			}

			msg := types.NewMsgInitDeal(cliCtx.GetFromAddress(), orderid, price, customer, maxTemp, minTemp)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

// GetCmdTransport receive orderid, transporter
// then assign transporter to deal with orderid
func GetCmdTransport(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "transport [orderid] [transporter]",
		Short: "select transporter",
		Args:  cobra.ExactArgs(2), // Does your request require arguments
		RunE: func(cmd *cobra.Command, args []string) error {

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			orderid := args[0]

			transporter, err := sdk.AccAddressFromBech32(args[1])
			if err != nil {
				return err
			}

			msg := types.NewMsgTransport(cliCtx.GetFromAddress(), transporter, orderid)

			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

// GetCmdTransport receive orderid, temp
// check if current temp is still acceptable or not
func GetCmdUpdateTemp(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "update-temp [orderid] [temp]",
		Short: "update current temp of the provided orderid",
		Args:  cobra.ExactArgs(2), // Does your request require arguments
		RunE: func(cmd *cobra.Command, args []string) error {

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			orderid := args[0]
			temp, err := strconv.Atoi(args[1])
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateTemp(cliCtx.GetFromAddress(), orderid, temp)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

// GetCmdReceive receive orderid, temp
// then customer accept product and make a payment to owner
func GetCmdReceive(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "receive [orderid]",
		Short: "receive product",
		Args:  cobra.ExactArgs(1), // Does your request require arguments
		RunE: func(cmd *cobra.Command, args []string) error {

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			orderid := args[0]

			msg := types.NewMsgReceive(cliCtx.GetFromAddress(), orderid)
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
