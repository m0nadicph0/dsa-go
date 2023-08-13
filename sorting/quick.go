package sorting

import "github.com/m0nadicph0/dsa-go/util"

func QuickSort(input []int) {
	quickSortRec(input, 0, len(input)-1)
}

func quickSortRec(input []int, start int, end int) {
	if start >= end {
		return
	}
	p := partition(input, start, end)
	quickSortRec(input, start, p-1)
	quickSortRec(input, p+1, end)
}

func partition(input []int, start int, end int) int {
	pivot := input[end]
	i := start - 1
	for j := start; j < end; j++ {
		if input[j] < pivot {
			i += 1
			util.Swap(input, i, j)
		}
	}

	util.Swap(input, i+1, end)
	return i + 1
}
