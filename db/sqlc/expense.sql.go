// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: expense.sql

package db

import (
	"context"
	"database/sql"
)

const createExpense = `-- name: CreateExpense :one
INSERT into expenses (
    name,
    balance,
    currency,
    tag
) values (
    $1, $2, $3, $4
    ) RETURNING id, name, balance, currency, tag, created_at
`

type CreateExpenseParams struct {
	Name     string         `json:"name"`
	Balance  int64          `json:"balance"`
	Currency string         `json:"currency"`
	Tag      sql.NullString `json:"tag"`
}

func (q *Queries) CreateExpense(ctx context.Context, arg CreateExpenseParams) (Expense, error) {
	row := q.db.QueryRowContext(ctx, createExpense,
		arg.Name,
		arg.Balance,
		arg.Currency,
		arg.Tag,
	)
	var i Expense
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Balance,
		&i.Currency,
		&i.Tag,
		&i.CreatedAt,
	)
	return i, err
}

const deleteExpense = `-- name: DeleteExpense :exec
DELETE from expenses where id = $1
`

func (q *Queries) DeleteExpense(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteExpense, id)
	return err
}

const getExpense = `-- name: GetExpense :one
SELECT id, name, balance, currency, tag, created_at from expenses where id = $1 LIMIT 1
`

func (q *Queries) GetExpense(ctx context.Context, id int64) (Expense, error) {
	row := q.db.QueryRowContext(ctx, getExpense, id)
	var i Expense
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Balance,
		&i.Currency,
		&i.Tag,
		&i.CreatedAt,
	)
	return i, err
}

const listExpenses = `-- name: ListExpenses :many
SELECT id, name, balance, currency, tag, created_at from expenses ORDER BY id LIMIT $1 OFFSET $2
`

type ListExpensesParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListExpenses(ctx context.Context, arg ListExpensesParams) ([]Expense, error) {
	rows, err := q.db.QueryContext(ctx, listExpenses, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Expense
	for rows.Next() {
		var i Expense
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Balance,
			&i.Currency,
			&i.Tag,
			&i.CreatedAt,
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

const updateExpense = `-- name: UpdateExpense :exec
UPDATE expenses set balance = $2 WHERE id = $1
`

type UpdateExpenseParams struct {
	ID      int64 `json:"id"`
	Balance int64 `json:"balance"`
}

func (q *Queries) UpdateExpense(ctx context.Context, arg UpdateExpenseParams) error {
	_, err := q.db.ExecContext(ctx, updateExpense, arg.ID, arg.Balance)
	return err
}