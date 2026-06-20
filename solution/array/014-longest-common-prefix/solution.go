package main

import "strings"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]

	for _, current := range strs[1:] {
		for !strings.HasPrefix(current, prefix) {
			prefix = prefix[:len(prefix)-1]

			if prefix == "" {
				return ""
			}
		}
	}

	return prefix
}