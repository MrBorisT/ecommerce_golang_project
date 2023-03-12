package domain

import (
	"context"
	"log"
	"route256/checkout/internal/model"

	"github.com/pkg/errors"
)

func (m *CheckoutService) Purchase(ctx context.Context, user int64) error {
	//TODO use cart of user with id "user"
	DUMMY_ITEMS := make([]model.Item, 1)
	DUMMY_ITEMS[0] = model.Item{
		SKU:   773297411,
		Count: 1,
	}
	//end of TODO

	order, err := m.Deps.CreateOrder(ctx, user, DUMMY_ITEMS)
	if err != nil {
		return errors.WithMessage(err, "creating order")
	}

	log.Printf("Created order: %v\n", order)

	return nil
}
