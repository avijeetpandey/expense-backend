package types

type Expense struct {
	Name  string `db:"name"`
	Price int    `db:"price"`
	Tag   string `db:"tag"`
}
