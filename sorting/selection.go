package sorting

func SelectionSort(a []int) {
	for i := 0; i < len(a)-1; i++ {
		value, pos := minWithPos(a, i+1)
		if value < a[i] {
			swap(a, i, pos)
		}
	}
}

func swap(a []int, u int, v int) {
	tmp := a[u]
	a[u] = a[v]
	a[v] = tmp
}

func minWithPos(a []int, start int) (int, int) {
	min := a[start]
	pos := start
	for i := start; i < len(a); i++ {
		if a[i] < min {
			min = a[i]
			pos = i
		}
	}
	return min, pos
}
