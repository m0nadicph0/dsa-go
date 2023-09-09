package backtrack

type Dictionary map[string]bool

func NewDictionary(words []string) Dictionary {
	dict := make(map[string]bool)
	for _, word := range words {
		dict[word] = true
	}
	return dict
}

func (d Dictionary) Contains(segment string) bool {
	return d[segment]
}

func wordBreakHelper(input string, dict Dictionary, start int) bool {
	if start == len(input) {
		return true
	}
	for i := start + 1; i <= len(input); i++ {
		segment := input[start:i]
		if dict.Contains(segment) && wordBreakHelper(input, dict, i) {
			return true
		}
	}
	return false
}

func WordBreak(input string, words []string) bool {
	return wordBreakHelper(input, NewDictionary(words), 0)
}
