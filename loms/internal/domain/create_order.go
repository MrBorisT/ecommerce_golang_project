package domain

import (
	"context"
	"route256/loms/internal/model"
)

func (m *service) CreateOrder(ctx context.Context, user int64, items []model.Item) (int64, error) {
	return m.OrderRepository.CreateOrder(ctx, user, items)
}
