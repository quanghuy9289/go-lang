package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

func handleServerConnection(c net.Conn) {
	// receive the message
	var msg string
	err := gob.NewDecoder(c).Decode(&msg)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Received:", msg)
	}
	c.Close()
}

func server() {
	// listen on a port
	listener, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		// accept a connection
		connection, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		// handle the connection
		go handleServerConnection(connection)
	}
}

func client() {
	// connect to the server
	c, err := net.Dial("tcp", "localhost:9999")
	if err != nil {
		fmt.Println(err)
		return
	}

	// send the message
	msg := "Hello world!"
	fmt.Println("Sending", msg)
	err = gob.NewEncoder(c).Encode(msg)
	if err != nil {
		fmt.Println(err)
	}

	c.Close()
}

func main() {
	go server()
	go client()

	var input string
	fmt.Scanln(&input)
}
