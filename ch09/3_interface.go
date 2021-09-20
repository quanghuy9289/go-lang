package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

// ===== Circle =====
type Circle struct {
	x, y, r float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

// ===== end Circle =====

// ===== Rectangle =====
type Rectangle struct {
	x1, y1, x2, y2 float64
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

// area is a method in GO
// (r *Rectangle) - receiver - like a parameter in function
// create function in this way allows us to call the function using the . operator
//   r.area() => easier to read, no longer need the & operator
// we use pointer here for receiver r - for effecting changes if any
func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

// ===== end Rectangle =====

func totalArea(shapes ...Shape) float64 {
	var area float64 = 0.0
	for _, v := range shapes {
		area += v.area()
	}
	return area
}

func main() {
	r := Rectangle{0, 0, 10, 10}
	c := Circle{0, 0, 15}
	fmt.Println(totalArea(&c, &r)) // no longer need the & operator, Go automatically knows to pass a pointer
}
