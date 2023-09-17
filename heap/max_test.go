package heap

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestMaxHeapInsert(t *testing.T) {
	testCases := []struct {
		name         string
		initialHeap  []int // Initial max-heap as an array
		element      int   // Element to insert
		expectedHeap []int // Expected max-heap after insertion
	}{
		{
			name:         "Inserting an element into an empty heap",
			initialHeap:  []int{},
			element:      10,
			expectedHeap: []int{10},
		},
		{
			name:         "Inserting an element into a non-empty heap",
			initialHeap:  []int{30, 20, 10, 5, 15},
			element:      25,
			expectedHeap: []int{30, 20, 25, 5, 15, 10},
		},
		{
			name:         "Inserting an element that is already in its correct position",
			initialHeap:  []int{40, 30, 20, 10, 5},
			element:      40,
			expectedHeap: []int{40, 30, 40, 10, 5, 20},
		},
		{
			name:         "Inserting an element that becomes the new root",
			initialHeap:  []int{30, 20, 10, 5, 15},
			element:      35,
			expectedHeap: []int{35, 20, 30, 5, 15, 10},
		},
		{
			name:         "Inserting an element that needs to move up multiple levels",
			initialHeap:  []int{30, 20, 10, 5, 15, 25, 12, 7, 8},
			element:      40,
			expectedHeap: []int{40, 30, 10, 5, 20, 25, 12, 7, 8, 15},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			h := NewMaxHeap()
			h.heap = append(h.heap, tc.initialHeap...)
			h.Insert(tc.element)
			assert.Equal(t, tc.expectedHeap, h.heap)
		})
	}
}

func TestMaxHeapExtractMax(t *testing.T) {

	tests := []struct {
		name   string
		heap   []int
		want   int
		result []int
	}{
		{
			name:   "Extract max from existing heap",
			heap:   []int{10, 6, 8, 3, 5, 4, 7},
			want:   10,
			result: []int{8, 6, 7, 3, 5, 4},
		},
		{
			name:   "Extract max from empty heap",
			heap:   []int{},
			want:   -1 * math.MaxInt,
			result: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &MaxHeap{
				heap: tt.heap,
			}
			assert.Equalf(t, tt.want, h.ExtractMax(), "ExtractMax()")
		})
	}
}
