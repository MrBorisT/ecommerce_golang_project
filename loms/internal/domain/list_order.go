package domain

import (
	"context"
	"route256/loms/internal/model"
)

func (m *service) ListOrder(ctx context.Context, orderID int64) (string, int64, []model.Item, error) {
	return m.OrderRepository.ListOrder(ctx, orderID)
}
