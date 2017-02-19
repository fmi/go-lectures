package main

import (
	"fmt"
	"io/ioutil"
	"net"
)

func main() {
	conn, _ := net.Dial("tcp", "golang.org:80")
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\nHost: golang.org\r\n\r\n")

	resp, _ := ioutil.ReadAll(conn)
	fmt.Printf("%s", resp)
}
