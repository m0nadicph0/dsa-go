package backtrack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMakeSquare(t *testing.T) {

	tests := []struct {
		name        string
		matchSticks []int
		want        bool
	}{
		{
			name:        "Case 1",
			matchSticks: []int{1, 1, 2, 2, 2},
			want:        true,
		},
		{
			name:        "Case 2",
			matchSticks: []int{3, 3, 3, 3, 4},
			want:        false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, MakeSquare(tt.matchSticks), "MakeSquare(%v)", tt.matchSticks)
		})
	}
}
