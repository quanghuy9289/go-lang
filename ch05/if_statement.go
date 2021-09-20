package main

import "fmt"

// GO just have ONE loop type is FOR

func main() {
	// odd-even from 1 - 10
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Printf("%d - event\n", i)
		} else {
			fmt.Printf("%d - odd\n", i)
		}
	}

	// 2. divisible by 3 between 1 and 100
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}

	// 3. divisible by 3 and 5 between 1 and 100
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
