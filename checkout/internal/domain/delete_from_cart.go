package domain

import (
	"context"
)

func (m *Model) DeleteFromCart(ctx context.Context, user int64, sku uint32, count uint16) error {
	//TODO cart logic
	return nil
}
