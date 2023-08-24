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

func forEachCombination(input []int, k int, start int, current []int, fn visitFn) {
	if k == len(current) {
		fn(current)
		return
	}

	for i := start; i < len(input); i++ {
		current = append(current, input[i])
		forEachCombination(input, k, i+1, current, fn)
		current = current[:len(current)-1]
	}
}
