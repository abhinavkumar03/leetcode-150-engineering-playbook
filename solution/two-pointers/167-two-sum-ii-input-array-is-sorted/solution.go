package twosumii

// twoSum returns the 1-based indices of the two numbers
// whose sum equals the target.
func twoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1

	for left < right {
		currentSum := numbers[left] + numbers[right]

		switch {
		case currentSum == target:
			return []int{left + 1, right + 1}
		case currentSum < target:
			left++
		default:
			right--
		}
	}

	return []int{}
}