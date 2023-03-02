package domain

import (
	"context"

	"github.com/pkg/errors"
)

var (
	ErrInvalidSKU = errors.New("invalid SKU")
)

func (m *Model) ListCart(ctx context.Context, user int64) error {
	//TODO get cart from user
	DUMMY_CART := []struct {
		SKU   uint32
		Count uint16
		Name  string
		Price uint32
	}{
		{
			SKU:   773297411,
			Count: 1,
		},
	}

	for i, cartItem := range DUMMY_CART {
		productName, productPrice, err := m.productChecker.GetProduct(ctx, cartItem.SKU)
		if err != nil {
			return errors.WithMessage(err, "checking product")
		}
		DUMMY_CART[i].Name = productName
		DUMMY_CART[i].Price = productPrice
	}

	return nil
}
