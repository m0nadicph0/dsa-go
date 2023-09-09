package backtrack

func UniquePaths(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}

	// Move either right or down, and sum the results.
	return UniquePaths(m-1, n) + UniquePaths(m, n-1)
}
