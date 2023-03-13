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
	//TODO
	return nil, nil
}
