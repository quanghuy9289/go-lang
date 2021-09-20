package main

import "fmt"

// the rest param is last parameter
// can take zero or more of int parameters
func total(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}

	return total
}

func main() {
	fmt.Println(total())
	fmt.Println(total(1, 2, 3))

	xs := []int{1, 2, 3}
	fmt.Println(total(xs...))
}
