package main

import (
	"io"
	"net/http"
)

func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	io.WriteString(
		res,
		`<doctype html>
		<html>
			<head>
				<title>Hello world</title>
			</head>
			<body>
				Hello World!
			</body>
		</html>`,
	)
}

func main() {
	// handle get request
	http.HandleFunc("/hello", hello)

	// static file
	http.Handle(
		"/assets/",
		http.StripPrefix(
			"/assets/",                          // remove /assets/ in request path, remain file name
			http.FileServer(http.Dir("assets")), // serve asset at assets folder in server
		),
	)

	// listen
	http.ListenAndServe(":9000", nil)
}
