package twoptr

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseWords(t *testing.T) {

	tests := []struct {
		name     string
		sentence string
		want     string
	}{
		{
			name:     "Case 1",
			sentence: "seven years in tibet",
			want:     "tibet in years seven",
		},
		{
			name:     "Case 2",
			sentence: "Way to GO",
			want:     "GO to Way",
		},
		{
			name:     "Case 3",
			sentence: "Reverse the words",
			want:     "words the Reverse",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, ReverseWords(tt.sentence), "ReverseWords(%v)", tt.sentence)
		})
	}
}
