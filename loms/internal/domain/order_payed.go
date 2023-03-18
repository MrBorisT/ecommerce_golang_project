package domain

import (
	"context"
)

func (m *service) OrderPayed(ctx context.Context, orderID int64) error {
	return m.OrderRepository.OrderPayed(ctx, orderID)
}
