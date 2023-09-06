package backtrack

func SubsetSum(input []int, target int) bool {
	if target == 0 {
		return true
	}

	if len(input) == 0 || target < 0 {
		return false
	}

	// Include current element
	if SubsetSum(input[1:], target-input[0]) {
		return true
	}

	// Exclude current element
	if SubsetSum(input[1:], target) {
		return true
	}

	return false
}
