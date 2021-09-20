package main

import "fmt"

func main() {
	// length of slice
	x := make([]int, 3, 9)
	fmt.Println(len(x)) // => 3
}
