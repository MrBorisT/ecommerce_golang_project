package repository

import (
	"context"
	"route256/loms/internal/model"

	"github.com/jackc/pgx/v4/pgxpool"
)

type OrdersRepo struct {
	pool *pgxpool.Pool
}

func NewOrdersRepo(pool *pgxpool.Pool) *OrdersRepo {
	return &OrdersRepo{
		pool: pool,
	}
}

//TODO
func (r *OrdersRepo) CancelOrder(ctx context.Context, orderID int64) error {
	return nil
}
func (r *OrdersRepo) CreateOrder(ctx context.Context, user int64, items []model.Item) (int64, error) {
	return 0, nil
}
func (r *OrdersRepo) ListOrder(ctx context.Context, orderID int64) (string, int64, []model.Item, error) {
	return "", 0, nil, nil
}
func (r *OrdersRepo) OrderPayed(ctx context.Context, orderID int64) error {
	return nil
}
