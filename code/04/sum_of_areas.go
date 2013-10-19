package main

import "fmt"

type Shape interface {
	Circumference() int
}

type Rectangle struct {
	x, y int
}

func (r *Rectangle) Circumference() int {
	return r.x + r.y
}

type Triangle struct {
	x, y, z int
}

func (t *Triangle) Circumference() int {
	return t.x + t.y + t.z
}

func sumOfCircumferences(shapes ...Shape) int {
	sum := 0
	for _, shape := range shapes {
		sum += shape.Circumference()
	}
	return sum
}

func main() {
	rect := &Rectangle{x: 12, y: 64}
	tr := &Triangle{x: 12, y: 64, z: 50}
	fmt.Println(sumOfCircumferences(rect, tr))
}
