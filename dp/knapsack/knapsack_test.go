package knapsack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKnapsack(t *testing.T) {
	tests := []struct {
		name     string
		weights  []int
		values   []int
		capacity int
		want     int
	}{
		{
			name:     "Case 1",
			weights:  []int{23, 31, 29, 44, 53, 38, 63, 85, 89, 82},
			values:   []int{92, 57, 49, 68, 60, 43, 67, 84, 87, 72},
			capacity: 165,
			want:     309,
		},
		{
			name:     "Case 2",
			weights:  []int{4, 5, 6},
			values:   []int{1, 2, 3},
			capacity: 3,
			want:     0,
		},
		{
			name:     "Case 3",
			weights:  []int{4, 5, 1},
			values:   []int{1, 2, 3},
			capacity: 4,
			want:     3,
		},
		{
			name:     "Case 3",
			weights:  []int{4, 5, 1},
			values:   []int{1, 2, 3},
			capacity: 4,
			want:     3,
		},
		{
			name:     "Case 4",
			weights:  []int{41, 50, 49, 59, 55, 57, 60},
			values:   []int{442, 525, 511, 593, 546, 564, 617},
			capacity: 170,
			want:     1735,
		},
		{
			name:     "Case 5",
			weights:  []int{4, 5, 6},
			values:   []int{1, 2, 3},
			capacity: 15,
			want:     6,
		},
		{
			name:     "Case 6",
			weights:  []int{4, 5, 1, 3, 2, 5},
			values:   []int{2, 3, 1, 5, 4, 7},
			capacity: 15,
			want:     19,
		},
		{
			name:     "Case 7",
			weights:  []int{4, 5, 1},
			values:   []int{1, 2, 3},
			capacity: 4,
			want:     3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Knapsack(tt.weights, tt.values, tt.capacity)
			assert.Equal(t, tt.want, got, "Knapsack(%v, %v, %d)=%d", tt.weights, tt.values, tt.capacity, got)
		})
	}
}
