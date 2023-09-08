package backtrack

func RegexMatch(input string, pattern string) bool {
	return regexMatchHelper(input, pattern, 0, 0)
}

func regexMatchHelper(input, pattern string, sIndex, pIndex int) bool {
	// Base case: If we have reached the end of both the string and pattern, it'input a match.
	if pIndex == len(pattern) {
		return sIndex == len(input)
	}

	// Check if the current characters match.
	currentMatch := sIndex < len(input) && (input[sIndex] == pattern[pIndex] || pattern[pIndex] == '.')

	// Check if the next character in the pattern is '*'.
	if pIndex+1 < len(pattern) && pattern[pIndex+1] == '*' {
		// Two choices: (1) Ignore '*' and move to the next character in the pattern
		// or (2) Match the current character and move to the next character in the string.
		return regexMatchHelper(input, pattern, sIndex, pIndex+2) || (currentMatch && regexMatchHelper(input, pattern, sIndex+1, pIndex))
	}

	// If no '*' is present, simply match the current character and move forward.
	return currentMatch && regexMatchHelper(input, pattern, sIndex+1, pIndex+1)
}
