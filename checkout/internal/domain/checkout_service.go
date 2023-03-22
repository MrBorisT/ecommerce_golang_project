package domain

//go:generate sh -c "mkdir -p mocks && rm -rf mocks/cart_repository_minimock.go && rm -rf mocks/product_checker_minimock.go"
//go:generate minimock -i CartRepository -o ./mocks/ -s "_minimock.go"
//go:generate minimock -i ProductChecker -o ./mocks/ -s "_minimock.go"

import (
	"context"
	"route256/checkout/internal/model"
	"route256/checkout/internal/repository/schema"
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

type CheckoutService interface {
	AddToCart(ctx context.Context, user int64, sku uint32, count uint16) error
	DeleteFromCart(ctx context.Context, user int64, sku uint32, count uint16) error
	ListCart(ctx context.Context, user int64) ([]model.Item, uint32, error)
	Purchase(ctx context.Context, user int64) error
}

type CartRepository interface {
	AddToCart(ctx context.Context, user int64, sku uint32, count uint16) error
	DeleteFromCart(ctx context.Context, user int64, sku uint32, count uint16) error
	ListCart(ctx context.Context, user int64) ([]schema.CartItems, error)
	Purchase(ctx context.Context, user int64) error
}

type Deps struct {
	LOMS
	ProductChecker
	CartRepository
}

type service struct {
	Deps
}

func NewCheckoutService(d Deps) CheckoutService {
	return &service{d}
}
