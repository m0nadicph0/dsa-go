package util

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
