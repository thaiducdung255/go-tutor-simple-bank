-- name: CreateAccount :one
insert into accounts (
	owner,
	balance,
	currency
) values (
	$1, $2, $3
)
returning *;

-- name: GetAccount :one
select * from accounts
where id = $1 limit 1;

-- name: ListAccounts :many
select * from accounts
order by id desc
limit $1
offset $2;

-- name: UpdateAccount :one
update accounts set owner = $1
where id = $1
returning *;

-- name: DeleteAccount :exec
delete from accounts
where id = $1;
