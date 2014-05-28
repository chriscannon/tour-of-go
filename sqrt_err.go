package main

import (
	"fmt"
	"math"
	"os"
)

const delta = 1e-15

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Cannot compute the square root of a negative number: %v", float64(e))
}

// Sqrt is a function that returns the square root of a float64 using Netwon's method.
func Sqrt(x float64) (float64, int, error) {
	if x < 0 {
		return 0, 0, ErrNegativeSqrt(x)
	}

	z := x
	p := 0.0
	iteration := 1
	for ; ; iteration++ {
		p, z = z, z-(z*z-x)/(2*z)
		if math.Abs(p-z) < delta {
			break
		}
	}
	return z, iteration, nil
}

func main() {
	const val = -2.0
	myFunc, iteration, err := Sqrt(val)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	goFunc := math.Sqrt(val)
	fmt.Printf("%d iterations square root: %v\n", iteration, myFunc)
	fmt.Printf("math package square root: %v\n", goFunc)
	fmt.Printf("diff: %v\n", math.Abs(myFunc-goFunc))
}
