package slidingwindow

func minSubArrayLen(target int, nums []int) int {
	left := 0
	windowSum := 0
	minLength := len(nums) + 1

	for right := 0; right < len(nums); right++ {
		windowSum += nums[right]

		for windowSum >= target {
			currentLength := right - left + 1
			if currentLength < minLength {
				minLength = currentLength
			}

			windowSum -= nums[left]
			left++
		}
	}

	if minLength == len(nums)+1 {
		return 0
	}

	return minLength
}