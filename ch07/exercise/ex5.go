package main

import "fmt"

func fibonaci(n int) int {
	if n <= 1 {
		return n
	}

	return fibonaci(n-1) + fibonaci(n-2)
}
func main() {
	fmt.Println(fibonaci(5))
}
