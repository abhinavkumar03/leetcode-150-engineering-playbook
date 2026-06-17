package main

func romanToInt(s string) int {
	values := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	result := 0

	for i := 0; i < len(s); i++ {
		currentValue := values[s[i]]

		if i < len(s)-1 && currentValue < values[s[i+1]] {
			result -= currentValue
		} else {
			result += currentValue
		}
	}

	return result
}