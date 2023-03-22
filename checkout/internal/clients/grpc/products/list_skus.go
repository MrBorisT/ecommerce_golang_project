package products

import (
	"context"
	productServiceAPI "route256/checkout/pkg/product"

	"github.com/pkg/errors"
)

func (c *client) ListSkus(ctx context.Context, start_after_sku, count uint32) ([]uint32, error) {
	if err := c.Limiter.Wait(ctx); err != nil {
		return nil, err
	}
	res, err := c.ProductClient.ListSkus(ctx, &productServiceAPI.ListSkusRequest{
		Token:         c.Token,
		StartAfterSku: start_after_sku,
		Count:         count,
	})
	if err != nil {
		return nil, errors.WithMessage(err, "rpc list skus")
	}

	return res.GetSkus(), nil
}
