package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	ListenAddress string
)

func init() {
	flag.StringVar(&ListenAddress, "addr", ":8080", "Listening address and port.")
}

func main() {
	flag.Parse()

	directory := "."
	if len(flag.Args()) >= 1 {
		directory = flag.Arg(0)
	}

	serverHandler := http.FileServer(http.Dir(directory))
	log.Fatal(http.ListenAndServe(ListenAddress, serverHandler))
}
