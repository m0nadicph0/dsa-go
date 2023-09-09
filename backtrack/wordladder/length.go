package wordladder

import "math"

func wordLadderLengthHelper(currentWord string, endWord string, depth int, dictionary Dictionary) int {
	if currentWord == endWord {
		return depth
	}

	// Mark the current word as visited by removing it from the dictionary
	dictionary.Delete(currentWord)

	// Find neighbors of the current word
	neighbors := dictionary.FindNeighbours(currentWord)
	minSteps := math.MaxInt

	for _, neighbor := range neighbors {
		steps := wordLadderLengthHelper(neighbor, endWord, depth+1, dictionary)
		if steps < minSteps {
			minSteps = steps
		}
	}

	// Restore the word in the dictionary
	dictionary[currentWord] = true

	return minSteps
}

func WordLadderLength(beginWord string, endWord string, words []string) int {
	result := wordLadderLengthHelper(beginWord, endWord, 1, NewDictionary(words))
	if result > len(words) {
		return 0
	}
	return result
}
