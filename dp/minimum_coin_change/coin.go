package minimum_coin_change

import "math"

func MinCoinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	minChange := math.MaxInt32

	for _, coin := range coins {
		if amount >= coin {
			change := MinCoinChange(coins, amount-coin) + 1
			minChange = min(minChange, change)
		}
	}
	return minChange
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
