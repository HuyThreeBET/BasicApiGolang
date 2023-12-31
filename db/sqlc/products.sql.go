// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: products.sql

package sqlc

import (
	"context"
)

const createProduct = `-- name: CreateProduct :one
INSERT INTO products (
  product_name, kind_of_product, owner, currency, price, quantity
) VALUES (
  $1, $2, $3, $4, $5, $6
) RETURNING id_product, product_name, kind_of_product, owner, currency, price, quantity, created_at, update_at
`

type CreateProductParams struct {
	ProductName   string `json:"product_name"`
	KindOfProduct string `json:"kind_of_product"`
	Owner         string `json:"owner"`
	Currency      string `json:"currency"`
	Price         int32  `json:"price"`
	Quantity      int32  `json:"quantity"`
}

func (q *Queries) CreateProduct(ctx context.Context, arg CreateProductParams) (Product, error) {
	row := q.db.QueryRowContext(ctx, createProduct,
		arg.ProductName,
		arg.KindOfProduct,
		arg.Owner,
		arg.Currency,
		arg.Price,
		arg.Quantity,
	)
	var i Product
	err := row.Scan(
		&i.IDProduct,
		&i.ProductName,
		&i.KindOfProduct,
		&i.Owner,
		&i.Currency,
		&i.Price,
		&i.Quantity,
		&i.CreatedAt,
		&i.UpdateAt,
	)
	return i, err
}

const deleteProduct = `-- name: DeleteProduct :exec
DELETE FROM products
WHERE id_product = $1
`

func (q *Queries) DeleteProduct(ctx context.Context, idProduct int32) error {
	_, err := q.db.ExecContext(ctx, deleteProduct, idProduct)
	return err
}

const getProduct = `-- name: GetProduct :one
SELECT id_product, product_name, kind_of_product, owner, currency, price, quantity, created_at, update_at FROM products
WHERE id_product = $1 
LIMIT 1
`

func (q *Queries) GetProduct(ctx context.Context, idProduct int32) (Product, error) {
	row := q.db.QueryRowContext(ctx, getProduct, idProduct)
	var i Product
	err := row.Scan(
		&i.IDProduct,
		&i.ProductName,
		&i.KindOfProduct,
		&i.Owner,
		&i.Currency,
		&i.Price,
		&i.Quantity,
		&i.CreatedAt,
		&i.UpdateAt,
	)
	return i, err
}

const listProducts = `-- name: ListProducts :many
SELECT id_product, product_name, kind_of_product, owner, currency, price, quantity, created_at, update_at FROM products
WHERE owner = $1
ORDER BY id_product
`

func (q *Queries) ListProducts(ctx context.Context, owner string) ([]Product, error) {
	rows, err := q.db.QueryContext(ctx, listProducts, owner)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Product{}
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.IDProduct,
			&i.ProductName,
			&i.KindOfProduct,
			&i.Owner,
			&i.Currency,
			&i.Price,
			&i.Quantity,
			&i.CreatedAt,
			&i.UpdateAt,
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

const updateQuantityOfProduct = `-- name: UpdateQuantityOfProduct :one
UPDATE products
SET quantity = quantity - $1
WHERE id_product = $2
RETURNING id_product, product_name, kind_of_product, owner, currency, price, quantity, created_at, update_at
`

type UpdateQuantityOfProductParams struct {
	Amount    int32 `json:"amount"`
	IDProduct int32 `json:"id_product"`
}

func (q *Queries) UpdateQuantityOfProduct(ctx context.Context, arg UpdateQuantityOfProductParams) (Product, error) {
	row := q.db.QueryRowContext(ctx, updateQuantityOfProduct, arg.Amount, arg.IDProduct)
	var i Product
	err := row.Scan(
		&i.IDProduct,
		&i.ProductName,
		&i.KindOfProduct,
		&i.Owner,
		&i.Currency,
		&i.Price,
		&i.Quantity,
		&i.CreatedAt,
		&i.UpdateAt,
	)
	return i, err
}
