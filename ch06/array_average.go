package main

import "fmt"

func main() {
	// var x [5]float64
	// x[0] = 100
	// x[1] = 98
	// x[2] = 70
	// x[3] = 89
	// x[4] = 92

	x := [5]float64{
		100,
		98,
		70,
		89,
		92,
	}

	var total float64 = 0
	// for i := 0; i < len(x); i++ {
	// 	total += x[i]
	// }
	// another way to write for statement
	for _, v := range x {
		total += v
	}

	fmt.Println(total / float64(len(x)))
}
