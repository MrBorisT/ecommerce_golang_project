package domain

import (
	"context"
)

func (m *CheckoutService) DeleteFromCart(ctx context.Context, user int64, sku uint32, count uint16) error {
	//TODO use cart of user with id "user"
	return nil
	//end of TODO
}
