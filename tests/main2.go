package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func Hello(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "Hello, %q", html.EscapeString(r.URL.Path))
}

func main() {
	http.HandleFunc("/foo", Hello)
	log.Fatal(http.ListenAndServe(":8082", nil))
}
