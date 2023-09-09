package backtrack

func BreakWordToSentences(input string, words []string) []string {
	memo := make(map[int][]string)
	return breakWordToSentencesHelper(input, NewDictionary(words), 0, memo)
}

func breakWordToSentencesHelper(input string, dictionary Dictionary, start int, memo map[int][]string) []string {
	if start == len(input) {
		return []string{""}
	}

	if val, ok := memo[start]; ok {
		return val
	}

	var result []string
	for i := start + 1; i <= len(input); i++ {
		substr := input[start:i]
		if dictionary.Contains(substr) {
			nextWords := breakWordToSentencesHelper(input, dictionary, i, memo)
			for _, nextWord := range nextWords {
				if nextWord == "" {
					result = append(result, substr)
				} else {
					result = append(result, substr+" "+nextWord)
				}
			}
		}
	}

	memo[start] = result
	return result
}
