package loms

import (
	"context"
	"route256/checkout/internal/domain"
	clientwrapper "route256/libs/client"
)

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
