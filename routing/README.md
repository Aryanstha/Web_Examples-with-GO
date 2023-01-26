## Overview
This code package is a simple Go web server that listens on port 8080 and serves different responses for different routes. It has three main routes:
- "/" which serves a "home" page
- "/about" which serves an "about" page
- "/contact" which serves a "contact" page

## Usage
1. Make sure you have Go installed on your machine.
2. Clone this repository to your machine.
3. Run the command `go run main.go` in the root directory of the repository.
4. Access the URLs of the server (http://localhost:8080/) in a web browser or via a GET request to see the different pages served by the server.

## Dependencies
- `net/http` package is used for creating the web server and handling HTTP requests.
- `fmt` package is used for formatting and printing the responses to the browser

## Additional notes
The `http.HandleFunc` function assigns a function to a specific route which will be called when a request is made to that route.
In this case, three functions are assigned to three different routes ("/", "/about", "/contact").
The `http.ListenAndServe` function starts the web server and listens for incoming requests on port 8080. The second argument passed to this function is the handler to use when a request is received, in this case `nil` is passed because we have already defined the handlers using `http.HandleFunc`.
It's possible to use a middleware like `http.StripPrefix` to remove the prefix from the URL, so the URLs for the pages will look prettier.
