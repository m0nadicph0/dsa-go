package twoptr

import "sort"

func ThreeSum(nums []int, target int) bool {
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			if sum == target {
				return true
			} else if sum < target {
				left++
			} else {
				right--
			}
		}
	}
	return false
}
