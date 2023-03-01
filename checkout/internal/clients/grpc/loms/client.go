package loms

import (
	"context"
	"route256/checkout/internal/model"
	lomsAPI "route256/checkout/pkg/loms_v1"

	"google.golang.org/grpc"
)

type Client interface {
	CancelOrder(ctx context.Context, orderID int64) error
	CreateOrder(ctx context.Context, user int64, items []model.Item) (int64, error)
	ListOrder(ctx context.Context, orderID int64) (string, int64, []model.Item, error)
	OrderPayed(ctx context.Context, orderID int64) error
	Stocks(ctx context.Context, SKU uint32) ([]model.Stock, error)
}

type client struct {
	lomsClient lomsAPI.LomsServiceClient
}

func NewClient(cc *grpc.ClientConn) *client {
	return &client{
		lomsClient: lomsAPI.NewLomsServiceClient(cc),
	}
}
