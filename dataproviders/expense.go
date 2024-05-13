package dataproviders

import (
	"expense-service/types"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// adds an expense in the database
func AddExpense(db *sqlx.DB, ex types.Expense) error {
	_, err := db.Exec("INSERT INTO expenses (name, price, tag) VALUES ($1, $2, $3)", ex.Name, ex.Price, ex.Tag)
	return err
}

// provides all the expenses
func GetAllExpense(db *sqlx.DB) ([]types.Expense, error) {
	var expenses []types.Expense
	err := db.Select(&expenses, "SELECT * from expenses")

	return expenses, err
}
