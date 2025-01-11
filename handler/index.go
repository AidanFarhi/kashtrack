package handler

import (
	"html/template"
	"net/http"
)

func IndexHandler(t *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t.ExecuteTemplate(w, "index", nil)
	}
}
