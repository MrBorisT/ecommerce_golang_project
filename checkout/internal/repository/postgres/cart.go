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
	const query = `
	INSERT INTO cart_items as c (user_id, sku, count)
	VALUES ($1,$2,$3)
	ON CONFLICT (user_id, sku) DO UPDATE SET count = EXCLUDED.count + c.count`

	_, err := r.pool.Query(ctx, query, user, sku, count)

	return err
}

func (r *cartRepo) DeleteFromCart(ctx context.Context, user int64, sku uint32, count uint16) error {
	const query = `
	UPDATE cart_items
	SET count = count - $3
	WHERE user_id = $1 AND sku = $2
	RETURNING count`
	row := r.pool.QueryRow(ctx, query, user, sku, count)

	var countAfterUpdate int
	if err := row.Scan(&count); err != nil {
		return err
	}

	if countAfterUpdate <= 0 {
		const deleteQuery = `
	DELETE FROM cart_items
	WHERE user_id = $1 AND sku = $2`

		_, err := r.pool.Query(ctx, query, user, sku)
		return err
	}

	return nil
}

func (r *cartRepo) ListCart(ctx context.Context, user int64) (*model.Cart, error) {
	const query = `
	SELECT sku, count
	FROM cart_items
	WHERE user_id = $1`

	rows, err := r.pool.Query(ctx, query, user)
	if err != nil {
		return nil, err
	}
	var cart *model.Cart

	defer rows.Close()

	cart.Items = make([]model.CartItem, 0)

	for rows.Next() {
		item := model.CartItem{}
		if err = rows.Scan(
			&item.SKU,
			&item.Count,
		); err != nil {
			return nil, err
		}

		cart.Items = append(cart.Items, item)
	}

	return cart, nil
}

func (r *cartRepo) Purchase(ctx context.Context, user int64) error {
	const query = `
	DELETE FROM cart_items
	WHERE user_id = $1`

	_, err := r.pool.Query(ctx, query, user)
	return err
}
