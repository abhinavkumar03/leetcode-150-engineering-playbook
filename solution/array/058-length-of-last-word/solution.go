package main

func lengthOfLastWord(s string) int {
	n := len(s)
	index := n - 1

	// Skip trailing spaces
	for index >= 0 && s[index] == ' ' {
		index--
	}

	length := 0

	// Count characters of the last word
	for index >= 0 && s[index] != ' ' {
		length++
		index--
	}

	return length
}