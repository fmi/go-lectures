package optimizations

type Cursor struct {
	X, Y int
}

const width, height = 640, 480

func Center(c *Cursor) {
	c.X += width / 2
	c.Y += height / 2
}

func CenterCursor() (int, int) {
	c := new(Cursor)
	Center(c)
	return c.X, c.Y
}

// BEGIN Sum OMIT
func SumFirst100() int {
	numbers := make([]int, 100)
	for i := range numbers {
		numbers[i] = i + 1
	}
	var sum int
	for _, i := range numbers {
		sum += i
	}
	return sum
}

// END Sum OMIT
