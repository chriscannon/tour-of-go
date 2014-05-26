package main

import (
	"strings"

	"code.google.com/p/go-tour/wc"
)

func WordCount(s string) map[string]int {
	wordMap := make(map[string]int)
	words := strings.Fields(s)
	if len(words) > 0 {
		for _, word := range words {
			val, present := wordMap[word]
			if present {
				wordMap[word] = val + 1
			} else {
				wordMap[word] = 1
			}
		}
	}

	return wordMap
}

func main() {
	wc.Test(WordCount)
}
