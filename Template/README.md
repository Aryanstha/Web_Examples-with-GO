This is a simple Go package that serves a webpage with a "Test Page" title and "This is a sample Page" as the body. The webpage is served on http://localhost:8080/view

## Requirements

- Go version 1.13 or higher

## Usage

- Clone the repository
- Navigate to the root directory of the package
- Run the command go run main.go
- Open a web browser and go to http://localhost:8080/view

## File Structure

- `main.go`: The entry point of the application. It defines the routing and starts the server.
- `view.tmpl`: The template file for the webpage. It uses the html/template package to render the page using the provided data.

## Customization

- To change the `title` and `body` of the page, modify the Title and Body fields of the Page struct in the main.go file.
- To change the template file, update the file path passed to the template.ParseFiles function in the viewHandler function.

## Notes

- This package uses the `net/http` package to handle HTTP requests and responses.
- This package uses the `html/template` package to parse and execute the template files.
- This package listens on port 8080, if you want to change this, you should update the `http.ListenAndServe` function accordingly.