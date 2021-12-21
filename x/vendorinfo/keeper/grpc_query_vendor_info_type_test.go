package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/zigbee-alliance/distributed-compliance-ledger/testutil/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/vendorinfo/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestVendorInfoTypeQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.VendorinfoKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNVendorInfoType(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetVendorInfoTypeRequest
		response *types.QueryGetVendorInfoTypeResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetVendorInfoTypeRequest{
				VendorID: msgs[0].VendorID,
			},
			response: &types.QueryGetVendorInfoTypeResponse{VendorInfoType: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetVendorInfoTypeRequest{
				VendorID: msgs[1].VendorID,
			},
			response: &types.QueryGetVendorInfoTypeResponse{VendorInfoType: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetVendorInfoTypeRequest{
				VendorID: 100000,
			},
			err: status.Error(codes.InvalidArgument, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.VendorInfoType(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.Equal(t, tc.response, response)
			}
		})
	}
}

func TestVendorInfoTypeQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.VendorinfoKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNVendorInfoType(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllVendorInfoTypeRequest {
		return &types.QueryAllVendorInfoTypeRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.VendorInfoTypeAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.VendorInfoType), step)
			require.Subset(t, msgs, resp.VendorInfoType)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.VendorInfoTypeAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.VendorInfoType), step)
			require.Subset(t, msgs, resp.VendorInfoType)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.VendorInfoTypeAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.VendorInfoTypeAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
