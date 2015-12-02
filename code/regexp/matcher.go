package main

import (
	"fmt"
	"regexp"
)

// second_start OMIT
func matcher(pattern, text string) {
	match := regexp.MustCompile(pattern).FindStringIndex(text)
	if match == nil {
		fmt.Printf("No match for `%s` in \"%s\"\n", pattern, text)
		return
	}

	begin := match[0]
	end := match[1]
	fmt.Printf("%s(%s)%s\n", text[:begin], text[begin:end], text[end:])
}

func main() {
	// first_start OMIT
	matcher(`pat`, "Find a pattern.") // Find a (pat)tern.
	matcher(`#`, "What the ###?")     // What the (#)##?
	// first_end OMIT
}
