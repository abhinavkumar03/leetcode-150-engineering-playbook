package main

func canJump(nums []int) bool {
	farthestReach := 0
	lastIndex := len(nums) - 1

	for index, jumpLength := range nums {
		if index > farthestReach {
			return false
		}

		if index+jumpLength > farthestReach {
			farthestReach = index + jumpLength
		}

		if farthestReach >= lastIndex {
			return true
		}
	}

	return true
}