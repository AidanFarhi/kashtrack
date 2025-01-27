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
		userID, err := service.GetUserIDFromSessionToken(db, r)
		if err != nil {
			w.Header().Add("HX-Redirect", "/")
			return
		}
		service.AddExpense(db, r, userID)
		w.Header().Add("HX-Redirect", "/")
	}
}

func ExpenseDistributionHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, err := service.GetUserIDFromSessionToken(db, r)
		if err != nil {
			w.Header().Add("HX-Redirect", "/")
			return
		}
		expenses := service.GetExpenseDistribution(db, userID)
		data, _ := json.Marshal(expenses)
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}
