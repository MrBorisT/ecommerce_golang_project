package domain

import (
	"context"
	"log"
	"route256/loms/internal/model"
	"time"

	"github.com/pkg/errors"
)

var (
	ErrCreatingOrder = errors.New("creating order")
	ErrReserveStocks = errors.New("reserving stocks")
)

const DEFAULT_CANCEL_ORDER_TIME = 10 * time.Minute

func (m *service) CreateOrder(ctx context.Context, user int64, items []model.Item) (int64, error) {
	orderID, err := m.OrderRepository.CreateOrder(ctx, user, items)
	if err != nil {
		return 0, errors.WithMessage(err, ErrCreatingOrder.Error())
	}
	m.StatusSender.SendStatusChange(orderID, "new")

	err = m.TransactionManager.RunRepeatableRead(ctx, func(ctxTX context.Context) error {
		var err error

		for _, item := range items {
			if err = m.StockRepository.ReserveStocks(ctxTX, item.SKU, item.Count); err != nil {
				return errors.WithMessage(err, ErrReserveStocks.Error())
			}
		}

		return nil
	})

	if err != nil {
		if errors.Is(err, ErrReserveStocks) {
			m.OrderRepository.OrderFailed(ctx, orderID)
		}
		log.Println("Create order failed", err)
		m.StatusSender.SendStatusChange(orderID, "failed")
		return 0, err
	}

	m.StatusSender.SendStatusChange(orderID, "awaiting payment")
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
