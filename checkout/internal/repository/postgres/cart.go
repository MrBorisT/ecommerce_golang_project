package repository

import (
	"context"
	"route256/checkout/internal/model"

	"github.com/jackc/pgx/v4/pgxpool"
)

type cartRepo struct {
	pool *pgxpool.Pool
}

func NewCartRepo(pool *pgxpool.Pool) *cartRepo {
	return &cartRepo{
		pool: pool,
	}
}

func (r *cartRepo) AddToCart(ctx context.Context, user int64, sku uint32, count uint16) error {
	return nil
}

func (r *cartRepo) DeleteFromCart(ctx context.Context, user int64, sku uint32, count uint16) error {
	return nil

}

func (r *cartRepo) ListCart(ctx context.Context, user int64) ([]model.Item, uint32, error) {
	return nil, 0, nil
}

func (r *cartRepo) Purchase(ctx context.Context, user int64) error {
	return nil
}
