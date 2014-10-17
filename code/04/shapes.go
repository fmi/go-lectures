package main

import (
	"fmt"
	"math"
)

// start interface OMIT
type Shape interface {
	Area() float64
	Circumference() int
}

// end interface OMIT

// start types OMIT
type Rectangle struct {
	X, Y int
}

type Triangle struct {
	X, Y, Z int
}

// end types OMIT

// start methods OMIT
func (r *Rectangle) Area() float64 {
	return float64(r.X * r.Y)
}

func (r *Rectangle) Circumference() int {
	return 2 * (r.X + r.Y)
}

func (t *Triangle) Area() float64 {
	p := float64(t.Circumference() / 2)
	return math.Sqrt(p * (p - float64(t.X)) * (p - float64(t.Y)) * (p - float64(t.Z)))
}

func (t *Triangle) Circumference() int {
	return t.X + t.Y + t.Z
}

// end methods OMIT

// start funcs OMIT
func sumOfCircumferences(shapes ...Shape) int {
	sum := 0
	for _, shape := range shapes {
		sum += shape.Circumference()
	}
	return sum
}

func biggestArea(shapes ...Shape) (result float64) {
	for _, shape := range shapes {
		area := shape.Area()
		if area > result {
			result = area
		}
	}
	return result
}

func main() {
	rect := &Rectangle{X: 12, Y: 64}
	tr := &Triangle{X: 12, Y: 64, Z: 50}
	fmt.Println(sumOfCircumferences(rect, tr))
	fmt.Println(biggestArea(rect, tr))
}

// end funcs OMIT
