----- CUSTOMERS -----

-- name: CreateCustomer :one
INSERT INTO customers (name, email) VALUES ($1, $2) RETURNING *;

-- name: CustomerAlreadyExistsByEmail :one
SELECT c.id FROM customers c WHERE c.email = $1 LIMIT 1;

-- name: GetCustomerById :one
SELECT * FROM customers WHERE id = $1;


----- PRODUCTS -----


-- name: CreateProduct :one
INSERT INTO products (name, price, stock) VALUES ($1, $2, $3) RETURNING *;

-- name: ProductAlreadyExistsByName :one
SELECT p.id FROM products p WHERE p.name = $1 LIMIT 1;

-- name: GetProductById :one
SELECT * FROM products WHERE id = $1 LIMIT 1;

-- name: GetProducts :many
SELECT * FROM products p
WHERE
  ($3::TEXT IS NULL OR p."name" ILIKE '%' || $3 || '%') AND -- filter by name
  ($4::INT IS NULL OR p.price >= $4) AND -- filter by min price
  ($5::INT IS NULL OR p.price <= $5) AND -- filter by max price
  ($6::INT IS NULL OR p.stock >= $6) AND -- filter by min stock
  ($7::INT IS NULL OR p.stock <= $7) -- filter by max stock
LIMIT $1 OFFSET $2;

-- name: IncrementProductStockById :exec
UPDATE products SET stock = stock + $2 WHERE id = $1;


----- ORDERS -----


-- name: CreateOrder :one
INSERT INTO orders (customer_id, product_id, units_amount, unit_price) 
VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetOrderById :one
SELECT * FROM orders WHERE id = $1 LIMIT 1;

-- name: GetOrdersByCustomerId :many
SELECT * FROM orders WHERE customer_id = $1 LIMIT $2 OFFSET $3;

-- name: GetOrdersByProductId :many
SELECT * FROM orders WHERE product_id = $1 LIMIT $2 OFFSET $3;