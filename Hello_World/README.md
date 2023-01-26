# README.md

## Overview
This code package is a simple Go web server that listens on port 8080 and responds with "Hello, World!" when the root URL is accessed.

## Usage
1. Make sure you have Go installed on your machine.
2. Clone this repository to your machine.
3. Run the command `go run main.go` in the root directory of the repository.
4. Access the root URL of the server (http://localhost:8080/) in a web browser or via a GET request to see the response "Hello, World!".

## Dependencies
- `net/http` package is used for creating the web server and handling HTTP requests.
- `fmt` package is used for printing the "Hello, World!" response.

## Additional notes
The `HandleFunc` function assigns a function to handle requests to the root URL. The function takes in a `http.ResponseWriter` and `http.Request` as arguments, and writes the "Hello, World!" string to the response writer. The `ListenAndServe` function starts the web server and listens for incoming requests on port 8080. The second argument passed to this function is the handler to use when a request is received, in this case `nil` is passed because we have already defined the handler using `http.HandleFunc`
