package sorting

import "github.com/m0nadicph0/dsa-go/util"

func BubbleSort(input []int) {
	for count := 1; count < len(input); count++ {
		for i := 0; i < len(input)-1; i++ {
			if input[i] > input[i+1] {
				util.Swap(input, i, i+1)
			}
		}
	}
}

func OptimizedBubbleSort(input []int) {

	for count := 1; count < len(input); count++ {
		sorted := true
		for i := 0; i < len(input)-1; i++ {
			if input[i] > input[i+1] {
				util.Swap(input, i, i+1)
				sorted = false
			}
		}
		if sorted {
			break
		}
	}
}
