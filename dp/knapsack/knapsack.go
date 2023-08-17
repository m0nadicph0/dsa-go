package knapsack

func Knapsack(weights []int, values []int, capacity int) int {
	// Base case: no items left or knapsack capacity is 0
	if len(weights) == 0 || capacity == 0 {
		return 0
	}

	// If the weight of the last item exceeds the knapsack capacity,
	// skip this item and recursively solve for the remaining items
	if last(weights) > capacity {
		return Knapsack(skipLast(weights), skipLast(values), capacity)
	}

	// Calculate the value when the last item is included
	include := last(values) + Knapsack(
		skipLast(weights),
		skipLast(values),
		capacity-last(weights),
	)

	// Calculate the value when the last item is excluded
	exclude := Knapsack(
		skipLast(weights),
		skipLast(values),
		capacity,
	)

	// Return the maximum value between
	// including and excluding the last item
	return max(include, exclude)
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func skipLast(weights []int) []int {
	return weights[:(len(weights) - 1)]
}

func last(a []int) int {
	return a[len(a)-1]
}
