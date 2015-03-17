package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	var a = strings.Fields(s)
	var b = make(map[string]int)
	for _, v := range a {
		b[v]++
	}
	return b
}

func main() {
	wc.Test(WordCount)
}
