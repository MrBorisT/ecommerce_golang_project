package loms

import (
	"context"
	"route256/checkout/internal/model"
	lomsServiceAPI "route256/checkout/pkg/loms_v1"

	"github.com/pkg/errors"
)

func (c *client) ListOrder(ctx context.Context, orderID int64) (string, int64, []model.Item, error) {
	res, err := c.lomsClient.ListOrder(ctx, &lomsServiceAPI.ListOrderRequest{
		OrderID: orderID,
	})

	if err != nil {
		return "", 0, nil, errors.WithMessage(err, "rpc cancel order")
	}

	convertedItems := convertItemsToModels(res.GetItems())

	return res.GetStatus(), res.GetUser(), convertedItems, nil
}

func convertItemsToModels(items []*lomsServiceAPI.Item) []model.Item {
	convertedItems := make([]model.Item, 0, len(items))

	for _, item := range items {
		convertedItems = append(convertedItems, model.Item{
			SKU:   item.GetSku(),
			Count: uint16(item.GetCount()),
		})
	}

	return convertedItems
}
