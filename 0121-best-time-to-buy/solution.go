package _121_best_time_to_buy

import "math"

func maxProfit(prices []int) int {
	maxProfit := 0
	minPrice := math.MaxInt

	for i := 0; i < len(prices); i++ {
		minPrice = min(minPrice, prices[i])
		maxProfit = max(maxProfit, prices[i]-minPrice)
	}

	return maxProfit
}
