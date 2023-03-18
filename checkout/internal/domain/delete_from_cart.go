package domain

import (
	"context"
)

func (m *CheckoutService) DeleteFromCart(ctx context.Context, user int64, sku uint32, count uint16) error {
	return m.CartRepository.DeleteFromCart(ctx, user, sku, count)
}
