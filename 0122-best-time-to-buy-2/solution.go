package _122_best_time_to_buy_2

func maxProfit(prices []int) int {
	controlPrice, profit := prices[0], 0
	for i := 0; i < len(prices); i++ {
		if prices[i] > controlPrice {
			profit += prices[i] - controlPrice
			controlPrice = prices[i]
		}

		controlPrice = min(controlPrice, prices[i])
	}

	return profit
}
