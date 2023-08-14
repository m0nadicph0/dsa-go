package list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkedListAppend(t *testing.T) {
	testsCases := []struct {
		name         string
		input        []int
		wantSize     int
		wantElements []int
	}{
		{
			name:         "Case 0",
			input:        []int{},
			wantSize:     0,
			wantElements: []int{},
		},
		{
			name:         "Case 1",
			input:        []int{1},
			wantSize:     1,
			wantElements: []int{1},
		},
		{
			name:         "Case 2",
			input:        []int{1, 2, 3, 4},
			wantSize:     4,
			wantElements: []int{1, 2, 3, 4},
		},
		{
			name:         "Case 3",
			input:        []int{1, 2, 3, 3, 2, 1},
			wantSize:     6,
			wantElements: []int{1, 2, 3, 3, 2, 1},
		},
	}

	for _, tc := range testsCases {
		t.Run(tc.name, func(t *testing.T) {
			list := NewLinkedList()
			for _, item := range tc.input {
				list.Append(item)
			}
			gotSize := tc.wantSize
			gotElements := make([]int, 0)
			list.ForEach(func(value int) {
				gotElements = append(gotElements, value)
			})

			assert.Equal(t, tc.wantSize, gotSize)
			assert.Equal(t, tc.wantElements, gotElements)
		})
	}
}
