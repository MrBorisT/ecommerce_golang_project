package domain

import "context"

type StocksChecker interface {
	Stocks(ctx context.Context, sku uint32) ([]Stock, error)
}

type ProductChecker interface {
	Product(ctx context.Context, sku uint32) (*Product, error)
}

type Model struct {
	stocksChecker  StocksChecker
	productChecker ProductChecker
}

func New(stocksChecker StocksChecker, productChecker ProductChecker) *Model {
	return &Model{
		stocksChecker:  stocksChecker,
		productChecker: productChecker,
	}
}
