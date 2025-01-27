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
		userID, err := service.GetUserIDFromSessionToken(db, r)
		if err != nil {
			pd.LoggedIn = false
		}
		if pd.LoggedIn {
			pd.Expenses, _ = service.GetExpenses(db, userID)
			pd.CurrentMonthSum, _ = service.GetCurrentMonthSum(db, userID)
		}
		t.ExecuteTemplate(w, "index", pd)
	}
}
