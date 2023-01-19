package main

import "fmt"

func first() {
	fmt.Println("1st")
}

func second() {
	fmt.Println("2nd")
}

// defer is often used when resources need to be freed
/* Ex:
	f, _ := os.Open(filename)
	defer f.Close()

Has 3 advantages:
	1 - It keeps our Close call near our Open call so its easier to understand
	2 - If our function had multiple return statements, Close will happen before both of them
	3 - Deferred functions are run event if a run-time panic occurs
*/
func main() {
	// defer moves the call to second to the end of the function
	defer second()
	first()
}
