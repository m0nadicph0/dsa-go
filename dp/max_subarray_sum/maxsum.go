package max_subarray_sum

import (
	"github.com/m0nadicph0/dsa-go/util"
	"math"
)

func MaximumSubarraySum(nums []int) int {
	return maxSubArraySumHelper(nums, 0, len(nums)-1)
}

func maxSubArraySumHelper(nums []int, left, right int) int {
	if left == right {
		return nums[left]
	}

	mid := (left + right) / 2

	leftMax := maxSubArraySumHelper(nums, left, mid)
	rightMax := maxSubArraySumHelper(nums, mid+1, right)
	crossingMax := maxCrossingSum(nums, left, mid, right)

	return util.Max(leftMax, rightMax, crossingMax)
}

func maxCrossingSum(nums []int, left, mid, right int) int {
	leftSumMax, rightSumMax := -1*math.MaxInt, -1*math.MaxInt
	sum := 0

	for i := mid; i >= left; i-- {
		sum += nums[i]
		if sum > leftSumMax {
			leftSumMax = sum
		}
	}

	sum = 0
	for i := mid + 1; i <= right; i++ {
		sum += nums[i]
		if sum > rightSumMax {
			rightSumMax = sum
		}
	}

	return leftSumMax + rightSumMax
}
