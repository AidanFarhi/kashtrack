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

	fs := http.FileServer(http.Dir("web"))

	m := http.NewServeMux()
	m.Handle("/web/", http.StripPrefix("/web/", fs))
	m.HandleFunc("/", handler.IndexHandler(db, t))
	m.HandleFunc("POST /add_expense", handler.AddExpenseHandler(db, t))
	m.HandleFunc("POST /login", handler.LoginHandler(db, t))
	m.HandleFunc("POST /logout", handler.LogoutHandler(db, t))
	m.HandleFunc("GET /expense_distribution", handler.ExpenseDistributionHandler(db))

	server := http.Server{
		Addr:    ":8080",
		Handler: m,
	}

	server.ListenAndServe()
}
