package backtrack

import "github.com/m0nadicph0/dsa-go/util"

func SubsetsOfSize(input []int, size int) [][]int {
	var subsets [][]int

	forEachSubsetOfSize(input, size, 0, []int{}, func(subset []int) {
		subsets = append(subsets, util.MakeCopy(subset))
	})
	return subsets
}

func forEachSubsetOfSize(input []int, size int, start int, current []int, fn func([]int)) {
	if len(current) == size {
		fn(current)
		return
	}

	for i := start; i < len(input); i++ {
		forEachSubsetOfSize(input, size, i+1, append(current, input[i]), fn)
	}
}
