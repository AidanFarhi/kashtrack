package handler

import (
	"database/sql"
	"html/template"
	"kashtrack/service"
	"net/http"
	"time"
)

func LoginHandler(db *sql.DB, t *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sessionToken, err := service.Login(db, r)
		if err != nil {
			w.Header().Add("HX-Redirect", "/")
			return
		}
		cookie := http.Cookie{
			Name:     "session_token",
			Value:    sessionToken,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)
		w.Header().Add("HX-Redirect", "/")
	}
}

func LogoutHandler(db *sql.DB, t *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId, _ := service.GetUserIDFromSessionToken(db, r)
		service.Logout(db, userId)
		http.SetCookie(w, &http.Cookie{
			Name:     "session_token",
			Value:    "",
			Path:     "/",
			Expires:  time.Unix(0, 0),
			MaxAge:   -1,
			HttpOnly: true,
		})
		w.Header().Add("HX-Redirect", "/")
	}
}
