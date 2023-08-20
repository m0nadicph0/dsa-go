package backtrack

import "github.com/m0nadicph0/dsa-go/util"

type visitFn func([]int)

func Combinations(input []int, k int) [][]int {
	result := make([][]int, 0)
	forEachCombination(input, k, 0, []int{}, func(combination []int) {
		result = append(result, util.MakeCopy(combination))
	})
	return result
}

func forEachCombination(input []int, k int, start int, combination []int, fn visitFn) {
	if k == 0 {
		fn(combination)
		return
	}

	for i := start; i < len(input); i++ {
		combination = append(combination, input[i])
		forEachCombination(input, k-1, i+1, combination, fn)
		combination = combination[:len(combination)-1] // backtrack
	}
}
