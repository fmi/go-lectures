package main

import (
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	go http.ListenAndServe(":8282", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %s", html.EscapeString(r.URL.Path))
	}))

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
