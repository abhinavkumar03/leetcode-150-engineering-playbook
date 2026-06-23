package main

func strStr(haystack string, needle string) int {
	n := len(haystack)
	m := len(needle)

	if m == 0 {
		return 0
	}

	for start := 0; start <= n-m; start++ {
		match := true

		for offset := 0; offset < m; offset++ {
			if haystack[start+offset] != needle[offset] {
				match = false
				break
			}
		}

		if match {
			return start
		}
	}

	return -1
}