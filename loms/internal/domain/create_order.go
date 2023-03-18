package domain

import (
	"context"
	"log"
	"route256/loms/internal/model"
	"time"
)

const DEFAULT_CANCEL_ORDER_TIME = 10 * time.Minute

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

	go m.CancelOrderByTimer(ctx, orderID)

	return orderID, nil
}

func (m *service) CancelOrderByTimer(ctx context.Context, orderID int64) {
	timer := time.NewTimer(DEFAULT_CANCEL_ORDER_TIME)
	<-timer.C
	var status string
	var err error
	if status, _, _, err = m.OrderRepository.ListOrder(ctx, orderID); err != nil {
		log.Printf("getting order when cancelling: %v\n", err)
	}
	if status != "new" {
		return
	}
	if err = m.OrderRepository.CancelOrder(ctx, orderID); err != nil {
		log.Printf("cancel order by timer: %v\n", err)
	}

}
