package loms

import (
	"context"
	"route256/checkout/internal/domain"
	clientwrapper "route256/libs/client"
)

type Client struct {
	url            string
	urlStocks      string
	urlCreateOrder string
}

func New(url string) *Client {
	return &Client{
		url:            url,
		urlStocks:      url + "/stocks",
		urlCreateOrder: url + "/createOrder",
	}
}

// STOCKS
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

// CREATE ORDER
type CreateOrderRequest struct {
	User int64 `json:"user"`
}

type CreateOrderResponse struct {
	OrderID int64 `json:"orderID"`
}

func (c *Client) CreateOrder(ctx context.Context, user int64) (*domain.Order, error) {
	request := CreateOrderRequest{User: user}
	response, err := clientwrapper.SendRequest[CreateOrderRequest, CreateOrderResponse](ctx, request, c.urlCreateOrder)
	if err != nil {
		return nil, err
	}

	order := domain.Order{
		OrderID: response.OrderID,
	}

	return &order, nil
}
