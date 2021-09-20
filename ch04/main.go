package main

import "fmt"

// GO is scope block

// global scope
var globalVar string = "Hello World!"

// constant
const PI = 3.14

func main() {
	// variable
	// specify variable type
	var x string = "Hello World!"
	fmt.Println(x)

	// let Go compiler to infer the type based on the literal value
	var y = "Hello World!"
	z := "Hello World!" // shorter statement => use this form whenever possible
	fmt.Println(y)
	fmt.Println(z)

	fmt.Println(PI)

}

func f() {
	fmt.Println(globalVar)
}
