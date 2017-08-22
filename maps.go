package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

// WordCount counts every occurence of words
func WordCount(s string) map[string]int {
    res := make(map[string]int)
	words := strings.Fields(s)
	for _, word := range words {
	    elem, ok := res[word]
		if ok {
		    res[word] = elem + 1
		} else {
		    res[word] = 1
		}
	}
	return res
}

func main() {
	wc.Test(WordCount)
}
