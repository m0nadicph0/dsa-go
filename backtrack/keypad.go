package backtrack

var keypad = map[uint8]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func LetterCombinations(input string) []string {
	result := make([]string, 0)
	forEachLetterCombinations(input, "", func(combination string) {
		result = append(result, combination)
	})
	return result
}

func forEachLetterCombinations(input string, current string, fn func(string)) {
	if len(input) == 0 {
		fn(current)
		return
	}

	letters := keypad[input[0]]

	for _, letter := range letters {
		forEachLetterCombinations(input[1:], current+string(letter), fn)
	}
}
