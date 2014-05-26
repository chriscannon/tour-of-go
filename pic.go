package main

import "code.google.com/p/go-tour/pic"

func Pic(dx, dy int) [][]uint8 {
	s := make([][]uint8, dy)
	for i := range s {
		s[i] = make([]uint8, dx)
	}

	for x, val := range s {
		for y := range val {
			val[x] = uint8((x + y) / 2)
		}
	}
	return s
}

func main() {
	pic.Show(Pic)
}
