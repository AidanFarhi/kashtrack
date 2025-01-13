package handler

import (
	"database/sql"
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
