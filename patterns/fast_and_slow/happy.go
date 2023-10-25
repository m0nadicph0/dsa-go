package fast_and_slow

func sumOfDigitsSquared(num int) int {
	sum := 0
	for num > 0 {
		last := num % 10
		sum += last * last
		num = num / 10
	}
	return sum
}

func IsHappy(num int) bool {
	slow, fast := num, sumOfDigitsSquared(num)
	for fast != 1 && slow != fast {
		slow = sumOfDigitsSquared(slow)
		fast = sumOfDigitsSquared(sumOfDigitsSquared(fast))
	}
	return fast == 1
}
