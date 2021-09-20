package main

import "fmt"

func main() {
	// get slice from array
	x := [6]string{"a", "b", "c", "d", "e", "f"}
	fmt.Println(x[2:5]) // => ["c", "d", "e"] (exclude the last index)
}
