-- name: CreateEntry :one
insert into entries (
	account_id,
	amount
) values (
	$1, $2
)
returning *;

-- name: GetEntry :one
select * from entries
where id = $1 limit 1;

-- name: ListEntries :many
select * from entries
order by id desc
limit $1
offset $2;

-- name: UpdateEntry :one
update entries set account_id = $1
where id = $2
returning *;

-- name: DeleteEntry :exec
delete from entries
where id = $1;
