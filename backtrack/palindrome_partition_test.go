package backtrack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPartitionPalindrome(t *testing.T) {

	tests := []struct {
		name  string
		input string
		want  [][]string
	}{
		{
			name:  "Case 1",
			input: "aab",
			want:  [][]string{{"a", "a", "b"}, {"aa", "b"}},
		},
		{
			name:  "Case 2",
			input: "a",
			want:  [][]string{{"a"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, PartitionPalindrome(tt.input), "PartitionPalindrome(%v)", tt.input)
		})
	}
}
