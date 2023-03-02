package domain

import (
	"context"
	"route256/checkout/internal/model"
)

// LOMS Service
type StocksChecker interface {
	Stocks(ctx context.Context, sku uint32) ([]model.Stock, error)
}

type OrderCreator interface {
	CreateOrder(ctx context.Context, user int64, items []model.Item) (int64, error)
}

// Product Service
type ProductChecker interface {
	GetProduct(ctx context.Context, sku uint32) (string, uint32, error)
}

type LOMS interface {
	StocksChecker
	OrderCreator
}

type Model struct {
	loms           LOMS
	productChecker ProductChecker
}

func New(loms LOMS, productChecker ProductChecker) *Model {
	return &Model{
		loms:           loms,
		productChecker: productChecker,
	}
}
