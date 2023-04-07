package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"
)

type offsetsRepo struct {
	pool *pgxpool.Pool
}

func NewOffsetsRepo(pool *pgxpool.Pool) *offsetsRepo {
	return &offsetsRepo{
		pool: pool,
	}
}

var (
	table   = "offsets"
	columns = []string{"id", "offset"}
)

func (r *offsetsRepo) CreateRepo(ctx context.Context, partitionID int32) error {
	query := sq.Insert(table).
		Columns(columns...).
		Values(partitionID, 0)

	queryRaw, args, err := query.ToSql()
	if err != nil {
		return err
	}

	_, err = r.pool.Query(ctx, queryRaw, args...)
	return err
}

func (r *offsetsRepo) GetOffsetForRepo(ctx context.Context, partitionID int32) (int64, error) {
	query := sq.Select("offset").
		From(table).
		Where(sq.Eq{"id": partitionID})

	queryRaw, args, err := query.ToSql()
	if err != nil {
		return 0, err
	}

	var offset int64
	if err := r.pool.QueryRow(ctx, queryRaw, args...).Scan(&offset); err != nil {
		return 0, err
	}

	return offset, nil
}

func (r *offsetsRepo) GetOffsets(ctx context.Context) (map[int32]int64, error) {
	query := sq.Select(columns...).
		From(table)

	queryRaw, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	offsets := make(map[int32]int64)
	if row, err := r.pool.Query(ctx, queryRaw, args...); err != nil {
		return nil, err
	} else {
		var id int32
		var offset int64
		for row.Next() {
			if err := row.Scan(&id, &offset); err != nil {
				return nil, err
			}
			offsets[id] = offset
		}
	}

	return offsets, nil
}

func (r *offsetsRepo) UpdateOffset(ctx context.Context, partitionID int32, newOffset int64) error {
	query := sq.Update(table).
		Set("offset", newOffset).
		Where(sq.Eq{"id": partitionID})

	queryRaw, args, err := query.ToSql()
	if err != nil {
		return err
	}

	_, err = r.pool.Query(ctx, queryRaw, args...)
	return err
}
