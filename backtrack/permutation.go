package backtrack

import "github.com/m0nadicph0/dsa-go/util"

func Permutations(input []int) [][]int {
	result := make([][]int, 0)
	forEachPermutation(input, 0, len(input)-1, func(permutation []int) {
		result = append(result, makeCopy(permutation))
	})
	return result
}

func makeCopy(xs []int) []int {
	tmp := make([]int, len(xs))
	copy(tmp, xs)
	return tmp
}

func forEachPermutation(input []int, start int, end int, fn visitFn) {
	if start == end {
		fn(input)
		return
	}

	for i := start; i <= end; i++ {
		util.Swap(input, start, i)
		forEachPermutation(input, start+1, end, fn)
		util.Swap(input, start, i)
	}
}
