package domain

import "context"

type StocksChecker interface {
	Stocks(ctx context.Context, sku uint32) ([]Stock, error)
}

type ProductChecker interface {
	Product(ctx context.Context, sku uint32) (*Product, error)
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
