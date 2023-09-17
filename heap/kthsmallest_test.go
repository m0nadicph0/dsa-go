package heap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindKthSmallest(t *testing.T) {
	tests := []struct {
		input    []int
		k        int
		expected int
	}{
		{[]int{3, 1, 4, 2, 5}, 3, 3},    // 3rd smallest element is 3
		{[]int{9, 7, 2, 6, 5, 8}, 2, 5}, // 2nd smallest element is 5
		{[]int{1, 2, 3, 4, 5}, 5, 5},    // 5th smallest element is 5
		{[]int{10, 10, 10, 10}, 1, 10},  // 1st smallest element is 10
		{[]int{6, 4, 9, 2, 1, 3}, 4, 4}, // 4th smallest element is 4
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			result := FindKthSmallest(test.input, test.k)
			assert.Equal(t, test.expected, result)
		})
	}
}
