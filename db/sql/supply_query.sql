-- name: DeleteSupply :one
delete from supply where uuid = $1 returning uuid;

-- name: UpdateSupplyInfo :exec
update supply set comment = $1, desired_date = $2, status = $3, responsible_user = $4, edited = true, edited_date = now(), cost = $5 where uuid = $6;

-- name: GetActiveSupplies :many
select * from supply where status != 'done';

-- name: GetSupplyById :one
select * from supply where uuid = $1;