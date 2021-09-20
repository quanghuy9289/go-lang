package main

import "fmt"

func main() {
	// create a slice of float64 type with length is 5
	// can add one more arg to set capacity of slice
	// make([]float64, 5, 10) => length can be increase to 10 (10 is capacity of array that slice underlying)
	x := make([]float64, 5)
	fmt.Println(x)

	// create slice from array
	arr := [5]float64{5.4, 6, 8.5, 9, 10}
	// from index to end index - exclude the element at end index
	// from index and end index can be ignore
	// y := arr[0:], y := arr[:5], y := arr[:] => the same
	y := arr[0:5]
	fmt.Println(y)

	// slice function
	// append
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)

	// copy
	slice3 := []int{1, 2, 3}
	slice4 := make([]int, 2)
	copy(slice4, slice3) // slice4 has length is 2 so just the first 2 elements of slice3 are copied into slice4
	fmt.Println(slice3, slice4)
}
