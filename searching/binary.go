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

func BinarySearchRange(a []int, key int) (int, int) {
	return binarySearchFirst(a, key), binarySearchLast(a, key)
}

func binarySearchFirst(a []int, key int) int {
	result := -1
	left, right := 0, len(a)-1

	for left <= right {
		mid := left + (right-left)/2

		if a[mid] == key {
			result = mid
			right = mid - 1
		} else if key < a[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return result
}

func binarySearchLast(a []int, key int) int {
	left, right := 0, len(a)-1
	result := -1

	for left <= right {
		mid := left + (right-left)/2

		if a[mid] == key {
			result = mid
			left = mid + 1 // Move right to find the last occurrence
		} else if key < a[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return result
}
