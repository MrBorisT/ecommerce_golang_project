package domain

import (
	"context"
	checkoutMocks "route256/checkout/internal/domain/mocks"
	"route256/checkout/internal/model"
	"route256/checkout/internal/repository/schema"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestListCart(t *testing.T) {
	type cartRepositoryMockFunc func(mc *minimock.Controller) CartRepository
	type productCheckerMockFunc func(mc *minimock.Controller) ProductChecker

	type args struct {
		ctx     context.Context
		reqUser uint64
	}

	type ListCartRes struct {
		items      []model.Item
		totalPrice uint32
	}

	var (
		mc  = minimock.NewController(t)
		ctx = context.Background()
		// repoErr = errors.New("repo error")
		userID = gofakeit.Int64()
		n      = 10

		repoRes     []schema.CartItems
		expectedRes = &ListCartRes{}
		sku         = 0
	)
	t.Cleanup(mc.Finish)

	for i := 0; i < n; i++ {
		repoRes = append(repoRes, schema.CartItems{
			UserID: gofakeit.Int64(),
			SKU:    gofakeit.Uint32(),
			Count:  gofakeit.Uint16(),
		})
	}

	for _, nt := range repoRes {
		expectedRes.items = append(expectedRes.items, model.Item{
			SKU:   nt.SKU,
			Count: nt.Count,
			Name:  "",
			Price: 0,
		})
	}

	tests := []struct {
		name               string
		args               args
		want               ListCartRes
		err                error
		cartRepositoryMock cartRepositoryMockFunc
		productCheckerMock productCheckerMockFunc
	}{
		{
			name: "positive case",
			args: args{
				ctx:     ctx,
				reqUser: 0,
			},
			want: *expectedRes,
			err:  nil,
			cartRepositoryMock: func(mc *minimock.Controller) CartRepository {
				mock := checkoutMocks.NewCartRepositoryMock(mc)
				mock.ListCartMock.Expect(ctx, userID).Return(repoRes, nil)
				return mock
			},
			productCheckerMock: func(mc *minimock.Controller) ProductChecker {
				mock := checkoutMocks.NewProductCheckerMock(mc)
				mock.GetProductMock.Expect(ctx, uint32(sku)).Return("", 0, nil)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			api := NewCheckoutService(Deps{
				CartRepository: tt.cartRepositoryMock(mc),
				ProductChecker: tt.productCheckerMock(mc),
			})

			resItems, resTotalPrice, err := api.ListCart(tt.args.ctx, int64(tt.args.reqUser))
			res := ListCartRes{
				items:      resItems,
				totalPrice: resTotalPrice,
			}
			require.Equal(t, tt.want, res)
			if tt.err != nil {
				require.ErrorContains(t, err, tt.err.Error())
			} else {
				require.Equal(t, tt.err, err)
			}
		})
	}
}
