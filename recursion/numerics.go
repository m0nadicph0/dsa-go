package recursion

func CountDigits(n int) int {
	if n < 10 {
		return 1
	}
	return CountDigits(n/10) + 1
}

func DigitSum(n int) int {
	if n < 10 {
		return n
	}
	return DigitSum(n/10) + n%10
}

func Multiply(m int, n int) int {
	if n == 0 {
		return 0
	}
	return Multiply(m, n-1) + m
}
