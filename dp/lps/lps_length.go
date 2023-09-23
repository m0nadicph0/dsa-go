package lps

import "github.com/m0nadicph0/dsa-go/util"

// LongestPalindromicSubsequence finds the length of the longest subsequence
// (not necessarily contiguous) in a given string that is also a palindrome
func LongestPalindromicSubsequence(input string) int {
	return lpsHelper(input, 0, len(input)-1)
}

func lpsHelper(input string, start, end int) int {
	if start == end {
		return 1
	}

	if start > end {
		return 0
	}

	if input[start] == input[end] {
		// If the characters at the current positions match,
		// include them in the palindrome subsequence.
		return 2 + lpsHelper(input, start+1, end-1)
	}

	// If the characters don't match, try excluding one character at a time
	// and find the maximum length of palindromic subsequences.
	withoutStart := lpsHelper(input, start+1, end)
	withoutEnd := lpsHelper(input, start, end-1)

	// Return the maximum length among the two cases.
	return util.Max(withoutStart, withoutEnd)
}
