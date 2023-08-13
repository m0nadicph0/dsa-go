package util

func Swap(a []int, u int, v int) {
	tmp := a[u]
	a[u] = a[v]
	a[v] = tmp
}
