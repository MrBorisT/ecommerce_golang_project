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

type InMemCache interface {
	SetValue(key string, value any)
	GetInt32(key string) (*int32, bool)
	GetInt64(key string) (*int64, bool)
	GetUint32(key string) (*uint32, bool)
	GetUint64(key string) (*uint64, bool)
	GetString(key string) (*string, bool)
}

type client struct {
	Deps
}

type Deps struct {
	ProductClient productServiceAPI.ProductServiceClient
	Token         string
	Limiter       RateLimiter
	Cache         InMemCache
}

func NewClient(d Deps) *client {
	return &client{d}
}
