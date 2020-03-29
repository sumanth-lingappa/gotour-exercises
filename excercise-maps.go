package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	retDict := make(map[string]int)
	for _, v := range strings.Fields(s) {
		if count, ok := retDict[v]; ok {
			retDict[v] = count + 1
		} else {
			retDict[v] = 1
		}
	}
	return retDict
}

func main() {
	wc.Test(WordCount)
}