package cli

import (
  "strconv"
	"github.com/spf13/cobra"

    "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/hirokryptor/xyzchain/x/xyzchain/types"
)

func CmdCreatePool() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-pool [token1] [token2] [total1] [total2]",
		Short: "Creates a new pool",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
      argsToken1 := string(args[0])
      argsToken2 := string(args[1])
      argsTotal1, _ := strconv.ParseInt(args[2], 10, 64)
      argsTotal2, _ := strconv.ParseInt(args[3], 10, 64)
      
        	clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			msg := types.NewMsgCreatePool(clientCtx.GetFromAddress().String(), string(argsToken1), string(argsToken2), int32(argsTotal1), int32(argsTotal2))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

    return cmd
}

func CmdUpdatePool() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-pool [id] [token1] [token2] [total1] [total2]",
		Short: "Update a pool",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
            id := args[0]
      argsToken1 := string(args[1])
      argsToken2 := string(args[2])
      argsTotal1, _ := strconv.ParseInt(args[3], 10, 64)
      argsTotal2, _ := strconv.ParseInt(args[4], 10, 64)
      
        	clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdatePool(clientCtx.GetFromAddress().String(), id, string(argsToken1), string(argsToken2), int32(argsTotal1), int32(argsTotal2))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

    return cmd
}

func CmdDeletePool() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-pool [id] [token1] [token2] [total1] [total2]",
		Short: "Delete a pool by id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
            id := args[0]

        	clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			msg := types.NewMsgDeletePool(clientCtx.GetFromAddress().String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

    return cmd
}
