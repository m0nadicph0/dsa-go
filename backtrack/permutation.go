package backtrack

import "github.com/m0nadicph0/dsa-go/util"

func Permutations(input []int) [][]int {
	result := make([][]int, 0)
	forEachPermutation(input, 0, len(input)-1, func(permutation []int) {
		result = append(result, util.MakeCopy(permutation))
	})
	return result
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

func PermutationsNK(input []int, k int) [][]int {
	result := make([][]int, 0)
	forEachPermutationNK(input, []int{}, make([]bool, len(input)), k, func(perm []int) {
		result = append(result, util.MakeCopy(perm))
	})
	return result
}

func forEachPermutationNK(items []int, current []int, used []bool, k int, fn visitFn) {
	if len(current) == k {
		fn(current)
		return
	}

	for i := 0; i < len(items); i++ {
		if used[i] {
			continue
		}

		used[i] = true
		current = append(current, items[i])
		forEachPermutationNK(items, current, used, k, fn)
		used[i] = false
		current = current[:len(current)-1]
	}
}
