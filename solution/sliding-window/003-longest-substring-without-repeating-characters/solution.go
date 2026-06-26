package slidingwindow

func lengthOfLongestSubstring(s string) int {
	left := 0
	maxLength := 0
	window := make(map[byte]bool)

	for right := 0; right < len(s); right++ {
		for window[s[right]] {
			window[s[left]] = false
			left++
		}

		window[s[right]] = true

		currentLength := right - left + 1
		if currentLength > maxLength {
			maxLength = currentLength
		}
	}

	return maxLength
}