package domain

import (
	"context"
	"log"
)

func (m *service) CancelOrder(ctx context.Context, orderID int64) error {
	err := m.TransactionManager.RunRepeatableRead(ctx, func(ctxTX context.Context) error {
		_, _, items, err := m.OrderRepository.ListOrder(ctxTX, orderID)
		if err != nil {
			return err
		}
		for _, item := range items {
			if err := m.StockRepository.UnreserveStocks(ctxTX, item.SKU, item.Count); err != nil {
				return err
			}
		}
		return m.OrderRepository.CancelOrder(ctxTX, orderID)
	})

	if err != nil {
		log.Println("Order cancel failed", err)
		return err
	}

	return nil
}
