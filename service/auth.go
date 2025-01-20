package service

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"errors"
	"net/http"
	"time"
)

func generateToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}

func validateLogin(db *sql.DB, r *http.Request) (int, error) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	var userId int
	err := db.QueryRow(`SELECT id FROM user WHERE username = ? AND password = ?`, username, password).Scan(&userId)
	if err != nil {
		return 0, errors.New("error validating password")
	}
	return userId, nil
}

func createNewSession(db *sql.DB, userId int) string {
	token := generateToken()
	db.Exec(`DELETE FROM session WHERE user_id = ?`, userId)
	db.Exec(`INSERT INTO session (token, user_id, created_at) VALUES (?, ?, ?)`, token, userId, time.Now())
	return token
}

func ValidateSession(db *sql.DB, r *http.Request) error {
	var userID string
	cookie, err := r.Cookie("session_token")
	if err != nil {
		return err
	}
	token := cookie.Value
	err = db.QueryRow("SELECT user_id FROM session WHERE token = ?", token).Scan(&userID)
	if err != nil {
		return err
	}
	return nil
}

func Login(db *sql.DB, r *http.Request) (string, error) {
	userId, err := validateLogin(db, r)
	if err != nil {
		return "", err
	}
	sessionToken := createNewSession(db, userId)
	return sessionToken, nil
}
