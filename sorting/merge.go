package sorting

func MergeSort(input []int) {
	sorted := mergeSortRec(input)
	copy(input, sorted)
}

func mergeSortRec(input []int) []int {
	if len(input) <= 1 {
		return input
	}
	mid := len(input) / 2
	left := mergeSortRec(input[0:mid])
	right := mergeSortRec(input[mid:])
	return merge(left, right)
}

func merge(left []int, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}
