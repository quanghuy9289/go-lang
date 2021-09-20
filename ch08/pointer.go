package main

import "fmt"

func square(x *float64) {
	*x = *x * *x
}

// func swap(x *float64, y *float64) {
// can collapse arguments that have the same type
func swap(x, y *float64) {
	z := *x
	*x = *y
	*y = z
}

func main() {
	x := 1.5
	y := 3.0
	square(&x)
	fmt.Println(x)

	fmt.Println("Before swap x =", x, "and y =", y)
	swap(&x, &y)
	fmt.Println("After swap x =", x, "and y =", y)
}
