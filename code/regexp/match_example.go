package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile(`Hello`)

	if re.MatchString("Hello Regular Expression.") {
		fmt.Println("Match")
	} else {
		fmt.Println("No match")
	}
}
