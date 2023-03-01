package loms

import (
	"context"
	"route256/checkout/internal/model"
	clientwrapper "route256/libs/client"
)

type StocksRequest struct {
	SKU uint32 `json:"sku"`
}

type StocksResponse struct {
	Stocks []model.Stock `json:"stocks"`
}

func (c *Client) Stocks(ctx context.Context, sku uint32) ([]model.Stock, error) {
	request := StocksRequest{SKU: sku}

	response, err := clientwrapper.SendRequest[StocksRequest, StocksResponse](ctx, request, c.urlStocks)
	if err != nil {
		return nil, err
	}

	stocks := make([]model.Stock, 0, len(response.Stocks))
	for _, stock := range response.Stocks {
		stocks = append(stocks, model.Stock{
			WarehouseID: stock.WarehouseID,
			Count:       stock.Count,
		})
	}

	return stocks, nil
}
