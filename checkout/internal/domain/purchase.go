package domain

import (
	"context"
	"log"

	"github.com/pkg/errors"
)

type Order struct {
	OrderID int64
}

type OrderCreator interface {
	CreateOrder(ctx context.Context, user int64) (*Order, error)
}

func (m *Model) Purchase(ctx context.Context, user int64) error {
	order, err := m.loms.CreateOrder(ctx, user)
	if err != nil {
		return errors.WithMessage(err, "creating order")
	}

	log.Printf("%v\n", order)

	return nil
}
