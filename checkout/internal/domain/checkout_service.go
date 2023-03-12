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

type CartRepository interface {
	AddToCart(ctx context.Context, user int64, sku uint32, count uint16) error
	DeleteFromCart(ctx context.Context, user int64, sku uint32, count uint16) error
	ListCart(ctx context.Context, user int64) (*model.Cart, error)
	Purchase(ctx context.Context, user int64) error
}

type Deps struct {
	LOMS
	ProductChecker
	CartRepository
}

type CheckoutService struct {
	Deps
}

func NewCheckoutService(d Deps) *CheckoutService {
	return &CheckoutService{d}
}
