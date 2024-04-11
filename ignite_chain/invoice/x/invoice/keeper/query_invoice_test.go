package keeper_test

import (
	"strconv"
	"testing"

	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "invoice/testutil/keeper"
	"invoice/testutil/nullify"
	"invoice/x/invoice/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestInvoiceQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.InvoiceKeeper(t)
	msgs := createNInvoice(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetInvoiceRequest
		response *types.QueryGetInvoiceResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetInvoiceRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetInvoiceResponse{Invoice: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetInvoiceRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetInvoiceResponse{Invoice: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetInvoiceRequest{
				Index: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Invoice(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestInvoiceQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.InvoiceKeeper(t)
	msgs := createNInvoice(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllInvoiceRequest {
		return &types.QueryAllInvoiceRequest{
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
			resp, err := keeper.InvoiceAll(ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Invoice), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Invoice),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.InvoiceAll(ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Invoice), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Invoice),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.InvoiceAll(ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Invoice),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.InvoiceAll(ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
