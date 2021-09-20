package main

import (
	"fmt"
	"math/rand"
	"time"
)

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Microsecond * amt)
	}
}

func main() {
	for i := 0; i < 10; i++ {
		go f(i) // goroutine, same as async function, f function should run simultaneously
	}
	var input string
	fmt.Scanln(&input)
}
