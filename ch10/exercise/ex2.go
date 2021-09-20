package main

import (
	"fmt"
	"time"
)

func Sleep(milisecond int) <-chan time.Time {
	return time.After(time.Duration(milisecond) * time.Millisecond)
}

func main() {
	for i := 0; i < 10; i++ {
		select {
		case <-Sleep(1000):
			fmt.Println(i)
		}
	}
}
