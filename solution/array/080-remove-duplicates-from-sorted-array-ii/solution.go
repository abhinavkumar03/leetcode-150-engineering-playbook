package main

func removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}

	write := 2

	for read := 2; read < len(nums); read++ {
		if nums[read] != nums[write-2] {
			nums[write] = nums[read]
			write++
		}
	}

	return write
}