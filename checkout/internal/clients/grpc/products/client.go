package products

import (
	"context"
	productServiceAPI "route256/checkout/pkg/product"

	"google.golang.org/grpc"
)

type Client interface {
	GetProduct(ctx context.Context, token string, sku uint32) (string, uint32, error)
	ListSkus(ctx context.Context, token string, start_after_sku, count uint32) ([]uint32, error)
}

type client struct {
	productClient productServiceAPI.ProductServiceClient
	token         string
}

func NewClient(cc *grpc.ClientConn, token string) *client {
	return &client{
		productClient: productServiceAPI.NewProductServiceClient(cc),
		token:         token,
	}
}
