package domain

import (
	"context"
)

func (m *service) OrderPayed(ctx context.Context, orderID int64) error {
	if err := m.OrderRepository.OrderPayed(ctx, orderID); err != nil {
		return err
	}
	m.StatusSender.SendStatusChange(orderID, "payed")
	return nil
}
