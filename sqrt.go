package main

import (
	"fmt"
	"math"
)

const delta = 1e-15

// Sqrt is a function that returns the square root of a float64 using Netwon's method.
func Sqrt(x float64) (float64, int) {
	z := x
	p := 0.0
	iteration := 1
	for ; ; iteration++ {
		p, z = z, z-(z*z-x)/(2*z)
		if math.Abs(p-z) < delta {
			break
		}
	}
	return z, iteration
}

func main() {
	const val = 10.0
	myFunc, iteration := Sqrt(val)
	goFunc := math.Sqrt(val)
	fmt.Printf("%d iterations square root: %v\n", iteration, myFunc)
	fmt.Printf("math package square root: %v\n", goFunc)
	fmt.Printf("diff: %v\n", math.Abs(myFunc-goFunc))
}
