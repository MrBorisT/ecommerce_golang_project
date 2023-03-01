package loms

import (
	"context"
	"route256/checkout/internal/domain"
	clientwrapper "route256/libs/client"
)

type StocksRequest struct {
	SKU uint32 `json:"sku"`
}

type StocksItem struct {
	WarehouseID int64  `json:"warehouseID"`
	Count       uint64 `json:"count"`
}

type StocksResponse struct {
	Stocks []StocksItem `json:"stocks"`
}

func (c *Client) Stocks(ctx context.Context, sku uint32) ([]domain.Stock, error) {
	request := StocksRequest{SKU: sku}

	response, err := clientwrapper.SendRequest[StocksRequest, StocksResponse](ctx, request, c.urlStocks)
	if err != nil {
		return nil, err
	}

	stocks := make([]domain.Stock, 0, len(response.Stocks))
	for _, stock := range response.Stocks {
		stocks = append(stocks, domain.Stock{
			WarehouseID: stock.WarehouseID,
			Count:       stock.Count,
		})
	}

	return stocks, nil
}
