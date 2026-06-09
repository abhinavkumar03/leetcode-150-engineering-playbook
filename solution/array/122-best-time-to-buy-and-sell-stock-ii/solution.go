package main

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	totalProfit := 0

	for day := 1; day < len(prices); day++ {
		priceDifference := prices[day] - prices[day-1]

		if priceDifference > 0 {
			totalProfit += priceDifference
		}
	}

	return totalProfit
}