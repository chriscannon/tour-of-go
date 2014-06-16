package main

import (
	"fmt"

	"code.google.com/p/go-tour/tree"
)

func traverse(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		traverse(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		traverse(t.Right, ch)
	}
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	traverse(t, ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	w1 := make(chan int)
	w2 := make(chan int)

	go Walk(t1, w1)
	go Walk(t2, w2)

	for {
		val1, ok1 := <-w1
		val2, ok2 := <-w2

		if !ok1 && !ok2 {
			break
		}

		if val1 != val2 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(2), tree.New(1)))
}
