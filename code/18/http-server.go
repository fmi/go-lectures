package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func init() {
	log.Println("Starting webserver")
}

func main() {
	http.HandleFunc("/world", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":8282", nil))
}
