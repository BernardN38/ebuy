-- name: CreateProduct :one
INSERT INTO products(user_id, name, description, price) VALUES ($1,$2,$3, $4) RETURNING id;
-- name: GetAllProducts :many
SELECT * FROM products;
-- name: GetRecentProducts :many
SELECT * FROM products ORDER BY upload_date DESC LIMIT 10;