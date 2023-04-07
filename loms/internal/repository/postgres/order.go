package repository

import (
	"context"
	"route256/libs/transactor"
	"route256/loms/internal/model"

	sq "github.com/Masterminds/squirrel"
)

type OrdersRepo struct {
	transactor.QueryEngineProvider
}

func NewOrdersRepo(provider transactor.QueryEngineProvider) *OrdersRepo {
	return &OrdersRepo{
		QueryEngineProvider: provider,
	}
}

const ordersTable = "orders"

func (r *OrdersRepo) CancelOrder(ctx context.Context, orderID int64) error {
	db := r.QueryEngineProvider.GetQueryEngine(ctx)
	query := sq.Update(ordersTable).
		Set("status", "cancelled").
		Where(sq.Eq{"id": orderID})

	queryRaw, args, err := query.ToSql()
	if err != nil {
		return err
	}

	_, err = db.Query(ctx, queryRaw, args...)
	return err
}

func (r *OrdersRepo) CreateOrder(ctx context.Context, user int64, items []model.Item) (int64, error) {
	db := r.QueryEngineProvider.GetQueryEngine(ctx)
	queryCreateOrder :=
		sq.Insert(ordersTable).
			Columns("status", "user_id").
			Values("new", user).
			Suffix("RETURNING id")

	queryCreateOrderRaw, args, err := queryCreateOrder.ToSql()
	if err != nil {
		return 0, nil
	}

	row, err := db.Query(ctx, queryCreateOrderRaw, args...)

	if err != nil {
		return 0, err
	}

	var orderID int64

	if err := row.Scan(&orderID); err != nil {
		return 0, err
	}

	for _, item := range items {
		queryCreateOrderItems := sq.Insert(ordersTable).
			Columns("order_id", "sku", "count").
			Values(orderID, item.SKU, item.Count)

		queryCreateOrderItemsRaw, args, err := queryCreateOrderItems.ToSql()
		if err != nil {
			return 0, err
		}

		_, err = db.Query(ctx, queryCreateOrderItemsRaw, args...)
		if err != nil {
			return orderID, err
		}
	}

	return orderID, nil
}

func (r *OrdersRepo) ListOrder(ctx context.Context, orderID int64) (string, int64, []model.Item, error) {
	db := r.QueryEngineProvider.GetQueryEngine(ctx)

	query := sq.Select("status", "user_id").
		From(ordersTable).
		Where(sq.Eq{"id": orderID})

	queryRaw, args, err := query.ToSql()
	if err != nil {
		return "", 0, nil, err
	}

	row, err := db.Query(ctx, queryRaw, args...)
	if err != nil {
		return "", 0, nil, err
	}

	var status string
	var userID int64

	if err := row.Scan(&status, &userID); err != nil {
		return "", 0, nil, err
	}

	query = sq.Select("sku", "count").
		From(ordersTable).
		Where(sq.Eq{"order_id": orderID})

	queryRaw, args, err = query.ToSql()
	if err != nil {
		return "", 0, nil, err
	}

	rowsItems, err := db.Query(ctx, queryRaw, args...)

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
	db := r.QueryEngineProvider.GetQueryEngine(ctx)
	query :=
		sq.Update(ordersTable).
			Set("status", "payed").
			Where(sq.Eq{"id": orderID})

	queryRaw, args, err := query.ToSql()
	if err != nil {
		return err
	}

	_, err = db.Query(ctx, queryRaw, args...)

	return err
}

func (r *OrdersRepo) OrderFailed(ctx context.Context, orderID int64) error {
	db := r.QueryEngineProvider.GetQueryEngine(ctx)
	query :=
		sq.Update(ordersTable).
			Set("status", "failed").
			Where(sq.Eq{"id": orderID})

	queryRaw, args, err := query.ToSql()
	if err != nil {
		return err
	}

	_, err = db.Query(ctx, queryRaw, args...)

	return err
}

func (r *OrdersRepo) OrderAwaitPayment(ctx context.Context, orderID int64) error {
	db := r.QueryEngineProvider.GetQueryEngine(ctx)
	query :=
		sq.Update(ordersTable).
			Set("status", "awaiting payment").
			Where(sq.Eq{"id": orderID})

	queryRaw, args, err := query.ToSql()
	if err != nil {
		return err
	}

	_, err = db.Query(ctx, queryRaw, args...)

	return err
}
