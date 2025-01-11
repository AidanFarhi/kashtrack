package main

import (
	"html/template"
	"kashtrack/handler"
	"net/http"
)

func main() {

	tmpl := template.Must(template.ParseGlob("web/html/*.html"))

	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.IndexHandler(tmpl))

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	server.ListenAndServe()
}
