package domain

import (
	"context"
	"log"
	"route256/loms/internal/model"
)

func (m *service) CreateOrder(ctx context.Context, user int64, items []model.Item) (int64, error) {
	var orderID int64
	err := m.TransactionManager.RunRepeatableRead(ctx, func(ctxTX context.Context) error {
		var err error
		for _, item := range items {
			if err = m.StockRepository.ReserveStocks(ctxTX, item.SKU, item.Count); err != nil {
				return err
			}
		}

		orderID, err = m.OrderRepository.CreateOrder(ctxTX, user, items)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		log.Println("Create order failed", err)
		return 0, err
	}

	return orderID, nil
}
