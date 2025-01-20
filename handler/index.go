package handler

import (
	"database/sql"
	"html/template"
	"kashtrack/model"
	"kashtrack/service"
	"net/http"
)

func IndexHandler(db *sql.DB, t *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pd := model.PageData{}
		pd.LoggedIn = true
		_, err := service.ValidateSession(db, r)
		if err != nil {
			pd.LoggedIn = false
		}
		pd.Expenses, _ = service.GetExpenses(db)
		pd.CurrentMonthSum, _ = service.GetCurrentMonthSum(db)
		t.ExecuteTemplate(w, "index", pd)
	}
}
