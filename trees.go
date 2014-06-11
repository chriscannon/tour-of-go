package main

import (
	"fmt"

	"code.google.com/p/go-tour/tree"
)

func traverse(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		traverse(t, ch)
	}
	fmt.Println(t.Value)
	if t.Right != nil {
		traverse(t, ch)
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
	return true
}

func main() {
	ch := make(chan int)
	Walk(tree.New(1), ch)
}
