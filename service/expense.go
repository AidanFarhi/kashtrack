package service

import (
	"database/sql"
	"kashtrack/model"
	"net/http"
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

func GetCurrentMonthSum(db *sql.DB) (float64, error) {
	var currentMonthSum float64
	row := db.QueryRow(`
		SELECT SUM(amount) FROM expense
		WHERE SUBSTRING(timestamp, 1, 7) = STRFTIME('%Y-%m', DATE('now'))
	`)
	row.Scan(&currentMonthSum)
	return currentMonthSum, nil
}

func AddExpense(db *sql.DB, r *http.Request) error {
	db.Exec(
		"INSERT INTO expense (user_id, category, amount) VALUES (1, ?, ?)",
		r.FormValue("category"), r.FormValue("amount"),
	)
	return nil
}

func GetExpenseDistribution(db *sql.DB) []model.ExpenseJSON {
	expenses := []model.ExpenseJSON{}
	rows, _ := db.Query(`
		SELECT category, ROUND(SUM(amount), 2)
		FROM expense
		WHERE SUBSTRING(timestamp, 1, 7) = STRFTIME('%Y-%m', DATE('now'))
		GROUP BY category
	`)
	for rows.Next() {
		e := model.ExpenseJSON{}
		rows.Scan(&e.Category, &e.Amount)
		expenses = append(expenses, e)
	}
	return expenses
}
