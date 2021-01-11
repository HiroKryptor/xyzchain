package cli

import (
    "context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
    "github.com/hirokryptor/xyzchain/x/xyzchain/types"
)

func CmdListPool() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-pool",
		Short: "list all pool",
		RunE: func(cmd *cobra.Command, args []string) error {
            clientCtx := client.GetClientContextFromCmd(cmd)
            clientCtx, err := client.ReadQueryCommandFlags(clientCtx, cmd.Flags())
            if err != nil {
                return err
            }

            pageReq, err := client.ReadPageRequest(cmd.Flags())
            if err != nil {
                return err
            }

            queryClient := types.NewQueryClient(clientCtx)

            params := &types.QueryAllPoolRequest{
                Pagination: pageReq,
            }

            res, err := queryClient.PoolAll(context.Background(), params)
            if err != nil {
                return err
            }

            return clientCtx.PrintOutput(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

    return cmd
}

func CmdShowPool() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-pool [id]",
		Short: "shows a pool",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
            clientCtx := client.GetClientContextFromCmd(cmd)
            clientCtx, err := client.ReadQueryCommandFlags(clientCtx, cmd.Flags())
            if err != nil {
                return err
            }

            queryClient := types.NewQueryClient(clientCtx)

            params := &types.QueryGetPoolRequest{
                Id: args[0],
            }

            res, err := queryClient.Pool(context.Background(), params)
            if err != nil {
                return err
            }

            return clientCtx.PrintOutput(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

    return cmd
}
