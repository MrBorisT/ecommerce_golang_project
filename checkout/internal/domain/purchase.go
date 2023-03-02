package domain

import (
	"context"
	"log"
	"route256/checkout/internal/model"

	"github.com/pkg/errors"
)

func (m *Model) Purchase(ctx context.Context, user int64) error {
	DUMMY_ITEMS := make([]model.Item, 0)
	order, err := m.loms.CreateOrder(ctx, user, DUMMY_ITEMS)
	if err != nil {
		return errors.WithMessage(err, "creating order")
	}

	log.Printf("%v\n", order)

	return nil
}
