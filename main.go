package main

import (
	"net/http"

	"github.com/sahdev0/golang-todo/views/home"
	"github.com/sahdev0/golang-todo/views/newtask"
)

func main() {

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// routes
	http.HandleFunc("/", home.Index)
	http.HandleFunc("/add-task", newtask.NewTask)
	// end

	http.ListenAndServe(":8081", nil)
}
