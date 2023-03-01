package loms_v1

import (
	"context"
	"route256/loms/internal/model"
	desc "route256/loms/pkg/loms_v1"
)

func (i *Implementation) CreateOrder(ctx context.Context, req *desc.CreateOrderRequest) (*desc.CreateOrderResponse, error) {
	convertedItems := convertItemsToModels(req.GetItems())

	orderID, err := i.lomsService.CreateOrder(ctx, req.GetUser(), convertedItems)
	if err != nil {
		return nil, err
	}

	return &desc.CreateOrderResponse{
		OrderID: orderID,
	}, nil
}

func convertItemsToModels(items []*desc.Item) []model.Item {
	convertedItems := make([]model.Item, 0, len(items))

	for _, item := range items {
		convertedItems = append(convertedItems, model.Item{
			SKU:   item.GetSku(),
			Count: uint16(item.GetCount()),
		})
	}

	return convertedItems
}
