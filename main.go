package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	mapResult := map[string]int{}
	words := strings.Fields(s)
	for _, word := range words {
		mapResult[word]++
	}
	return mapResult
}

func main() {
	wc.Test(WordCount)
}
