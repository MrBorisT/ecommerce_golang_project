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

func (r *StocksRepo) ReserveStocks(ctx context.Context, SKU uint32, count uint16) error {
	const query = `
	SELECT warehouse_id, count
	FROM stocks
	WHERE sku = $1 AND reserved = FALSE`

	rows, err := r.pool.Query(ctx, query, SKU)
	if err != nil {
		return err
	}

	for rows.Next() && count > 0 {
		var stockWarehouseID int64
		var stockCount uint64
		if err := rows.Scan(&stockWarehouseID, &stockCount); err != nil {
			return err
		}

		if stockCount >= uint64(count) {
			// warehouse has more or equal amount
			if stockCount == uint64(count) {
				// equal amounts - reserve whole warehouse
				const queryUpdate = `
		UPDATE stocks
		SET reserved = TRUE
		WHERE warehouse_id = $1 AND sku = $2 AND reserved = FALSE`
				_, err := r.pool.Query(ctx, queryUpdate, stockWarehouseID, SKU)
				if err != nil {
					return err
				}
			} else {
				// in warehouse amount is less
				// decrease reserved amount in warehouse
				const queryUpdate = `
		UPDATE stocks
		SET reserved = TRUE, count = $1
		WHERE warehouse_id = $2 AND sku = $3 AND reserved = FALSE`
				_, err := r.pool.Query(ctx, queryUpdate, stockCount, stockWarehouseID, SKU)
				if err != nil {
					return err
				}
				// then insert into warehouse with reserved = true
				const insertQuery = `
		INSERT INTO stocks (warehouse_id, sku, count, reserved)
		VALUES ($1, $2, $3, TRUE)`
				_, err = r.pool.Query(ctx, insertQuery, stockWarehouseID, SKU, count)
				if err != nil {
					return err
				}
			}
			count = 0
		} else {
			// warehouse has less - reserve whole warehouse
			const queryUpdate = `
			UPDATE stocks
			SET reserved = TRUE
			WHERE warehouse_id = $1 AND sku = $2 AND reserved = FALSE`
			_, err := r.pool.Query(ctx, queryUpdate, stockWarehouseID, SKU)
			if err != nil {
				return err
			}
			// then go to next
			count -= uint16(stockCount)
		}
	}
	return nil
}

func (r *StocksRepo) UnreserveStocks(ctx context.Context, SKU uint32, count uint16) error {
	const query = `
	SELECT warehouse_id, count
	FROM stocks
	WHERE sku = $1 AND reserved = TRUE`

	rows, err := r.pool.Query(ctx, query, SKU)
	if err != nil {
		return err
	}

	for rows.Next() && count > 0 {
		var stockWarehouseID int64
		var stockCount uint64
		if err := rows.Scan(&stockWarehouseID, &stockCount); err != nil {
			return err
		}

		if stockCount >= uint64(count) {
			// warehouse has more or equal amount
			if stockCount == uint64(count) {
				// equal amounts - reserve whole warehouse
				const queryUpdate = `
		UPDATE stocks
		SET reserved = FALSE
		WHERE warehouse_id = $1 AND sku = $2 AND reserved = TRUE`
				_, err := r.pool.Query(ctx, queryUpdate, stockWarehouseID, SKU)
				if err != nil {
					return err
				}
			} else {
				// in warehouse amount is less
				// decrease reserved amount in warehouse
				const queryUpdate = `
		UPDATE stocks
		SET reserved = FALSE, count = $1
		WHERE warehouse_id = $2 AND sku = $3 AND reserved = TRUE`
				_, err := r.pool.Query(ctx, queryUpdate, stockCount, stockWarehouseID, SKU)
				if err != nil {
					return err
				}
				// then insert into warehouse with reserved = FALSE
				const insertQuery = `
		INSERT INTO stocks (warehouse_id, sku, count, reserved)
		VALUES ($1, $2, $3, FALSE)`
				_, err = r.pool.Query(ctx, insertQuery, stockWarehouseID, SKU, count)
				if err != nil {
					return err
				}
			}
			count = 0
		} else {
			// warehouse has less - reserve whole warehouse
			const queryUpdate = `
			UPDATE stocks
			SET reserved = TRUE
			WHERE warehouse_id = $1 AND sku = $2 AND reserved = TRUE`
			_, err := r.pool.Query(ctx, queryUpdate, stockWarehouseID, SKU)
			if err != nil {
				return err
			}
			// then go to next
			count -= uint16(stockCount)
		}
	}
	return nil
}
