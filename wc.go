package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

// WordCount ...
func WordCount(s string) map[string]int {
	var words []string = strings.Fields(s)
	res := make(map[string]int)
	for _, v := range words {
		res[v]++
	}
	return res
}

func main() {
	wc.Test(WordCount)
}
