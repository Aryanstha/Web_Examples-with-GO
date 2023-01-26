## Overview
This code package is a simple Go web server that listens on port 8080 and serves the contents of the "./static" directory as static files.

## Usage
1. Make sure you have Go installed on your machine.
2. Clone this repository to your machine.
3. Run the command `go run main.go` in the root directory of the repository.
4. Access the root URL of the server (http://localhost:8080/) in a web browser or via a GET request to see the static files served by the server.

## Dependencies
- `net/http` package is used for creating the web server and handling HTTP requests.
- `log` package is used for printing log messages.

## Additional notes
The `http.Handle` function assigns the root URL to the `http.FileServer` handler, which serves files from the specified directory. The `http.ListenAndServe` function starts the web server and listens for incoming requests on port 8080. The second argument passed to this function is the handler to use when a request is received, in this case `nil` is passed because we have already defined the handler using `http.Handle`.
The log.Println("Listening on :8080...") will print a message to the console that the server is listening. 
