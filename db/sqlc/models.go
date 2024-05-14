// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"database/sql"
	"time"
)

type Expense struct {
	ID        int64          `json:"id"`
	Name      string         `json:"name"`
	Balance   int64          `json:"balance"`
	Currency  string         `json:"currency"`
	Tag       sql.NullString `json:"tag"`
	CreatedAt time.Time      `json:"created_at"`
}
