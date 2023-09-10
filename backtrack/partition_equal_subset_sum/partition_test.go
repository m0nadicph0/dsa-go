package partition_equal_subset_sum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCanPartition(t *testing.T) {

	tests := []struct {
		name  string
		input []int
		want  bool
	}{
		{
			name:  "Case 1",
			input: []int{1, 5, 11, 5},
			want:  true,
		},
		{
			name:  "Case 2",
			input: []int{1, 2, 3, 5},
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, CanPartition(tt.input))
		})
	}
}
