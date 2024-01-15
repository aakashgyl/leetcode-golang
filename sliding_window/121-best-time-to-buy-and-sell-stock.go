package main

func maxProfit(prices []int) int {
	i := 0
	j := 1
	max := 0

	for j < len(prices) {
		if prices[i] > prices[j] { //
			i = j
			j++
		} else {
			profit := prices[j] - prices[i]
			max = getMax(max, profit)
			j++
		}
	}

	return max
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
