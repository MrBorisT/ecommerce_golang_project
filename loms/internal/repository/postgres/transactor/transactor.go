package transactor

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type QueryEngine interface {
	Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error)
}

type QueryEngineProvider interface {
	GetQueryEngine(ctx context.Context) QueryEngine // tx/pool
}

type TransactionManager struct {
	pool *pgxpool.Pool
}

func NewTransactionManager(pool *pgxpool.Pool) TransactionManager {
	return TransactionManager{pool}
}

type txkey string

const key = txkey("tx")

func (tm TransactionManager) RunRepeatableRead(ctx context.Context, fx func(ctxTX context.Context) error) error {
	tx, err := tm.pool.BeginTx(ctx,
		pgx.TxOptions{
			IsoLevel: pgx.RepeatableRead,
		})
	if err != nil {
		return err
	}

	if err := fx(context.WithValue(ctx, key, tx)); err != nil {
		tx.Rollback(ctx)
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		tx.Rollback(ctx)
		return err
	}

	return nil
}

func (tm TransactionManager) GetQueryEngine(ctx context.Context) QueryEngine {
	tx, ok := ctx.Value(key).(QueryEngine)
	if ok && tx != nil {
		return tx
	}

	return tm.pool
}
