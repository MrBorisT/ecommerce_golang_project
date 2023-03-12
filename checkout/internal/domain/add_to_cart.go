package domain

import (
	"context"

	"github.com/pkg/errors"
)

var (
	ErrInsufficientStocks = errors.New("insufficient stocks")
)

func (m *CheckoutService) AddToCart(ctx context.Context, user int64, sku uint32, count uint16) error {
	stocks, err := m.LOMS.Stocks(ctx, sku)
	if err != nil {
		return errors.WithMessage(err, "checking stocks")
	}

	counter := int64(count)
	for _, stock := range stocks {
		counter -= int64(stock.Count)
		if counter <= 0 {
			return m.CartRepository.AddToCart(ctx, user, sku, count)
		}
	}

	return ErrInsufficientStocks
}
