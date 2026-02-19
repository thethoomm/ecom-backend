-- name: ListProducts :many
SELECT * FROM products;

-- name: FindProductById :one
SELECT * FROM products WHERE id = $1;

-- name: CreateUser :one
INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id, name, email, created_at;
