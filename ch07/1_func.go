package main

import "fmt"

// function take in a slice of float64 and return one float64
func average(slice []float64) float64 {
	var total float64 = 0.0
	for _, v := range slice {
		total += v
	}

	return total / float64(len(slice))
}

// naming for return type
// r will be saved returned value
func f2() (r int) {
	r = 1
	return
}

func main() {
	x := []float64{
		100,
		98,
		70,
		89,
		92,
	}
	fmt.Println(average(x))
	fmt.Println(f2())
}
