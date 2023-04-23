package domain

import (
	"context"
	"route256/checkout/internal/model"

	"github.com/pkg/errors"
)

func (m *service) Purchase(ctx context.Context, user int64) error {
	cartItems, err := m.CartRepository.ListCart(ctx, user)
	if err != nil {
		return errors.WithMessage(err, "getting cart for purchase")
	}
	items := model.BindSchemaCartItemToItem(cartItems)

	_, err = m.LOMS.CreateOrder(ctx, user, items)
	if err != nil {
		return errors.WithMessage(err, "creating order")
	}

	return nil
}
