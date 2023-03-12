package products

import (
	"context"
	productServiceAPI "route256/checkout/pkg/product"

	"github.com/pkg/errors"
)

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