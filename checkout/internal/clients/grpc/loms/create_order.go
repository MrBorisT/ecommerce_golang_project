package loms

import (
	"context"
	"route256/checkout/internal/model"
	lomsServiceAPI "route256/checkout/pkg/loms_v1"

	"github.com/pkg/errors"
)

func (c *client) CreateOrder(ctx context.Context, user int64, items []model.Item) (int64, error) {
	convertedItems := convertItems(items)

	res, err := c.lomsClient.CreateOrder(ctx, &lomsServiceAPI.CreateOrderRequest{
		User:  user,
		Items: convertedItems,
	})

	if err != nil {
		return 0, errors.WithMessage(err, "rpc cancel order")
	}

	orderID := res.GetOrderID()

	return orderID, nil
}

func convertItems(items []model.Item) []*lomsServiceAPI.Item {
	convertedItems := make([]*lomsServiceAPI.Item, 0, len(items))

	for _, item := range items {
		convertedItems = append(convertedItems, &lomsServiceAPI.Item{
			Sku:   item.SKU,
			Count: uint32(item.Count),
		})
	}

	return convertedItems
}
