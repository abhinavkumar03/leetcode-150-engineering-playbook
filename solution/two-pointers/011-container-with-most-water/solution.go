package twopointers

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxArea := 0

	for left < right {
		width := right - left

		var currentHeight int
		if height[left] < height[right] {
			currentHeight = height[left]
		} else {
			currentHeight = height[right]
		}

		currentArea := currentHeight * width
		if currentArea > maxArea {
			maxArea = currentArea
		}

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxArea
}