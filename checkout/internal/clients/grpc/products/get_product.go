package products

import (
	"context"
	productServiceAPI "route256/checkout/pkg/product"
	"route256/libs/logger"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type product struct {
	name  string
	price uint32
}

func (c *client) GetProduct(ctx context.Context, sku uint32) (string, uint32, error) {
	productInfo, ok := c.Cache.GetRawValue(sku)
	if !ok {
		return c.requestProductAndCache(ctx, sku)
	}

	productInfoConverted, ok := (*productInfo).(product)
	if !ok {
		logger.Error("invalid cache", zap.Uint32("sku", sku))
		c.Cache.ClearValue(sku)
		return c.requestProductAndCache(ctx, sku)
	}

	return productInfoConverted.name, productInfoConverted.price, nil
}

func (c *client) requestProductAndCache(ctx context.Context, sku uint32) (string, uint32, error) {
	if err := c.Limiter.Wait(ctx); err != nil {
		return "", 0, err
	}
	res, err := c.ProductClient.GetProduct(ctx, &productServiceAPI.GetProductRequest{
		Token: c.Token,
		Sku:   sku,
	})
	if err != nil {
		return "", 0, errors.WithMessage(err, "rpc get product")
	}

	productToCache := product{
		name:  res.GetName(),
		price: res.GetPrice(),
	}
	c.Cache.SetValue(sku, productToCache)

	return res.GetName(), res.GetPrice(), nil
}
