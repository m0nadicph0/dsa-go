package searching

func BinarySearch(input []int, key int) int {
	return binarySearchRec(input, key, 0, len(input)-1)
}

func binarySearchRec(input []int, key int, start int, end int) int {
	if start <= end {
		mid := start + (end-start)/2
		if input[mid] == key {
			return mid
		}

		if key < input[mid] {
			return binarySearchRec(input, key, start, mid-1)
		} else {
			return binarySearchRec(input, key, mid+1, end)
		}
	}
	return -1
}
