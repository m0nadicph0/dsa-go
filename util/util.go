package util

import "strconv"

func Swap(a []int, u int, v int) {
	tmp := a[u]
	a[u] = a[v]
	a[v] = tmp
}

func MakeCopy(xs []int) []int {
	tmp := make([]int, len(xs))
	copy(tmp, xs)
	return tmp
}

func Max(nums ...int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	return max(nums[0], Max(nums[1:]...))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(nums ...int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	return min(nums[0], Min(nums[1:]...))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func Sum(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func StrToInt(s string) int {
	value, _ := strconv.Atoi(s)
	return value
}
