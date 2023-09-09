package wordladder

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDictionary_FindNeighbours(t *testing.T) {

	tests := []struct {
		name string
		dict Dictionary
		word string
		want []string
	}{
		{
			name: "Case 1",
			dict: NewDictionary([]string{"hot", "dot", "dog", "lot", "log", "cog"}),
			word: "hot",
			want: []string{"hot", "dot", "lot"},
		},
		{
			name: "Case 2",
			dict: NewDictionary([]string{"hot", "dot", "dog", "lot", "log", "cog"}),
			word: "lit",
			want: []string{"lot"},
		},
		{
			name: "Case 3",
			dict: NewDictionary([]string{"hot", "dot", "dog", "lot", "log", "cog"}),
			word: "ate",
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.dict.FindNeighbours(tt.word)
			assert.ElementsMatch(t, got, tt.want)
		})
	}
}
