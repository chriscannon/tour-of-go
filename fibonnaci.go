package main

import "fmt"

// Fibonacci is a function that returns a function that returns an int.
func Fibonacci() func() int {
	fibOne, fibTwo := 0, 1
	return func() int {
		fibOne, fibTwo = fibTwo, fibOne+fibTwo
		return fibOne
	}
}

func main() {
	f := Fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
