package sliding_window

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindRepeatedSequences(t *testing.T) {

	tests := []struct {
		name string
		s    string
		k    int
		want []string
	}{
		{
			name: "Case 1",
			s:    "AAAAACCCCCAAAAACCCCCC",
			k:    8,
			want: []string{"AAAAACCC", "AAAACCCC", "AAACCCCC"},
		},
		{
			name: "Case 2",
			s:    "GGGGGGGGGGGGGGGGGGGGGGGGG",
			k:    9,
			want: []string{"GGGGGGGGG"},
		},
		{
			name: "Case 3",
			s:    "TTTTTCCCCCCCTTTTTTCCCCCCCTTTTTTT",
			k:    10,
			want: []string{"CCCCCCCTTT", "CCCCCCTTTT", "CCCCCTTTTT", "CCCCTTTTTT", "TCCCCCCCTT", "TTCCCCCCCT", "TTTCCCCCCC", "TTTTCCCCCC", "TTTTTCCCCC"},
		},
		{
			name: "Case 4",
			s:    "AAAAAACCCCCCCAAAAAAAACCCCCCCTG",
			k:    10,
			want: []string{"AAAAAACCCC", "AAAAACCCCC", "AAAACCCCCC", "AAACCCCCCC"},
		},
		{
			name: "Case 5",
			s:    "ATATATATATATATAT",
			k:    6,
			want: []string{"ATATAT", "TATATA"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindRepeatedSequences(tt.s, tt.k)
			assert.ElementsMatch(t, tt.want, got.All())
		})
	}
}
