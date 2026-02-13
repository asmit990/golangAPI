-- name: ListProducts :many
SELECT *
FROM products
ORDER BY id;

-- name: FindProductById :one
SELECT *
FROM products
WHERE id = $1;

-- name: CreateProduct :one
INSERT INTO products (
  name,
  price_in_centers,
  quantity
)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateProduct :one
UPDATE products
SET name = $2,
    price_in_centers = $3,
    quantity = $4
WHERE id = $1
RETURNING *;

-- name: DeleteProduct :exec
DELETE FROM products
WHERE id = $1;
