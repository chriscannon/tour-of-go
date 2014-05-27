package main

import (
	"strings"

	"code.google.com/p/go-tour/wc"
)

// WordCount is a function that takes a string and returns the number of occurrences of each word.
func WordCount(s string) map[string]int {
	wordMap := make(map[string]int)
	for _, word := range strings.Fields(s) {
		wordMap[word]++
	}
	return wordMap
}

func main() {
	wc.Test(WordCount)
}
