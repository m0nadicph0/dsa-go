package backtrack

type parenVisitFn func(string)

func GenerateParenthesisCombinations(n int) []string {
	result := make([]string, 0)
	forEachParenthesisCombination("", n, n, func(combination string) {
		result = append(result, combination)
	})
	return result
}

func forEachParenthesisCombination(current string, openCount int, closeCount int, fn parenVisitFn) {
	if openCount == 0 && closeCount == 0 {
		fn(current)
		return
	}

	if openCount > 0 {
		forEachParenthesisCombination(current+"(", openCount-1, closeCount, fn)
	}

	if closeCount > openCount {
		forEachParenthesisCombination(current+")", openCount, closeCount-1, fn)
	}
}
