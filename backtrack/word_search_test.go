package backtrack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWordSearch(t *testing.T) {

	tests := []struct {
		name  string
		board WordBoard
		word  string
		want  bool
	}{
		{
			name: "Case 1",
			board: WordBoard{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word: "ABCCED",
			want: true,
		},
		{
			name: "Case 2",
			board: WordBoard{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word: "SEE",
			want: true,
		},
		{
			name: "Case 3",
			board: WordBoard{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word: "ABCB",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, WordSearch(tt.board, tt.word), "WordSearch(%v, %v)", tt.board, tt.word)
		})
	}
}
