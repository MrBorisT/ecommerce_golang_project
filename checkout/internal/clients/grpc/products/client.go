package products

import (
	"context"
	productServiceAPI "route256/checkout/pkg/product"
)

type Client interface {
	GetProduct(ctx context.Context, token string, sku uint32) (string, uint32, error)
	ListSkus(ctx context.Context, token string, start_after_sku, count uint32) ([]uint32, error)
}

type RateLimiter interface {
	Wait(ctx context.Context) error
}

type InMemCache[T comparable] interface {
	SetValue(key T, value any)
	GetRawValue(key T) (*any, bool)
	ClearValue(key T)
}

type client struct {
	Deps
}

type Deps struct {
	ProductClient productServiceAPI.ProductServiceClient
	Token         string
	Limiter       RateLimiter
	Cache         InMemCache[uint32]
}

func NewClient(d Deps) *client {
	return &client{d}
}
