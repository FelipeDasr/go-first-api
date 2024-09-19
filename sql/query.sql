----- CUSTOMERS -----

-- name: CreateCustomer :one
INSERT INTO customers (name, email) VALUES ($1, $2) RETURNING *;

-- name: CustomerAlreadyExistsByEmail :one
SELECT c.id FROM customers c WHERE c.email = $1 LIMIT 1;

-- name: GetCustomerById :one
SELECT * FROM customers WHERE id = $1;

-- name: GetCustomers :many
SELECT * FROM customers c LIMIT $1 OFFSET $2;


----- PRODUCTS -----


-- name: CreateProduct :one
INSERT INTO products (name, price, stock) VALUES ($1, $2, $3) RETURNING *;

-- name: ProductAlreadyExistsByName :one
SELECT p.id FROM products p WHERE p.name = $1 LIMIT 1;

-- name: GetProductById :one
SELECT * FROM products WHERE id = $1 LIMIT 1;

-- name: GetProducts :many
SELECT * FROM products p LIMIT $1 OFFSET $2;

-- name: IncrementProductStockById :exec
UPDATE products SET stock = stock + $2 WHERE id = $1;


----- ORDERS -----


-- name: CreateOrder :one
INSERT INTO orders (customer_id, product_id, units_amount, unit_price) 
VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetOrderById :one
SELECT * FROM orders WHERE id = $1 LIMIT 1;

-- name: GetManyOrders :many
SELECT * FROM orders o LIMIT $1 OFFSET $2;