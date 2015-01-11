package main

import (
	"fmt"
	"github.com/codegangsta/negroni"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	// Use the default middleware.
	n := negroni.Classic()
	// ... Add any other middlware here

	// add the router as the last handler in the stack
	n.UseHandler(mux)
	n.Run(":3000")
}
