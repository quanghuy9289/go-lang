package main

import "fmt"

func main() {
	// access the 4nd element of array or slice
	x := [5]int{1, 2, 3, 4, 5} // array
	y := []int{9, 8, 7, 6, 6}  // slice
	fmt.Println(x[4], y[4])
}
