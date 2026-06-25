package twopointers

import "sort"

// threeSum returns all unique triplets such that their sum is zero.
func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	var result [][]int
	n := len(nums)

	for i := 0; i < n-2; i++ {
		// Skip duplicate fixed elements.
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, n-1

		for left < right {
			currentSum := nums[i] + nums[left] + nums[right]

			switch {
			case currentSum < 0:
				left++

			case currentSum > 0:
				right--

			default:
				result = append(result, []int{
					nums[i],
					nums[left],
					nums[right],
				})

				left++
				right--

				// Skip duplicate left values.
				for left < right && nums[left] == nums[left-1] {
					left++
				}

				// Skip duplicate right values.
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			}
		}
	}

	return result
}