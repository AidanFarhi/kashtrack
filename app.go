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

	tmpl := template.Must(template.ParseGlob("web/templates/*.html"))

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.IndexHandler(db, tmpl))

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	server.ListenAndServe()
}
