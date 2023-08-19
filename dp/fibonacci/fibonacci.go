package fibonacci

// Fibonacci brute force recursion
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-2) + Fibonacci(n-1)
}

// MemoFibonacci memoised solution
func MemoFibonacci(n int) int {
	return memoFibonacciHelper(n, make(map[int]int))
}

func memoFibonacciHelper(n int, memo map[int]int) int {
	if n <= 1 {
		return n
	}

	value, ok := memo[n]
	if ok {
		return value
	}

	computedValue := memoFibonacciHelper(n-2, memo) + memoFibonacciHelper(n-1, memo)

	memo[n] = computedValue

	return computedValue
}

// BottomUpDPFibonacci bottom up DP
func BottomUpDPFibonacci(n int) interface{} {
	if n <= 1 {
		return n
	}
	table := make([]int, n+1)
	table[0] = 0
	table[1] = 1

	for i := 2; i <= n; i++ {
		table[i] = table[i-2] + table[i-1]
	}

	return table[n]
}
