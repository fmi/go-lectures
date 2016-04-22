package main

import "fmt"

type mystruct struct{ a, b int }

func (m mystruct) String() string {
	return fmt.Sprintf("mystruct(%d, %d) yo!", m.a, m.b)
}

func stringify(value interface{}) string {
	switch s := value.(type) {
	case string:
		return s
	case fmt.Stringer:
		return s.String()
	case int, uint:
		return fmt.Sprintf("%06d", s) // leftpad
	default:
		return "No clue, mate :)"
	}
}

func main() {
	fmt.Printf("%#v\n", stringify(mystruct{1, 2}))
}
