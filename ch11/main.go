package main

import (
	"fmt"
	m "golang-book/ch11/math" // module alias as m
)

func main() {
	xs := []float64{1, 2, 3, 4, 5}
	avg := m.Average(xs)
	fmt.Println(avg)
	min := m.Min(xs)
	fmt.Println(min)
	max := m.Max(xs)
	fmt.Println(max)
}
