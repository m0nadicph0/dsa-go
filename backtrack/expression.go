package backtrack

import "github.com/m0nadicph0/dsa-go/util"

func forEachExpression(input string, target int, expr string, index, currentVal, prevNum int, fn func(string)) {
	if index == len(input) {
		if currentVal == target {
			fn(expr)
		}
		return
	}

	for i := index; i < len(input); i++ {
		if i > index && input[index] == '0' {
			break // Avoid numbers with leading zeros
		}

		currentStr := input[index : i+1]
		currentNum := util.StrToInt(currentStr)

		if index == 0 {
			forEachExpression(input, target, currentStr, i+1, currentNum, currentNum, fn)
		} else {
			// Addition
			forEachExpression(input, target, expr+"+"+currentStr, i+1, currentVal+currentNum, currentNum, fn)

			// Subtraction
			forEachExpression(input, target, expr+"-"+currentStr, i+1, currentVal-currentNum, -currentNum, fn)

			// Multiplication
			forEachExpression(input, target, expr+"*"+currentStr, i+1, currentVal-prevNum+prevNum*currentNum, prevNum*currentNum, fn)
		}
	}
}
func GenerateExpressions(input string, target int) []string {
	var result []string
	forEachExpression(input, target, "", 0, 0, 0, func(expr string) {
		result = append(result, expr)
	})
	return result
}
