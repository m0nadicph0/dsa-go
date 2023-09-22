package lcs

import "github.com/m0nadicph0/dsa-go/util"

func LengthOfLongestCommonSubsequence(text1 string, text2 string) int {
	return lcsLenHelper(text1, text2, 0, 0)
}

func lcsLenHelper(text1 string, text2 string, i int, j int) int {
	if i == len(text1) || j == len(text2) {
		return 0
	}

	if text1[i] == text2[j] {
		// If the current characters match, include them in the LCS.
		return 1 + lcsLenHelper(text1, text2, i+1, j+1)
	} else {
		// If the current characters don't match, try both possibilities (excluding one character at a time).
		return util.Max(lcsLenHelper(text1, text2, i+1, j), lcsLenHelper(text1, text2, i, j+1))
	}
}
