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

type OrderRepository interface {
	CancelOrder(ctx context.Context, orderID int64) error
	CreateOrder(ctx context.Context, user int64, items []model.Item) (int64, error)
	ListOrder(ctx context.Context, orderID int64) (string, int64, []model.Item, error)
	OrderPayed(ctx context.Context, orderID int64) error
}

type StockRepository interface {
	Stocks(ctx context.Context, SKU uint32) ([]model.Stock, error)
	ReserveStocks(ctx context.Context, SKU uint32, count uint16) error
	UnreserveStocks(ctx context.Context, SKU uint32, count uint16) error
}

type TransactionManager interface {
	RunRepeatableRead(ctx context.Context, f func(ctxTX context.Context) error) error
}

type StatusSender interface {
	SendStatusChange(orderID int64, status string)
}

type Deps struct {
	OrderRepository
	StockRepository
	TransactionManager
	StatusSender
}

type service struct {
	Deps
}

func NewService(d Deps) *service {
	return &service{d}
}
