package domain

import (
	"context"
	"log"
	"route256/checkout/internal/model"

	"github.com/pkg/errors"
)

func (m *CheckoutService) Purchase(ctx context.Context, user int64) error {
	cart, err := m.CartRepository.ListCart(ctx, user)
	if err != nil {
		return errors.WithMessage(err, "getting cart for purchase")
	}
	items := model.BindCartItemToItem(cart.Items)

	order, err := m.Deps.CreateOrder(ctx, user, items)
	if err != nil {
		return errors.WithMessage(err, "creating order")
	}

	log.Printf("Created order: %v\n", order)

	return nil
}
