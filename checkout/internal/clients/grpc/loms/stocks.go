package loms

import (
	"context"
	"route256/checkout/internal/model"
	lomsServiceAPI "route256/checkout/pkg/loms_v1"

	"github.com/pkg/errors"
)

func (c *client) Stocks(ctx context.Context, SKU uint32) ([]model.Stock, error) {
	res, err := c.lomsClient.Stocks(ctx, &lomsServiceAPI.StocksRequest{
		Sku: SKU,
	})

	if err != nil {
		return nil, errors.WithMessage(err, "rpc stocks")
	}

	convertedStocks := convertStocks(res.GetStocks())
	return convertedStocks, nil
}

func convertStocks(stocks []*lomsServiceAPI.Stock) []model.Stock {
	convertedStocks := make([]model.Stock, 0, len(stocks))

	for _, stock := range stocks {
		convertedStocks = append(convertedStocks, model.Stock{
			WarehouseID: stock.GetWarehouseID(),
			Count:       stock.GetCount(),
		})
	}

	return convertedStocks
}
