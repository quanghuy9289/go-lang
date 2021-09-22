package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

// receive a max argument from command line
// ex: go run 12_cmd_arguments.go -max=100 -min=10

// non-flag arguments can be retrieved with flag.Args() => []string
func main() {
	// define flags
	maxp := flag.Int("max", 100, "the max value")
	minp := flag.Int("min", 1, "the min value")

	// parse
	flag.Parse()

	// generate a number between 0 and max
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(*maxp-*minp) + *minp)
}
