package backtrack

import "github.com/m0nadicph0/dsa-go/util"

func CombinationSum(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	forEachCombinationSum(candidates, target, 0, []int{}, func(combination []int) {
		result = append(result, util.MakeCopy(combination))
	})
	return result
}

func forEachCombinationSum(candidates []int, target int, start int, current []int, fn func([]int)) {
	if target == 0 {
		fn(current)
		return
	}

	if target < 0 || start >= len(candidates) {
		return
	}

	for i := start; i < len(candidates); i++ {
		forEachCombinationSum(candidates, target-candidates[i], i, append(current, candidates[i]), fn)
	}
}
