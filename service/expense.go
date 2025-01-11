package service

import (
	"database/sql"
	"kashtrack/model"
)

func GetExpenses(db *sql.DB) ([]model.Expense, error) {
	expenses := []model.Expense{}
	rows, _ := db.Query("SELECT timestamp, category, amount FROM expense ORDER BY timestamp DESC")
	for rows.Next() {
		e := model.Expense{}
		rows.Scan(&e.TimeStamp, &e.Category, &e.Amount)
		e.TimeStamp = e.TimeStamp[:10]
		expenses = append(expenses, e)
	}
	return expenses, nil
}