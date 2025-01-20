package handler

import (
	"database/sql"
	"html/template"
	"kashtrack/service"
	"net/http"
)

func LoginHandler(db *sql.DB, t *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sessionToken, err := service.Login(db, r)
		if err != nil {
			w.Header().Add("HX-Redirect", "/")
			return
		}
		cookie := http.Cookie{
			Name:  "session_token",
			Value: sessionToken,
		}
		http.SetCookie(w, &cookie)
		w.Header().Add("HX-Redirect", "/")
	}
}
