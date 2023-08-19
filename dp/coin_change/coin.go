package coin_change

// CoinChange Brute force recursion
func CoinChange(coins []int, amount int) int {

	if amount == 0 {
		return 1
	}

	if amount < 0 || len(coins) == 0 {
		return 0
	}

	return CoinChange(coins, amount-last(coins)) + // include
		CoinChange(excludeLast(coins), amount) // exclude
}

func last(coins []int) int {
	return coins[len(coins)-1]
}

func excludeLast(coins []int) []int {
	return coins[:len(coins)-1]
}
