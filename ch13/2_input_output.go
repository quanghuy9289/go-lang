package main

import (
	"fmt"
	"os"
)

// read file
func main() {
	file, err := os.Open("1_strings.go")
	if err != nil {
		// handle error here
		return
	}
	defer file.Close() // close file as soon as the function completes

	// get file size
	stat, err := file.Stat()
	if err != nil {
		return
	}

	// read the file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return
	}

	str := string(bs)
	fmt.Println(str)
}
