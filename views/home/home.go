package home

import (
	"html/template"
	"net/http"
)

// Index function
func Index(w http.ResponseWriter, req *http.Request) {

	files := []string{
		"pages/home/home.html",
		"pages/base.html",
	}

	tmpl := template.Must(template.ParseFiles(files...))
	tmpl.Execute(w, nil)
}
