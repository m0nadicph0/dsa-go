package backtrack

import "github.com/m0nadicph0/dsa-go/util"

func robHelperII(wealth []int) int {
	if len(wealth) == 0 {
		return 0
	}

	if len(wealth) == 1 {
		return wealth[0]
	}

	return util.Max(wealth[0]+robHelperII(wealth[2:]), robHelperII(wealth[1:]))
}

func RobII(wealth []int) int {
	if len(wealth) == 0 {
		return 0
	}

	if len(wealth) == 1 {
		return wealth[0]
	}

	// Robbing the first house
	robFirst := wealth[0] + robHelperII(wealth[2:len(wealth)-1])

	// Skip the first house
	skipFirst := robHelperII(wealth[1:])

	// Return the maximum of the two options
	return util.Max(robFirst, skipFirst)
}
