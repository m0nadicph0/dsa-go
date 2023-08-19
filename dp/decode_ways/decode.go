package decode_ways

import "strconv"

// DecodeWays brute force recursion
func DecodeWays(input string) int {
	if len(input) == 0 {
		return 1
	}

	ways := 0

	// count single digits
	if input[0] != '0' {
		ways += DecodeWays(input[1:])
	}

	// count double digits within range 1 to 26
	if len(input) >= 2 {
		twoDigits := toInt(input[:2])
		if twoDigits >= 10 && twoDigits <= 26 {
			ways += DecodeWays(input[2:])
		}
	}
	return ways
}

func toInt(input string) int {
	n, err := strconv.Atoi(input)
	if err != nil {
		return 0
	}

	return n
}
