package handler

import (
	"database/sql"
	"html/template"
	"kashtrack/service"
	"net/http"
)

func LoginHandler(db *sql.DB, t *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sessionToken, _ := service.Login(db, r)
		cookie := http.Cookie{
			Name:  "session_token",
			Value: sessionToken,
		}
		http.SetCookie(w, &cookie)
		w.Header().Add("HX-Redirect", "/")
	}
}
