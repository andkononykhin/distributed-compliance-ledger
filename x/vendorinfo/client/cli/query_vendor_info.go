package cli

import (
	"context"

	"github.com/spf13/cast"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/vendorinfo/types"
)

func CmdListVendorInfo() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-vendors",
		Short: "Get information about all vendors",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllVendorInfoRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.VendorInfoAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowVendorInfo() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "vendor",
		Short: "Get vendor details for the given vendorID",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argVendorID, err := cast.ToInt32E(FlagVID)
			if err != nil {
				return err
			}

			params := &types.QueryGetVendorInfoRequest{
				VendorID: argVendorID,
			}

			res, err := queryClient.VendorInfo(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	cmd.Flags().String(FlagVID, "", "Unique ID assigned to the vendor")

	flags.AddQueryFlagsToCmd(cmd)

	_ = cmd.MarkFlagRequired(FlagVID)

	return cmd
}