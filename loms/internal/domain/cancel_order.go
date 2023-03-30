package domain

import (
	"context"

	"github.com/pkg/errors"
)

func (m *service) CancelOrder(ctx context.Context, orderID int64) error {
	err := m.OrderRepository.CancelOrder(ctx, orderID)
	if err != nil {
		return err
	}

	m.StatusSender.SendStatusChange(orderID, "cancelled")

	errTX := m.TransactionManager.RunRepeatableRead(ctx, func(ctxTX context.Context) error {
		_, _, items, err := m.OrderRepository.ListOrder(ctxTX, orderID)
		if err != nil {
			return err
		}
		for _, item := range items {
			if err := m.StockRepository.UnreserveStocks(ctxTX, item.SKU, item.Count); err != nil {
				return err
			}
		}
		return nil
	})
	if errTX != nil {
		return errors.WithMessage(errTX, "unreserving stocks")
	}

	return nil
}
