-- name: ListProducts :many
SELECT * FROM products;

-- name: FindProductById :one
SELECT * FROM products WHERE id = $1;

-- name: UpdateProductQuantity :exec
UPDATE products SET quantity = $2 WHERE id = $1;

-- name: CreateUser :one
INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id, name, email, created_at;

-- name: CreateOrder :one
INSERT INTO orders (customer_id, status) VALUES ($1, $2) RETURNING *;

-- name: CreateOrderItem :one
INSERT INTO order_items (order_id, product_id, price, quantity) VALUES ($1, $2, $3, $4) RETURNING *;
