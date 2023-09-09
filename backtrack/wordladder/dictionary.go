package wordladder

type Dictionary map[string]bool

func NewDictionary(words []string) Dictionary {
	dict := make(Dictionary)
	for _, word := range words {
		dict[word] = true
	}
	return dict
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
func (d Dictionary) FindNeighbours(word string) []string {
	neighbors := make(map[string]bool)
	for i := 0; i < len(word); i++ {
		for c := 'a'; c <= 'z'; c++ {
			newWord := word[:i] + string(c) + word[i+1:]
			if d[newWord] {
				neighbors[newWord] = true
			}
		}
	}
	return keys(neighbors)
}

func keys(neighbors map[string]bool) []string {
	result := make([]string, 0)
	for k, _ := range neighbors {
		result = append(result, k)
	}
	return result
}
