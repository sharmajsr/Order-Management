-- name: CreateProduct :exec
INSERT INTO products (
    id,
    price,
    details,
    quantity_available
)
VALUES (
    $1, $2, $3, $4
) RETURNING *;


-- name: UpdateProduct :exec
UPDATE products
    SET 
        price = COALESCE($2, price),
        details = COALESCE($3, details),
        quantity_available = COALESCE($4, quantity_available)
    WHERE id=$1;


-- name: ProductById :one 
SELECT * FROM products WHERE id = $1;

-- name: Products :many
SELECT * FROM products;

