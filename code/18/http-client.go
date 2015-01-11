package main

import (
	"fmt"
	"net/http" // HL
)

func main() {
	resp, err := http.Get("http://fmi.golang.bg/") // HL
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("GET fmi golang site said HTTP %d\n", resp.StatusCode)
}
