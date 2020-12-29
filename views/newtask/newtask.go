package newtask

import (
	"fmt"
	"html/template"
	"net/http"
)

type contactDetail struct {
	Title string
}

// NewTask function
func NewTask(w http.ResponseWriter, req *http.Request) {

	files := []string{
		"pages/NewTask/NewTask.html",
		"pages/base.html",
	}

	tmpl := template.Must(template.ParseFiles(files...))

	if req.Method == "POST" {
		req.ParseForm()
		fmt.Println("Form Data: ", req.FormValue("title"))
		tmpl.Execute(w, struct{ Success bool }{Success: true})
	}
	
tmpl.Execute(w, nil)
}
