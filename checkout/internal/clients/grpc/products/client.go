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
	GetInt32(key T) (*int32, bool)
	GetInt64(key T) (*int64, bool)
	GetUint32(key T) (*uint32, bool)
	GetUint64(key T) (*uint64, bool)
	GetString(key T) (*string, bool)
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
