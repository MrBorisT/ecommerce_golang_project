package domain

import (
	"context"
	"route256/checkout/internal/model"

	"github.com/pkg/errors"
)

func (m *CheckoutService) ListCart(ctx context.Context, user int64) ([]model.Item, uint32, error) {
	//TODO use cart of user with id "user"
	DUMMY_CART := []model.Item{
		{
			SKU:   773297411,
			Count: 1,
		},
	}

	for i, cartItem := range DUMMY_CART {
		productName, productPrice, err := m.Deps.GetProduct(ctx, cartItem.SKU)
		if err != nil {
			return nil, 0, errors.WithMessage(err, "checking product")
		}
		DUMMY_CART[i].Name = productName
		DUMMY_CART[i].Price = productPrice
	}

	return DUMMY_CART, GetTotalPrice(DUMMY_CART), nil
	//end of TODO
}

func GetTotalPrice(items []model.Item) uint32 {
	var totalPrice uint32

	for _, item := range items {
		totalPrice += item.Price * uint32(item.Count)
	}

	return totalPrice
}
