// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: query.sql

package products

import (
	"context"
)

const createProduct = `-- name: CreateProduct :one
INSERT INTO products(user_id, name, description, price) VALUES ($1,$2,$3, $4) RETURNING id
`

type CreateProductParams struct {
	UserID      int32
	Name        string
	Description string
	Price       int32
}

func (q *Queries) CreateProduct(ctx context.Context, arg CreateProductParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, createProduct,
		arg.UserID,
		arg.Name,
		arg.Description,
		arg.Price,
	)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const getAllProducts = `-- name: GetAllProducts :many
SELECT id, user_id, name, description, price, upload_date FROM products
`

func (q *Queries) GetAllProducts(ctx context.Context) ([]Product, error) {
	rows, err := q.db.QueryContext(ctx, getAllProducts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Product
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Name,
			&i.Description,
			&i.Price,
			&i.UploadDate,
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

const getRecentProducts = `-- name: GetRecentProducts :many
SELECT id, user_id, name, description, price, upload_date FROM products ORDER BY upload_date DESC LIMIT 10
`

func (q *Queries) GetRecentProducts(ctx context.Context) ([]Product, error) {
	rows, err := q.db.QueryContext(ctx, getRecentProducts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Product
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Name,
			&i.Description,
			&i.Price,
			&i.UploadDate,
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