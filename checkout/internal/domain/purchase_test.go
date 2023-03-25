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

func TestPurchase(t *testing.T) {
	type cartRepositoryMockFunc func(mc *minimock.Controller) CartRepository
	type lomsMockFunc func(mc *minimock.Controller) LOMS

	type args struct {
		ctx     context.Context
		reqUser int64
	}

	var (
		mc  = minimock.NewController(t)
		ctx = context.Background()
		n   = 1

		userID = gofakeit.Int64()
		sku    = gofakeit.Uint32()

		repoRes []schema.CartItems
		lomsReq []model.Item

		// repoErr = errors.New("repo error")
	)
	t.Cleanup(mc.Finish)

	for i := 0; i < n; i++ {
		repoRes = append(repoRes, schema.CartItems{
			UserID: userID,
			SKU:    sku,
			Count:  gofakeit.Uint16(),
		})
	}

	lomsReq = model.BindSchemaCartItemToItem(repoRes)

	tests := []struct {
		name               string
		args               args
		err                error
		cartRepositoryMock cartRepositoryMockFunc
		lomsMock           lomsMockFunc
	}{
		{
			name: "positive case",
			args: args{
				ctx:     ctx,
				reqUser: userID,
			},
			err: nil,
			cartRepositoryMock: func(mc *minimock.Controller) CartRepository {
				mock := checkoutMocks.NewCartRepositoryMock(mc)
				mock.ListCartMock.Expect(ctx, userID).Return(repoRes, nil)
				return mock
			},
			lomsMock: func(mc *minimock.Controller) LOMS {
				mock := checkoutMocks.NewLOMSMock(mc)
				//returned order doesn't matter so loms returns 0
				mock.CreateOrderMock.Expect(ctx, userID, lomsReq).Return(0, nil)
				return mock
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// t.Parallel()
			api := NewCheckoutService(Deps{
				CartRepository: tt.cartRepositoryMock(mc),
				LOMS:           tt.lomsMock(mc),
			})

			err := api.Purchase(tt.args.ctx, tt.args.reqUser)

			if tt.err != nil {
				require.ErrorContains(t, err, tt.err.Error())
			} else {
				require.Equal(t, tt.err, err)
			}
		})
	}
}
