package backtrack

import "github.com/m0nadicph0/dsa-go/util"

type VisitFn func([]int)

func subsets(nums []int) [][]int {
	result := make([][]int, 0)
	forEachSubset(nums, []int{}, func(subset []int) {

		result = append(result, util.MakeCopy(subset))
	})
	return result
}

func forEachSubset(nums []int, output []int, fn VisitFn) {
	if len(nums) == 0 {
		fn(output)
		return
	}

	// Exclude current element
	forEachSubset(nums[1:], output, fn)

	// Include current element
	forEachSubset(nums[1:], append(output, nums[0]), fn)
}
