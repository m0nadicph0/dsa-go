package backtrack

func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func partitionHelper(input string, start int, current []string, fn func([]string)) {
	if start == len(input) {
		fn(current)
		return
	}

	for i := start; i < len(input); i++ {
		if isPalindrome(input[start : i+1]) {
			partitionHelper(input, i+1, append(current, input[start:i+1]), fn)
		}
	}
}

func PartitionPalindrome(input string) [][]string {
	result := make([][]string, 0)

	partitionHelper(input, 0, []string{}, func(strings []string) {
		result = append(result, strings)
	})
	return result
}
