package service

import (
	"database/sql"
	"net/http"
)

func Login(db *sql.DB, r *http.Request) (string, error) {
	return "sagetoken", nil
}
