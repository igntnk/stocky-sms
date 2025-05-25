-- name: CreateProduct :one
insert into products (product_code) values ($1) returning uuid;

-- name: DeleteProduct :one
delete from products where uuid = $1 returning uuid;

-- name: SetStoreCost :exec
update products set store_cost = $1 where uuid = $2;

-- name: SetStoreAmount :exec
update products set store_amount = $1 where uuid = $2;

-- name: GetStoreAmount :one
select store_amount from products where uuid = $1;
