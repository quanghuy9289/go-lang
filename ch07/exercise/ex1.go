package main

import "fmt"

// the rest param is last parameter
// can take zero or more of int parameters
func sum(numbers []int) int {
	total := 0
	for _, v := range numbers {
		total += v
	}

	return total
}

func main() {
	xs := []int{1, 2, 3}
	fmt.Println(sum(xs))
}
