package main

import "fmt"

func largestNumber(numbers ...int) int {
	largestNumber := 0
	for _, v := range numbers {
		if largestNumber < v {
			largestNumber = v
		}
	}
	return largestNumber
}

func main() {
	fmt.Println(largestNumber(5, 9, 16, 2, 1))
}
