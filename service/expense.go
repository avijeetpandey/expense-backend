package service

import (
	"expense-service/database"
	"expense-service/dataproviders"
	"expense-service/types"
	"fmt"
)

// service to add expense
func AddExpense(ex types.Expense) error {
	db, err := database.GetDBInstance()

	if err != nil {
		fmt.Println("Error getting database instance")
		return err
	}

	dataproviders.AddExpense(db, ex)

	return nil
}

// service to get all the expenses
func GetAllExpense(ex types.Expense) ([]types.Expense, error) {
	db, err := database.GetDBInstance()

	if err != nil {
		fmt.Println("Error getting database instance")
		return nil, err
	}

	expenses, err := dataproviders.GetAllExpense(db)

	return expenses, err
}
