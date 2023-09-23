package editdistance

import "github.com/m0nadicph0/dsa-go/util"

func EditDistance(word1 string, word2 string) int {
	return editDistanceHelper(word1, word2, len(word1), len(word2))
}

func editDistanceHelper(word1 string, word2 string, m int, n int) int {
	// Base cases
	if m == 0 {
		return n
	}
	if n == 0 {
		return m
	}

	// If the last characters of the strings match, no operation is needed.
	if word1[m-1] == word2[n-1] {
		return editDistanceHelper(word1, word2, m-1, n-1)
	}

	// Find the minimum of three operations: insertion, deletion, and replacement.
	return 1 + util.Min(
		editDistanceHelper(word1, word2, m, n-1),   // Insert
		editDistanceHelper(word1, word2, m-1, n),   // Delete
		editDistanceHelper(word1, word2, m-1, n-1), // Replace
	)
}
