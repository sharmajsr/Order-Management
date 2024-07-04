-- name: CreateOrders :exec
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
) RETURNING *;


-- name: UpdateOrderStatus :exec
UPDATE orders
    SET 
        status = $4
    WHERE id=$1 and product_id=$2 and user_id=$3;

-- name: UpdateOrderQuantity :exec
UPDATE orders
    SET 
        quantity_ordered = COALESCE($2, quantity_ordered)
    WHERE id=$1 and product_id=$2 and user_id=$3;


-- name: OrderById :one 
SELECT * FROM orders WHERE id = $1;

-- name: OrderByUserId :one 
SELECT * FROM orders WHERE user_id = $1;

-- name: OrderByStatus :one 
SELECT * FROM orders WHERE status = $1;

-- name: orders :many
SELECT * FROM orders;


