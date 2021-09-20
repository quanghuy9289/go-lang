package main

import "fmt"

// GO just have ONE loop type is FOR

func main() {

	// one way to define loop
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i += 1
	}

	// another way
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}
