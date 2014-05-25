package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) (float64, int) {
	z := 1.0
	p := 0.0
	iteration := 0
	for {
		z = z - (z*z-x)/(2*z)
		iteration++
		if math.Abs(p-z) < 1e-15 {
			break
		}
		p = z
	}
	return z, iteration
}

func main() {
	val := 10.0
	z, iteration := Sqrt(val)
	z_math := math.Sqrt(val)
	fmt.Printf("%d iterations square root: %v\n", iteration, z)
	fmt.Printf("math package square root: %v\n", math.Sqrt(val))
	fmt.Printf("diff: %v\n", math.Abs(z-z_math))
}
