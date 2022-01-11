package cli_test

/* TODO issue 99
import (
	"fmt"
	"strconv"
	"testing"

	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/stretchr/testify/require"
	tmcli "github.com/tendermint/tendermint/libs/cli"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/network"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/nullify"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliancetest/client/cli"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliancetest/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Prevent strconv unused error.
var _ = strconv.IntSize

func networkWithTestingResultsObjects(t *testing.T, n int) (*network.Network, []types.TestingResults) {
	t.Helper()
	cfg := network.DefaultConfig()
	state := types.GenesisState{}
	require.NoError(t, cfg.Codec.UnmarshalJSON(cfg.GenesisState[types.ModuleName], &state))

	for i := 0; i < n; i++ {
		testingResults := types.TestingResults{
			Vid:             int32(i),
			Pid:             int32(i),
			SoftwareVersion: uint32(i),
		}
		nullify.Fill(&testingResults)
		state.TestingResultsList = append(state.TestingResultsList, testingResults)
	}
	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[types.ModuleName] = buf
	return network.New(t, cfg), state.TestingResultsList
}

func TestShowTestingResults(t *testing.T) {
	net, objs := networkWithTestingResultsObjects(t, 2)

	ctx := net.Validators[0].ClientCtx
	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	for _, tc := range []struct {
		desc              string
		idVid             int32
		idPid             int32
		idSoftwareVersion uint32

		args []string
		err  error
		obj  types.TestingResults
	}{
		{
			desc:              "found",
			idVid:             objs[0].Vid,
			idPid:             objs[0].Pid,
			idSoftwareVersion: objs[0].SoftwareVersion,

			args: common,
			obj:  objs[0],
		},
		{
			desc:              "not found",
			idVid:             100000,
			idPid:             100000,
			idSoftwareVersion: 100000,

			args: common,
			err:  status.Error(codes.NotFound, "not found"),
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			args := []string{
				strconv.Itoa(int(tc.idVid)),
				strconv.Itoa(int(tc.idPid)),
				strconv.Itoa(int(tc.idSoftwareVersion)),
			}
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowTestingResults(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryGetTestingResultsResponse
				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.TestingResults)
				require.Equal(t,
					nullify.Fill(&tc.obj),
					nullify.Fill(&resp.TestingResults),
				)
			}
		})
	}
}

func TestListTestingResults(t *testing.T) {
	net, objs := networkWithTestingResultsObjects(t, 5)

	ctx := net.Validators[0].ClientCtx
	request := func(next []byte, offset, limit uint64, total bool) []string {
		args := []string{
			fmt.Sprintf("--%s=json", tmcli.OutputFlag),
		}
		if next == nil {
			args = append(args, fmt.Sprintf("--%s=%d", flags.FlagOffset, offset))
		} else {
			args = append(args, fmt.Sprintf("--%s=%s", flags.FlagPageKey, next))
		}
		args = append(args, fmt.Sprintf("--%s=%d", flags.FlagLimit, limit))
		if total {
			args = append(args, fmt.Sprintf("--%s", flags.FlagCountTotal))
		}
		return args
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(objs); i += step {
			args := request(nil, uint64(i), uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListTestingResults(), args)
			require.NoError(t, err)
			var resp types.QueryAllTestingResultsResponse
			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.TestingResults), step)
			require.Subset(t,
				nullify.Fill(objs),
				nullify.Fill(resp.TestingResults),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(objs); i += step {
			args := request(next, 0, uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListTestingResults(), args)
			require.NoError(t, err)
			var resp types.QueryAllTestingResultsResponse
			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.TestingResults), step)
			require.Subset(t,
				nullify.Fill(objs),
				nullify.Fill(resp.TestingResults),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		args := request(nil, 0, uint64(len(objs)), true)
		out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListTestingResults(), args)
		require.NoError(t, err)
		var resp types.QueryAllTestingResultsResponse
		require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
		require.NoError(t, err)
		require.Equal(t, len(objs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(objs),
			nullify.Fill(resp.TestingResults),
		)
	})
}
*/
