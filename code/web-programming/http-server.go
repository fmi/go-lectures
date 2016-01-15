package main

import (
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/world", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	go http.ListenAndServe(":8282", nil)

	res, err := http.Get("http://localhost:8282/world")
	if err != nil {
		log.Fatal(err)
	}
	contents, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Server responded with: %s", contents)
}
