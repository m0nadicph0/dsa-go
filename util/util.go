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

func Max(a, b int) int {
	if a > b {
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
