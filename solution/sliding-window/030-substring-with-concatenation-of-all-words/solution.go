package slidingwindow

func findSubstring(s string, words []string) []int {
	if len(s) == 0 || len(words) == 0 {
		return []int{}
	}

	wordLength := len(words[0])
	wordCount := len(words)
	windowLength := wordLength * wordCount

	if len(s) < windowLength {
		return []int{}
	}

	targetFrequency := make(map[string]int)
	for _, word := range words {
		targetFrequency[word]++
	}

	result := make([]int, 0)

	for offset := 0; offset < wordLength; offset++ {
		left := offset
		currentFrequency := make(map[string]int)
		wordsInWindow := 0

		for right := offset; right+wordLength <= len(s); right += wordLength {
			currentWord := s[right : right+wordLength]

			if _, exists := targetFrequency[currentWord]; !exists {
				currentFrequency = make(map[string]int)
				wordsInWindow = 0
				left = right + wordLength
				continue
			}

			currentFrequency[currentWord]++
			wordsInWindow++

			for currentFrequency[currentWord] > targetFrequency[currentWord] {
				leftWord := s[left : left+wordLength]
				currentFrequency[leftWord]--
				wordsInWindow--
				left += wordLength
			}

			if wordsInWindow == wordCount {
				result = append(result, left)

				leftWord := s[left : left+wordLength]
				currentFrequency[leftWord]--
				wordsInWindow--
				left += wordLength
			}
		}
	}

	return result
}