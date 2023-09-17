package heap

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestMinHeapInsert(t *testing.T) {

	tests := []struct {
		name         string
		initialHeap  []int
		element      int
		expectedHeap []int
	}{
		{
			name:         "Inserting an element into an empty heap",
			initialHeap:  []int{},
			element:      10,
			expectedHeap: []int{10},
		},
		{
			name:         "Inserting an element into a non-empty heap",
			initialHeap:  []int{1, 3, 5, 4, 7, 10, 6},
			element:      25,
			expectedHeap: []int{1, 3, 5, 4, 7, 10, 6, 25},
		},
		{
			name:         "Inserting an element that becomes the new root",
			initialHeap:  []int{5, 6, 10, 8, 7},
			element:      1,
			expectedHeap: []int{1, 6, 5, 8, 7, 10},
		},
		{
			name:         "Inserting an element that needs to move up multiple levels",
			initialHeap:  []int{4, 6, 9, 7, 8, 14, 16, 21, 26, 25, 30, 28, 21, 32, 34},
			element:      2,
			expectedHeap: []int{2, 4, 9, 6, 8, 14, 16, 7, 26, 25, 30, 28, 21, 32, 34, 21},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &MinHeap{
				heap: tt.initialHeap,
			}
			h.Insert(tt.element)
			assert.Equal(t, tt.expectedHeap, h.heap)
		})
	}
}

func TestMinHeapExtractMin(t *testing.T) {
	tests := []struct {
		name   string
		heap   []int
		want   int
		result []int
	}{
		{
			name:   "Extract min from existing heap",
			heap:   []int{1, 3, 5, 4, 7, 10, 6},
			want:   1,
			result: []int{3, 4, 5, 6, 7, 10},
		},
		{
			name:   "Extract min from empty heap",
			heap:   []int{},
			want:   math.MaxInt,
			result: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &MinHeap{
				heap: tt.heap,
			}
			assert.Equalf(t, tt.want, h.ExtractMin(), "ExtractMin()")
		})
	}
}
