package main

func jump(nums []int) int {
	n := len(nums)

	if n <= 1 {
		return 0
	}

	jumps := 0
	currentEnd := 0
	farthestReach := 0

	for i := 0; i < n-1; i++ {
		if i+nums[i] > farthestReach {
			farthestReach = i + nums[i]
		}

		if i == currentEnd {
			jumps++
			currentEnd = farthestReach

			if currentEnd >= n-1 {
				break
			}
		}
	}

	return jumps
}