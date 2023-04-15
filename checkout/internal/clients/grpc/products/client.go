package products

import (
	"context"
	productServiceAPI "route256/checkout/pkg/product"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
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
	GetValue(key T) (*any, bool)
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

var cacheHitCount = promauto.NewCounter(prometheus.CounterOpts{
	Name: "app_cache_hits_total",
})

var cacheRequestsTotal = promauto.NewCounter(prometheus.CounterOpts{
	Name: "app_cache_requests_total",
})

var cacheErrorsTotal = promauto.NewCounter(prometheus.CounterOpts{
	Name: "app_cache_errors_total",
})
