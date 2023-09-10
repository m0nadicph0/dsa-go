package backtrack

import "github.com/m0nadicph0/dsa-go/util"

func robHelper(wealth []int, index int) int {
	if index >= len(wealth) {
		return 0
	}

	// Rob the current house and move two steps forward
	robCurrent := wealth[index] + robHelper(wealth, index+2)

	// Skip the current house and move one step forward
	skipCurrent := robHelper(wealth, index+1)

	// Return the maximum of the two options
	return util.Max(robCurrent, skipCurrent)
}

func Rob(wealth []int) int {
	if len(wealth) == 0 {
		return 0
	}
	return robHelper(wealth, 0)
}
