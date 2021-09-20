package main

import (
	"fmt"
)

type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

// struct's fields usually represent the has-a relationship
// Android has-a Person => don't like that
// => embedded type, is-a relationship: Android is a person
type Android struct {
	Person // just define type Person without name
	Model  string
}

func main() {
	a := new(Android)
	a.Talk() // can call Person method directly on the Android instance
}
