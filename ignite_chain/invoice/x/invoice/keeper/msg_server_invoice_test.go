package keeper_test

import (
	"strconv"
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "invoice/testutil/keeper"
	"invoice/x/invoice/keeper"
	"invoice/x/invoice/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestInvoiceMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.InvoiceKeeper(t)
	srv := keeper.NewMsgServerImpl(k)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateInvoice{Creator: creator,
			Index: strconv.Itoa(i),
		}
		_, err := srv.CreateInvoice(ctx, expected)
		require.NoError(t, err)
		rst, found := k.GetInvoice(ctx,
			expected.Index,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestInvoiceMsgServerUpdate(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgUpdateInvoice
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateInvoice{Creator: creator,
				Index: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateInvoice{Creator: "B",
				Index: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateInvoice{Creator: creator,
				Index: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.InvoiceKeeper(t)
			srv := keeper.NewMsgServerImpl(k)
			expected := &types.MsgCreateInvoice{Creator: creator,
				Index: strconv.Itoa(0),
			}
			_, err := srv.CreateInvoice(ctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateInvoice(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetInvoice(ctx,
					expected.Index,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestInvoiceMsgServerDelete(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgDeleteInvoice
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteInvoice{Creator: creator,
				Index: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteInvoice{Creator: "B",
				Index: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteInvoice{Creator: creator,
				Index: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.InvoiceKeeper(t)
			srv := keeper.NewMsgServerImpl(k)

			_, err := srv.CreateInvoice(ctx, &types.MsgCreateInvoice{Creator: creator,
				Index: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteInvoice(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetInvoice(ctx,
					tc.request.Index,
				)
				require.False(t, found)
			}
		})
	}
}
