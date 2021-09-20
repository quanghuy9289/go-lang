package main

import "fmt"

func main() {
	// declare map
	var x map[string]int // x is a map with key type is string and value type is int
	//x["key"] = 10        // runtime error => map has to be initialized before they can be used
	fmt.Println(x)

	y := make(map[string]int) // initialize map
	y["key"] = 10             // ok
	fmt.Println(y["key"])

}
