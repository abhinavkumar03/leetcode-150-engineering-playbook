package main

import (
	"strings"
)

func reverseWords(s string) string {
	words := strings.Fields(s)

	for i := 0; i < len(words)/2; i++ {
		j := len(words) - 1 - i
		words[i], words[j] = words[j], words[i]
	}

	return strings.Join(words, " ")
}