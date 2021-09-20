package main

import "fmt"

func main() {
	// int and float64
	fmt.Println("1 + 1 = ", 1+1)
	fmt.Println("1 + 1 = ", 1.0+1.0)
	fmt.Println(" 32132 × 42452 = ", 32132*42452)

	// string
	fmt.Println("Hello " + "World!")
	fmt.Println(len("Hello World!"))
	fmt.Println("Hello World!"[1])

	// boolean
	fmt.Println(true)
}
