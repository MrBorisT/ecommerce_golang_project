package domain

import (
	"context"
	"route256/loms/internal/model"
)

func (m *service) CreateOrder(ctx context.Context, user int64, items []model.Item) (int64, error) {
	for _, item := range items {
		if err := m.StockRepository.ReserveStocks(ctx, item.SKU, item.Count); err != nil {
			return 0, err
		}
	}

	orderID, err := m.OrderRepository.CreateOrder(ctx, user, items)
	if err != nil {
		return 0, err
	}

	return orderID, nil
}
