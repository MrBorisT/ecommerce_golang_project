package domain

import (
	"context"
	"route256/loms/internal/model"
)

func (m *service) Stocks(ctx context.Context, SKU uint32) ([]model.Stock, error) {
	return m.StockRepository.Stocks(ctx, SKU)
}
