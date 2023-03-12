package repository

import (
	"context"
	"route256/checkout/internal/model"
	"route256/checkout/internal/repository/schema"

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

func (r *cartRepo) ListCart(ctx context.Context, user int64) (*model.Cart, error) {
	return nil, nil
}

func (r *cartRepo) Purchase(ctx context.Context, user int64) error {
	return nil
}

func bindSchemaCartToModelsCart(cart schema.Cart) model.Cart {
	items := make([]model.CartItem, 0, len(cart.Items))
	for _, item := range items {
		items = append(items, model.CartItem{
			SKU:   item.SKU,
			Count: item.Count,
		})
	}
	return model.Cart{
		User:  cart.User,
		Items: items,
	}
}
