package main

import (
	"fmt"
	"time"
)

// A channel that doesn't have these restrictions is known as bi-directional
// A bi-directional channel can be passed to a function that takes send-only or receive-only channels
// Not reverse

// specify a direction on a channel
// => restricting it to either sending to channel
func pinger(c chan<- string) {
	for i := 0; ; i++ {
		fmt.Println(i, "pinger")
		c <- "ping" // send "ping" message to channel c
	}
}

// specify a direction on a channel
// => restricting it to either sending to channel
func ponger(c chan<- string) {
	for i := 0; ; i++ {
		fmt.Println(i, "ponger")
		c <- "pong"
	}
}

// specify a direction on a channel
// => restricting it to either receiving from channel
func printer(c <-chan string) {
	for {
		msg := <-c // receive message from channel and store it in msg
		fmt.Println("printer:", msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var c chan string = make(chan string)
	go pinger(c) // pinger and printer communicate through channel c
	go ponger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}
