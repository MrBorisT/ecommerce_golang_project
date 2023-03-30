package domain

import (
	"context"
	checkoutMocks "route256/checkout/internal/domain/mocks"
	"route256/checkout/internal/model"
	"route256/checkout/internal/repository/schema"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func TestListCart(t *testing.T) {
	type cartRepositoryMockFunc func(mc *minimock.Controller) CartRepository
	type productCheckerMockFunc func(mc *minimock.Controller) ProductChecker

	type args struct {
		ctx     context.Context
		reqUser int64
	}

	type ListCartRes struct {
		items      []model.Item
		totalPrice uint32
	}

	var (
		mc  = minimock.NewController(t)
		ctx = context.Background()
		n   = 100

		userID       = gofakeit.Int64()
		sku          = gofakeit.Uint32()
		count        = gofakeit.Uint16()
		productName  = gofakeit.BeerName()
		productPrice = gofakeit.Uint32()

		repoRes          []schema.CartItems
		expectedRes      = ListCartRes{}
		expectedEmptyRes = ListCartRes{}

		repoErr        = errors.New("repo error")
		prodServiceErr = errors.New("product checker error")
	)
	t.Cleanup(mc.Finish)

	for i := 0; i < n; i++ {
		repoRes = append(repoRes, schema.CartItems{
			UserID: userID,
			SKU:    sku,
			Count:  count,
		})
	}

	for _, nt := range repoRes {
		expectedRes.items = append(expectedRes.items, model.Item{
			SKU:   nt.SKU,
			Count: nt.Count,
			Name:  productName,
			Price: productPrice,
		})

		expectedRes.totalPrice += uint32(nt.Count) * productPrice
	}

	tests := []struct {
		name               string
		args               args
		want               *ListCartRes
		err                error
		cartRepositoryMock cartRepositoryMockFunc
		productCheckerMock productCheckerMockFunc
	}{
		{
			name: "positive case",
			args: args{
				ctx:     ctx,
				reqUser: userID,
			},
			want: &expectedRes,
			err:  nil,
			cartRepositoryMock: func(mc *minimock.Controller) CartRepository {
				mock := checkoutMocks.NewCartRepositoryMock(mc)
				mock.ListCartMock.Expect(ctx, userID).Return(repoRes, nil)
				return mock
			},
			productCheckerMock: func(mc *minimock.Controller) ProductChecker {
				mock := checkoutMocks.NewProductCheckerMock(mc)
				mock.GetProductMock.Expect(ctx, sku).Return(productName, productPrice, nil)
				return mock
			},
		},
		{
			name: "repo error",
			args: args{
				ctx:     ctx,
				reqUser: userID,
			},
			want: &expectedEmptyRes,
			err:  repoErr,
			cartRepositoryMock: func(mc *minimock.Controller) CartRepository {
				mock := checkoutMocks.NewCartRepositoryMock(mc)
				mock.ListCartMock.Expect(ctx, userID).Return(nil, repoErr)
				return mock
			},
			productCheckerMock: func(mc *minimock.Controller) ProductChecker {
				mock := checkoutMocks.NewProductCheckerMock(mc)
				return mock
			},
		},
		{
			name: "product checker error",
			args: args{
				ctx:     ctx,
				reqUser: userID,
			},
			want: &expectedEmptyRes,
			err:  prodServiceErr,
			cartRepositoryMock: func(mc *minimock.Controller) CartRepository {
				mock := checkoutMocks.NewCartRepositoryMock(mc)
				mock.ListCartMock.Expect(ctx, userID).Return(repoRes, nil)
				return mock
			},
			productCheckerMock: func(mc *minimock.Controller) ProductChecker {
				mock := checkoutMocks.NewProductCheckerMock(mc)
				mock.GetProductMock.Expect(ctx, sku).Return(productName, productPrice, prodServiceErr)
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

			resItems, resTotalPrice, err := api.ListCart(tt.args.ctx, tt.args.reqUser)
			res := &ListCartRes{
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
