package products

import (
	"context"
	productServiceAPI "route256/checkout/pkg/product"

	"github.com/pkg/errors"
)

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
