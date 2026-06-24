package main

import "strings"

func fullJustify(words []string, maxWidth int) []string {
	var result []string
	n := len(words)
	index := 0

	for index < n {
		lineStart := index
		lineLength := len(words[index])
		index++

		for index < n &&
			lineLength+1+len(words[index]) <= maxWidth {
			lineLength += 1 + len(words[index])
			index++
		}

		lineWords := words[lineStart:index]
		isLastLine := index == n

		result = append(
			result,
			buildLine(lineWords, maxWidth, isLastLine),
		)
	}

	return result
}

func buildLine(words []string, maxWidth int, isLastLine bool) string {
	wordCount := len(words)

	totalWordLength := 0
	for _, word := range words {
		totalWordLength += len(word)
	}

	if isLastLine || wordCount == 1 {
		line := strings.Join(words, " ")
		line += strings.Repeat(" ", maxWidth-len(line))
		return line
	}

	totalSpaces := maxWidth - totalWordLength
	gaps := wordCount - 1

	baseSpaces := totalSpaces / gaps
	extraSpaces := totalSpaces % gaps

	var builder strings.Builder

	for i := 0; i < wordCount; i++ {
		builder.WriteString(words[i])

		if i == wordCount-1 {
			continue
		}

		spaces := baseSpaces

		if i < extraSpaces {
			spaces++
		}

		builder.WriteString(strings.Repeat(" ", spaces))
	}

	return builder.String()
}