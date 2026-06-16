func candy(ratings []int) int {
	n := len(ratings)
	if n == 0 {
		return 0
	}

	candies := make([]int, n)

	for i := range candies {
		candies[i] = 1
	}

	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}

	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			requiredCandies := candies[i+1] + 1
			if requiredCandies > candies[i] {
				candies[i] = requiredCandies
			}
		}
	}

	totalCandies := 0

	for _, candyCount := range candies {
		totalCandies += candyCount
	}

	return totalCandies
}