package handler

import (
	"database/sql"
	"encoding/json"
	"html/template"
	"kashtrack/service"
	"net/http"
)

func AddExpenseHandler(db *sql.DB, t *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		service.AddExpense(db, r)
		w.Header().Add("HX-Redirect", "/")
	}
}

func ExpenseDistributionHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		expenses := service.GetExpenseDistribution(db)
		data, _ := json.Marshal(expenses)
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}
