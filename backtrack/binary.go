package backtrack

func BinaryStrings(n int) []string {
	result := make([]string, 0)
	forEachBinaryString(n, "", func(binStr string) {
		result = append(result, binStr)
	})
	return result
}

func forEachBinaryString(n int, current string, fn func(string)) {
	if n == 0 {
		fn(current)
		return
	}
	forEachBinaryString(n-1, current+"0", fn)
	forEachBinaryString(n-1, current+"1", fn)
}
