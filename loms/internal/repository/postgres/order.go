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

func (r *OrdersRepo) CancelOrder(ctx context.Context, orderID int64) error {
	const query = `
	DELETE FROM orders
	WHERE order_id = $1`

	_, err := r.pool.Query(ctx, query, orderID)
	return err
}

func (r *OrdersRepo) CreateOrder(ctx context.Context, user int64, items []model.Item) (int64, error) {
	const queryCreateOrder = `
	INSERT INTO orders (status, user_id)
	VALUES ('new', $1)
	RETURNING id`

	row := r.pool.QueryRow(ctx, queryCreateOrder, user)

	var orderID int64

	if err := row.Scan(&orderID); err != nil {
		return 0, err
	}

	for _, item := range items {
		const queryCreateOrderItems = `
	INSERT INTO order_items (order_id, sku, count)
	VALUES ($1, $2, $3)`
		_, err := r.pool.Query(ctx, queryCreateOrderItems, orderID, item.SKU, item.Count)
		if err != nil {
			return orderID, err
		}
	}

	return orderID, nil
}
func (r *OrdersRepo) ListOrder(ctx context.Context, orderID int64) (string, int64, []model.Item, error) {
	const query = `
	SELECT status, user_id
	FROM order
	WHERE id = $1`

	row := r.pool.QueryRow(ctx, query, orderID)

	var status string
	var userID int64

	if err := row.Scan(&status, &userID); err != nil {
		return "", 0, nil, err
	}

	const queryItems = `
	SELECT sku, count
	FROM order_items
	WHERE order_id = $1`

	rowsItems, err := r.pool.Query(ctx, queryItems, orderID)

	if err != nil {
		return status, userID, nil, err
	}

	items := make([]model.Item, 0)
	for rowsItems.Next() {
		item := model.Item{}
		if err := rowsItems.Scan(
			&item.SKU,
			&item.Count,
		); err != nil {
			return status, userID, nil, err
		}

		items = append(items, item)
	}

	return status, userID, items, nil
}
func (r *OrdersRepo) OrderPayed(ctx context.Context, orderID int64) error {
	const query = `
	UPDATE orders
	SET status = 'payed'
	WHERE id = $1`

	_, err := r.pool.Query(ctx, query, orderID)

	return err
}
