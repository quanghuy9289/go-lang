package main

import "fmt"

func main() {
	/*	if we write as below, the recorver will never happen,
		because the call to panic immediately stops execution of the function

		func main() {
			panic("PANIC")   // => stop execution
			str := recover()
			fmt.Println(str)
		}
	*/

	// using defer to make sure recover alway be called when panic happen
	defer func() {
		str := recover() // collect panic information to report and avoid stop execution
		fmt.Println(str)
	}()

	panic("PANIC")
}
