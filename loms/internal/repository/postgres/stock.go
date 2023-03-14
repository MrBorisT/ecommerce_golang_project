package repository

import (
	"context"
	"route256/loms/internal/model"

	"github.com/jackc/pgx/v4/pgxpool"
)

type StocksRepo struct {
	pool *pgxpool.Pool
}

func NewStocksRepo(pool *pgxpool.Pool) *StocksRepo {
	return &StocksRepo{
		pool: pool,
	}
}

func (r *StocksRepo) Stocks(ctx context.Context, SKU uint32) ([]model.Stock, error) {
	const query = `
	SELECT warehouse_id, count
	FROM stocks
	WHERE sku = $1 AND reserved = FALSE`

	stocks := make([]model.Stock, 0)

	rows, err := r.pool.Query(ctx, query, SKU)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		newStock := model.Stock{}
		if err := rows.Scan(&newStock.WarehouseID, &newStock.Count); err != nil {
			return nil, err
		}
		stocks = append(stocks, newStock)
	}

	return stocks, nil
}
