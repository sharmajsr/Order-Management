-- name: CreateUser :exec
INSERT INTO users (
    id,
    name,
    username,
    password
)
VALUES (
    $1, $2, $3, $4
) RETURNING *;


-- name: UpdateNameOfUser :exec
UPDATE users
    SET 
        name = $2
    WHERE id=$1;


-- name: UserById :one 
SELECT * FROM users WHERE id = $1;

-- name: Users :many
SELECT * FROM users;

-- name: UpdatePasswordOfUser :exec
UPDATE users
    SET 
        password = $2
    WHERE id=$1;

