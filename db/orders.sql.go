// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: orders.sql

package db

import (
	"context"
)

const createOrders = `-- name: CreateOrders :exec
INSERT INTO orders (
    id,
    product_id,
    user_id,
    price,
    quantity_ordered,
    status
)
VALUES (
    $1, $2, $3, $4, $5,$6
) RETURNING id, product_id, user_id, price, status, quantity_ordered
`

type CreateOrdersParams struct {
	ID              int64  `json:"id"`
	ProductID       int64  `json:"product_id"`
	UserID          int64  `json:"user_id"`
	Price           int64  `json:"price"`
	QuantityOrdered int64  `json:"quantity_ordered"`
	Status          string `json:"status"`
}

func (q *Queries) CreateOrders(ctx context.Context, arg CreateOrdersParams) error {
	_, err := q.db.ExecContext(ctx, createOrders,
		arg.ID,
		arg.ProductID,
		arg.UserID,
		arg.Price,
		arg.QuantityOrdered,
		arg.Status,
	)
	return err
}

const orderByStatus = `-- name: OrderByStatus :one
SELECT id, product_id, user_id, price, status, quantity_ordered FROM orders WHERE status = $1
`

func (q *Queries) OrderByStatus(ctx context.Context, status string) (Order, error) {
	row := q.db.QueryRowContext(ctx, orderByStatus, status)
	var i Order
	err := row.Scan(
		&i.ID,
		&i.ProductID,
		&i.UserID,
		&i.Price,
		&i.Status,
		&i.QuantityOrdered,
	)
	return i, err
}

const orderByUserId = `-- name: OrderByUserId :one
SELECT id, product_id, user_id, price, status, quantity_ordered FROM orders WHERE user_id = $1
`

func (q *Queries) OrderByUserId(ctx context.Context, userID int64) (Order, error) {
	row := q.db.QueryRowContext(ctx, orderByUserId, userID)
	var i Order
	err := row.Scan(
		&i.ID,
		&i.ProductID,
		&i.UserID,
		&i.Price,
		&i.Status,
		&i.QuantityOrdered,
	)
	return i, err
}

const updateOrderQuantity = `-- name: UpdateOrderQuantity :exec
UPDATE orders
    SET 
        quantity_ordered = COALESCE($2, quantity_ordered)
    WHERE id=$1 and product_id=$2 and user_id=$3
`

type UpdateOrderQuantityParams struct {
	ID              int64 `json:"id"`
	QuantityOrdered int64 `json:"quantity_ordered"`
	UserID          int64 `json:"user_id"`
}

func (q *Queries) UpdateOrderQuantity(ctx context.Context, arg UpdateOrderQuantityParams) error {
	_, err := q.db.ExecContext(ctx, updateOrderQuantity, arg.ID, arg.QuantityOrdered, arg.UserID)
	return err
}

const updateOrderStatus = `-- name: UpdateOrderStatus :exec
UPDATE orders
    SET 
        status = $4
    WHERE id=$1 and product_id=$2 and user_id=$3
`

type UpdateOrderStatusParams struct {
	ID        int64  `json:"id"`
	ProductID int64  `json:"product_id"`
	UserID    int64  `json:"user_id"`
	Status    string `json:"status"`
}

func (q *Queries) UpdateOrderStatus(ctx context.Context, arg UpdateOrderStatusParams) error {
	_, err := q.db.ExecContext(ctx, updateOrderStatus,
		arg.ID,
		arg.ProductID,
		arg.UserID,
		arg.Status,
	)
	return err
}

const orders = `-- name: orders :many
SELECT id, product_id, user_id, price, status, quantity_ordered FROM orders
`

func (q *Queries) orders(ctx context.Context) ([]Order, error) {
	rows, err := q.db.QueryContext(ctx, orders)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Order
	for rows.Next() {
		var i Order
		if err := rows.Scan(
			&i.ID,
			&i.ProductID,
			&i.UserID,
			&i.Price,
			&i.Status,
			&i.QuantityOrdered,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}