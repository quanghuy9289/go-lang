package main

import "fmt"

func half(number int) (int, bool) {
	if number%2 == 0 {
		return number / 2, true
	} else {
		return number / 2, false
	}
}

func main() {
	fmt.Println(half(1))
	fmt.Println(half(2))
	fmt.Println(half(3))
}
