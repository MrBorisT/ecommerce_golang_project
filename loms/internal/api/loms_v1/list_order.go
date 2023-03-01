package loms_v1

import (
	"context"
	"route256/loms/internal/model"
	desc "route256/loms/pkg/loms_v1"
)

func (i *Implementation) ListOrder(ctx context.Context, req *desc.ListOrderRequest) (*desc.ListOrderResponse, error) {
	status, user, items, err := i.lomsService.ListOrder(ctx, req.GetOrderID())
	if err != nil {
		return nil, err
	}

	convertedItems := convertItemsToDesc(items)

	return &desc.ListOrderResponse{
		Status: status,
		User:   user,
		Items:  convertedItems,
	}, nil
}

func convertItemsToDesc(items []model.Item) []*desc.Item {
	convertedItems := make([]*desc.Item, 0, len(items))

	for _, item := range items {
		convertedItems = append(convertedItems, &desc.Item{
			Sku:   item.SKU,
			Count: uint32(item.Count),
		})
	}

	return convertedItems
}
