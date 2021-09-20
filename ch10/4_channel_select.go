package main

import (
	"fmt"
	"time"
)

// Select works like a switch but for channel
func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "from 1"
			time.Sleep(time.Second * 2)
		}
	}()

	go func() {
		for {
			c2 <- "from 2"
			time.Sleep(time.Second * 3)
		}
	}()

	go func() {
		for {
			// select picks the first channel that is ready and receives from it (or sends to it)
			// if more than one of the channels are ready then it randomly picks which one to receive from
			// if none of the channels are ready, statement blocks until one becomes available
			select {
			case msg1 := <-c1:
				fmt.Println("Message 1", msg1)
			case msg2 := <-c2:
				fmt.Println("Message 2", msg2)
			case timer := <-time.After(time.Second):
				fmt.Println("Time:", timer.String())
				fmt.Println("timeout")
				// default:
				// 	fmt.Println("nothing ready")
			}

		}
	}()

	var input string
	fmt.Scanln(&input)
}
