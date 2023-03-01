package domain

import (
	"context"
	"route256/loms/internal/model"
)

type Service interface {
	CancelOrder(ctx context.Context, orderID int64) error
	CreateOrder(ctx context.Context, user int64, items []model.Item) (int64, error)
	ListOrder(ctx context.Context, orderID int64) (string, int64, []model.Item, error)
	OrderPayed(ctx context.Context, orderID int64) error
	Stocks(ctx context.Context, SKU uint32) ([]model.Stock, error)
}

type service struct {
}

func NewService() *service {
	return &service{}
}
