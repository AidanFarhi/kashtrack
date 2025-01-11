package main

import (
	"database/sql"
	"html/template"
	"kashtrack/handler"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, _ := sql.Open("sqlite3", "db/expense.db")

	t := template.Must(template.ParseGlob("web/templates/*.html"))

	m := http.NewServeMux()
	m.HandleFunc("/", handler.IndexHandler(db, t))
	m.HandleFunc("POST /add_expense", handler.AddExpenseHandler(db, t))

	server := http.Server{
		Addr:    ":8080",
		Handler: m,
	}

	server.ListenAndServe()
}
