package products

import (
	"context"
	productServiceAPI "route256/checkout/pkg/product"

	"github.com/pkg/errors"
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

func (c *client) GetProduct(ctx context.Context, sku uint32) (string, uint32, error) {
	res, err := c.productClient.GetProduct(ctx, &productServiceAPI.GetProductRequest{
		Token: c.token,
		Sku:   sku,
	})
	if err != nil {
		return "", 0, errors.WithMessage(err, "rpc get product")
	}

	return res.GetName(), res.GetPrice(), nil
}

func (c *client) ListSkus(ctx context.Context, start_after_sku, count uint32) ([]uint32, error) {
	res, err := c.productClient.ListSkus(ctx, &productServiceAPI.ListSkusRequest{
		Token:         c.token,
		StartAfterSku: start_after_sku,
		Count:         count,
	})
	if err != nil {
		return nil, errors.WithMessage(err, "rpc list skus")
	}

	return res.GetSkus(), nil
}
