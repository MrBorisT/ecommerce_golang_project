package repository

import (
	"context"
	"fmt"
	"route256/checkout/internal/model"

	sq "github.com/Masterminds/squirrel"
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

var (
	itemsColumns = []string{"user_id", "sku", "count"}
)

const (
	itemsTable = "cart_items"
)

func (r *cartRepo) AddToCart(ctx context.Context, user int64, sku uint32, count uint16) error {
	insertQuery := sq.Insert(fmt.Sprintf("%s as s", itemsTable)).
		Columns(itemsColumns...).
		Values(user, sku, count).
		Suffix("ON CONFLICT (user_id, sku) DO UPDATE SET count = EXCLUDED.count + c.count")

	insertQueryRaw, args, err := insertQuery.ToSql()
	if err != nil {
		return err
	}

	_, err = r.pool.Query(ctx, insertQueryRaw, args...)
	return err
}

func (r *cartRepo) DeleteFromCart(ctx context.Context, user int64, sku uint32, count uint16) error {
	query := sq.Update(itemsTable).
		Set("count", sq.Expr("count - ?", count)).
		Where(sq.And{sq.Eq{"user_id": user, "sku": sku}}).
		Suffix("RETURNING count")
	queryRaw, args, err := query.ToSql()
	if err != nil {
		return err
	}
	row, err := r.pool.Query(ctx, queryRaw, args...)
	if err != nil {
		return err
	}

	var countAfterUpdate int
	if err := row.Scan(&count); err != nil {
		return err
	}

	if countAfterUpdate <= 0 {
		const deleteQuery = `
	DELETE FROM cart_items
	WHERE user_id = $1 AND sku = $2`

		_, err := r.pool.Query(ctx, deleteQuery, user, sku)
		return err
	}

	return nil
}

func (r *cartRepo) ListCart(ctx context.Context, user int64) (*model.Cart, error) {
	query := sq.Select("sku", "count").
		From(itemsTable).
		Where("user_id = ?", user)

	queryRaw, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.pool.Query(ctx, queryRaw, args...)
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
	query := sq.Delete(itemsTable).
		Where("user_id", user)

	queryRaw, args, err := query.ToSql()
	if err != nil {
		return err
	}

	_, err = r.pool.Query(ctx, queryRaw, args...)
	return err
}
