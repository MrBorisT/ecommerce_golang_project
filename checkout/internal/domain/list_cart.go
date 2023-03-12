package domain

import (
	"context"
	"route256/checkout/internal/model"

	"github.com/pkg/errors"
)

func (m *CheckoutService) ListCart(ctx context.Context, user int64) ([]model.Item, uint32, error) {
	cart, err := m.CartRepository.ListCart(ctx, user)
	if err != nil {
		return nil, 0, err
	}

	items := make([]model.Item, 0, len(cart.Items))
	for _, cartItem := range cart.Items {
		productName, productPrice, err := m.ProductChecker.GetProduct(ctx, cartItem.SKU)
		if err != nil {
			return nil, 0, errors.WithMessage(err, "checking product")
		}
		items = append(items, model.Item{
			SKU: cartItem.SKU,
			Count: cartItem.Count,
			Name: productName,
			Price: productPrice,
		})
	}

	return items, GetTotalPrice(items), nil
}

func GetTotalPrice(items []model.Item) uint32 {
	var totalPrice uint32

	for _, item := range items {
		totalPrice += item.Price * uint32(item.Count)
	}

	return totalPrice
}
