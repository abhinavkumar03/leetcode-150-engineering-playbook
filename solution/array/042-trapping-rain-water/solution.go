package main

func trap(height []int) int {
	n := len(height)
	if n < 3 {
		return 0
	}

	left, right := 0, n-1
	leftMax, rightMax := 0, 0
	trappedWater := 0

	for left < right {
		if height[left] > leftMax {
			leftMax = height[left]
		}

		if height[right] > rightMax {
			rightMax = height[right]
		}

		if leftMax < rightMax {
			trappedWater += leftMax - height[left]
			left++
		} else {
			trappedWater += rightMax - height[right]
			right--
		}
	}

	return trappedWater
}