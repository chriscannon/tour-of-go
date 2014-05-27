package main

import (
	"fmt"
	"math/cmplx"
)

const delta = 1e-15

// Cbrt computes the cube root of a number using Newton's method.
func Cbrt(x complex128) complex128 {
	switch x {
	case 0:
		fallthrough
	case 1:
		return x
	default:
		z := x
		p := complex128(0)
		for {
			p, z = z, z-(cmplx.Pow(z, 3)-x)/(3*cmplx.Pow(z, 2))
			if cmplx.Abs(p-z) < delta {
				break
			}
		}
		return z
	}
}

func main() {
	fmt.Println(Cbrt(2))
}
