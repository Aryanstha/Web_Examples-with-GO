package main

import (
	"log"
	"net/http"
)

func main() {
	// Serve the directory "static" on the root URL
	http.Handle("/", http.FileServer(http.Dir("./static")))

	// Start the server
	log.Println("Listening on :8080...")
	http.ListenAndServe(":8080", nil)
}
