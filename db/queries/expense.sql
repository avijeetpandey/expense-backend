-- name: CreateExpense :one
INSERT into expenses (
    name,
    balance,
    currency,
    tag
) values (
    $1, $2, $3, $4
    ) RETURNING *;


-- name: GetExpense :one
SELECT * from expenses where id = $1 LIMIT 1;

-- name: ListExpenses :many
SELECT * from expenses ORDER BY id LIMIT $1 OFFSET $2;

-- name: UpdateExpense :exec
UPDATE expenses set balance = $2 WHERE id = $1;

-- name: DeleteExpense :exec
DELETE from expenses where id = $1;