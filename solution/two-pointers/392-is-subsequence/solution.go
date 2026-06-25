package twopointers

// isSubsequence returns true if s is a subsequence of t.
func isSubsequence(s string, t string) bool {
	sIndex := 0
	tIndex := 0

	for sIndex < len(s) && tIndex < len(t) {
		if s[sIndex] == t[tIndex] {
			sIndex++
		}
		tIndex++
	}

	return sIndex == len(s)
}