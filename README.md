# Go Web Server Examples

This repository contains two simple examples of building an HTML web server using Go.

## Project Structure

```text
hello-web/
├── main.go
├── page.html
├── go.mod
├── README.md
└── html-template/
    ├── main.go
    └── page2.html
```

## Example 1: Serving a Static HTML File

The first example uses Go's `net/http` package and `http.ServeFile` to serve a static HTML page.

**Run:**

```bash
go run .
```

Open in your browser:

http://localhost:8081

## Example 2: Using HTML Templates

The second example uses Go's `html/template` package to render HTML templates with dynamic data.

The template receives a name from the Go program and displays it on the page.

**Run:**

```bash
go run ./html-template
```

Open in your browser:

http://localhost:8082

## Requirements

- Go
- A web browser

Both examples use Go's standard library.

## Purpose

This project was created as a practical exercise to learn the basics of web servers and HTML rendering in Go.
