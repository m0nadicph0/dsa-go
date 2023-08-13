package searching

func LinearSearch(input []int, key int) int {
	for i := 0; i < len(input); i++ {
		if input[i] == key {
			return i
		}
	}
	return -1
}
