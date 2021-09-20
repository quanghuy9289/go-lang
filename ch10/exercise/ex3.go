package main

import (
	"fmt"
	"time"
)

func pinger(c chan int) {
	for i := 0; ; i++ {
		fmt.Println(i, "pinger")
		c <- i // send "ping" message to channel c
	}
}

func printer(c chan int) {
	for {
		msg := <-c // receive message from channel and store it in msg
		fmt.Println("printer:", msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	c := make(chan int, 20)
	go pinger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}
