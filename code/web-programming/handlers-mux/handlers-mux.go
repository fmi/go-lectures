package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/doc/", http.FileServer(http.Dir("/usr/share/")))
	mux.Handle("/google/", http.RedirectHandler("https://www.google.com", http.StatusTemporaryRedirect))
	mux.Handle("/proxy/", http.StripPrefix("/proxy/", httputil.NewSingleHostReverseProxy(
		&url.URL{Scheme: "https", Host: "mirrors.kernel.org", Path: "/"},
	)))

	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		if req.URL.Path == "/" {
			fmt.Fprintf(w, "Welcome to the jungle!")
			return
		}
		http.NotFound(w, req)
	})

	s := &http.Server{Addr: ":8282", Handler: mux, WriteTimeout: 1 * time.Second}
	log.Printf("Starting server on %s", s.Addr)
	log.Fatal(s.ListenAndServe())
}
