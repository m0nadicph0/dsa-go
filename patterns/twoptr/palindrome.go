package twoptr

func IsPalindrome(input string) bool {
	left, right := 0, len(input)-1
	for left <= right {
		if input[left] != input[right] {
			return false
		}
		left++
		right--
	}
	return true
}
