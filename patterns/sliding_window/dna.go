package sliding_window

func FindRepeatedSequences(s string, k int) Set {
	output := *NewSet()
	store := *NewSet()

	left, right := 0, k-1

	for right < len(s) {
		subSequence := s[left : right+1]
		if store.Exists(subSequence) {
			output.Add(subSequence)
		} else {
			store.Add(subSequence)
		}
		left++
		right++
	}
	return output
}
