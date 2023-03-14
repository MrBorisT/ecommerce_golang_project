package domain

import (
	"context"
)

func (m *service) CancelOrder(ctx context.Context, orderID int64) error {
	_, _, items, err := m.OrderRepository.ListOrder(ctx, orderID)
	if err != nil {
		return err
	}
	for _, item := range items {
		if err := m.StockRepository.UnreserveStocks(ctx, item.SKU, item.Count); err != nil {
			return err
		}
	}

	return m.OrderRepository.CancelOrder(ctx, orderID)
}
