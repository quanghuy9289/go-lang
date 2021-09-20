package shape

import (
	"math"
)

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
