package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "<h1>This is home page.</h1><br>")
		fmt.Fprint(w, "<p>Go to <a href='/about'>About</a></p><br>")
		fmt.Fprint(w, "<p>Go to <a href='/contact'>Contact</a></p>")

	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "<h1>This is the about page.</h1><br>")
		fmt.Fprint(w, "<p>Go to <a href='/'>Home</a></p><br>")
		fmt.Fprint(w, "<p>Go to <a href='/contact'>Contact</a></p>")
	})

	http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "<h1>This is the contact page.</h1><br>")
		fmt.Fprint(w, "<p>Go to <a href='/'>Home</a></p><br>")
		fmt.Fprint(w, "<p>Go to <a href='/about'>About</a></p>")
	})

	http.ListenAndServe(":8080", nil)
}
