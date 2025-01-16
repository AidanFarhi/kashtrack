package service

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"net/http"
)

func generateToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}

func Login(db *sql.DB, r *http.Request) (string, error) {
	return "sagetoken", nil
}
