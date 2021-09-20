package main

import (
	"fmt"
	"time"
)

// Channel provide a way for 2 goroutines to communicate with one other
// and synchronize their execution

// when pinger attempts to send a message on the channel it will wait until printer is ready to receive the message
// => blocking
func pinger(c chan string) {
	for i := 0; ; i++ {
		fmt.Println(i, "pinger")
		c <- "ping" // send "ping" message to channel c
	}
}
func ponger(c chan string) {
	for i := 0; ; i++ {
		fmt.Println(i, "ponger")
		c <- "pong"
	}
}
func printer(c chan string) {
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
