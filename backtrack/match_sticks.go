package backtrack

import "github.com/m0nadicph0/dsa-go/util"

func canMakeSquare(matchSticks []int, sides []int, target int, index int) bool {
	if index == len(matchSticks) {
		return sides[0] == target && sides[1] == target && sides[2] == target
	}

	for i := 0; i < 4; i++ {
		if sides[i]+matchSticks[index] <= target {
			sides[i] += matchSticks[index]
			if canMakeSquare(matchSticks, sides, target, index+1) {
				return true
			}
			sides[i] -= matchSticks[index]
		}
	}

	return false
}

func MakeSquare(matchSticks []int) bool {
	if len(matchSticks) < 4 {
		return false
	}

	sum := util.Sum(matchSticks)

	if sum%4 != 0 {
		return false
	}

	return canMakeSquare(matchSticks, make([]int, 4), sum/4, 0)
}
