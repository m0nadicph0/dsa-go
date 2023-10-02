package twoptr

func ReverseWords(sentence string) string {
	inputBytes := []byte(sentence)
	reverse(inputBytes, 0, len(sentence)-1)

	start, end := 0, 0
	for start < len(sentence) {
		for end < len(sentence) && inputBytes[end] != ' ' {
			end++
		}
		reverse(inputBytes, start, end-1)
		start = end + 1
		end++
	}

	return string(inputBytes)
}

func reverse(input []byte, start, end int) {
	for start < end {
		input[start], input[end] = input[end], input[start]
		start++
		end--
	}
}
