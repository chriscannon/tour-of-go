package main

import (
	"fmt"
	"math"
)

const delta = 1e-15

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
	z, iteration := Sqrt(val)
	z_math := math.Sqrt(val)
	fmt.Printf("%d iterations square root: %v\n", iteration, z)
	fmt.Printf("math package square root: %v\n", z_math)
	fmt.Printf("diff: %v\n", math.Abs(z-z_math))
}
