package editdistance

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEditDistance(t *testing.T) {

	tests := []struct {
		name  string
		word1 string
		word2 string
		want  int
	}{
		{
			name:  "Case 1",
			word1: "kitten",
			word2: "sitting",
			want:  3,
		},
		{
			name:  "Case 2",
			word1: "cat",
			word2: "catalogue",
			want:  6,
		},
		{
			name:  "Case 3",
			word1: "cat",
			word2: "bata",
			want:  2,
		},
		{
			name:  "Case 3",
			word1: "sunday",
			word2: "saturday",
			want:  3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, EditDistance(tt.word1, tt.word2))
		})
	}
}
