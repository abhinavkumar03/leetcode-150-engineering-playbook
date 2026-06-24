package main

import "unicode"

func isPalindrome(s string) bool {
	left := 0
	right := len(s) - 1

	for left < right {
		for left < right && !isAlphaNumeric(rune(s[left])) {
			left++
		}

		for left < right && !isAlphaNumeric(rune(s[right])) {
			right--
		}

		if unicode.ToLower(rune(s[left])) != unicode.ToLower(rune(s[right])) {
			return false
		}

		left++
		right--
	}

	return true
}

func isAlphaNumeric(char rune) bool {
	return unicode.IsLetter(char) || unicode.IsDigit(char)
}