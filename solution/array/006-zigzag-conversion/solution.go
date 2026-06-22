package main

import "strings"

func convert(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}

	rows := make([]strings.Builder, numRows)

	currentRow := 0
	direction := 1

	for _, ch := range s {
		rows[currentRow].WriteRune(ch)

		if currentRow == 0 {
			direction = 1
		} else if currentRow == numRows-1 {
			direction = -1
		}

		currentRow += direction
	}

	var result strings.Builder

	for i := 0; i < numRows; i++ {
		result.WriteString(rows[i].String())
	}

	return result.String()
}