package domain

import (
	"context"
	"log"

	"github.com/pkg/errors"
)

type Product struct {
	Name  string
	Price uint32
}

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
		productInfo, err := m.productChecker.Product(ctx, cartItem.SKU)
		if err != nil {
			return errors.WithMessage(err, "checking product")
		}
		DUMMY_CART[i].Name = productInfo.Name
		DUMMY_CART[i].Price = productInfo.Price
		log.Println("got product from grcp:", productInfo)
	}

	return nil
}
