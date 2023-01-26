package main

import (
	"html/template"
	"net/http"
)

type Page struct {
	Title string
	Body  []byte
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	p := &Page{Title: "Test Page", Body: []byte("This is a sample Page.")}
	t, _ := template.ParseFiles("view.tmpl")
	t.Execute(w, p)
}

func main() {
	http.HandleFunc("/view", viewHandler)
	http.ListenAndServe(":8080", nil)
}
