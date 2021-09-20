package main

import "fmt"

func main() {
	// array is fixed length in GO, cannot change length after initialization
	// => use SLICE to change length dynamically
	var x [5]int
	x[4] = 100
	fmt.Println(x)
}
