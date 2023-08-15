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

func TestLinkedListPrepend(t *testing.T) {
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
			wantElements: []int{4, 3, 2, 1},
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
				list.Prepend(item)
			}
			gotSize := list.Size()
			gotElements := make([]int, 0)
			list.ForEach(func(value int) {
				gotElements = append(gotElements, value)
			})

			assert.Equal(t, tc.wantSize, gotSize)
			assert.Equal(t, tc.wantElements, gotElements)
		})
	}
}

func TestLinkedListPrependAndAppend(t *testing.T) {
	testsCases := []struct {
		name         string
		inputPrepend []int
		inputAppend  []int
		wantSize     int
		wantElements []int
	}{
		{
			name:         "Case 0",
			inputPrepend: []int{},
			inputAppend:  []int{},
			wantSize:     0,
			wantElements: []int{},
		},
		{
			name:         "Case 1",
			inputPrepend: []int{1},
			inputAppend:  []int{2},
			wantSize:     2,
			wantElements: []int{1, 2},
		},
		{
			name:         "Case 2",
			inputPrepend: []int{1, 2, 3},
			inputAppend:  []int{4, 5, 6},
			wantSize:     6,
			wantElements: []int{3, 2, 1, 4, 5, 6},
		},
		{
			name:         "Case 3",
			inputPrepend: []int{4, 3, 2, 1},
			inputAppend:  []int{5, 6, 7, 8},
			wantSize:     8,
			wantElements: []int{1, 2, 3, 4, 5, 6, 7, 8},
		},
	}

	for _, tc := range testsCases {
		t.Run(tc.name, func(t *testing.T) {
			list := NewLinkedList()
			for _, item := range tc.inputPrepend {
				list.Prepend(item)
			}
			for _, item := range tc.inputAppend {
				list.Append(item)
			}

			gotSize := list.Size()
			gotElements := make([]int, 0)
			list.ForEach(func(value int) {
				gotElements = append(gotElements, value)
			})

			assert.Equal(t, tc.wantSize, gotSize)
			assert.Equal(t, tc.wantElements, gotElements)
		})
	}
}

func TestLinkedListGet(t *testing.T) {
	testsCases := []struct {
		name   string
		input  []int
		target int
		want   int
	}{
		{
			name:   "Case 0",
			input:  []int{},
			target: 0,
			want:   -1,
		},
		{
			name:   "Case 1",
			input:  []int{1},
			target: 0,
			want:   1,
		},
		{
			name:   "Case 2",
			input:  []int{1, 2, 3, 4},
			target: 3,
			want:   4,
		},
		{
			name:   "Case 3",
			input:  []int{1, 2, 3, 3, 2, 1},
			target: 3,
			want:   3,
		},
		{
			name:   "Case 4",
			input:  []int{1, 2, 3, 3, 2, 1},
			target: 0,
			want:   1,
		},
		{
			name:   "Case 5",
			input:  []int{1, 2, 3, 4},
			target: 4,
			want:   -1,
		},
	}

	for _, tc := range testsCases {
		t.Run(tc.name, func(t *testing.T) {
			list := NewLinkedList()
			for _, item := range tc.input {
				list.Append(item)
			}
			got := list.Get(tc.target)
			assert.Equal(t, tc.want, got, "list.Get(%d)=%d", tc.target, got)
		})
	}
}

func TestLinkedListInsertAt(t *testing.T) {
	testsCases := []struct {
		name         string
		initial      []int
		value        int
		position     int
		wantSize     int
		wantElements []int
		wantErr      bool
	}{
		{
			name:         "Case 0",
			initial:      []int{},
			value:        2,
			position:     0,
			wantSize:     1,
			wantElements: []int{2},
			wantErr:      false,
		},
		{
			name:         "Case 1",
			initial:      []int{1, 2, 3},
			value:        4,
			position:     1,
			wantSize:     4,
			wantElements: []int{1, 4, 2, 3},
			wantErr:      false,
		},
		{
			name:         "Case 2",
			initial:      []int{1, 2, 3},
			value:        0,
			position:     0,
			wantSize:     4,
			wantElements: []int{0, 1, 2, 3},
			wantErr:      false,
		},
		{
			name:         "Case 3",
			initial:      []int{1, 2, 3},
			value:        4,
			position:     3,
			wantSize:     4,
			wantElements: []int{1, 2, 3, 4},
			wantErr:      false,
		},
		{
			name:         "Case 4",
			initial:      []int{1, 2, 3},
			value:        4,
			position:     -1,
			wantSize:     3,
			wantElements: []int{1, 2, 3},
			wantErr:      true,
		},
		{
			name:         "Case 5",
			initial:      []int{1, 2, 3},
			value:        4,
			position:     4,
			wantSize:     3,
			wantElements: []int{1, 2, 3},
			wantErr:      true,
		},
	}

	for _, tc := range testsCases {
		t.Run(tc.name, func(t *testing.T) {
			list := NewLinkedList()
			for _, item := range tc.initial {
				list.Append(item)
			}
			err := list.InsertAt(tc.position, tc.value)
			if tc.wantErr {
				assert.Error(t, err, "should throw error")
			} else {
				assert.NoError(t, err, "should not throw error")
			}
			gotSize := list.Size()
			gotElements := make([]int, 0)
			list.ForEach(func(v int) {
				gotElements = append(gotElements, v)
			})
			assert.Equal(t, tc.wantSize, gotSize, "list.Size()=%d", gotSize)
			assert.Equal(t, tc.wantElements, gotElements, "list=[%v]", gotElements)
		})
	}
}

func TestInsertAtAndPrepend(t *testing.T) {
	ll := NewLinkedList()
	ll.InsertAt(0, 1)
	ll.Prepend(2)
	ll.InsertAt(0, 3)

	assert.Equal(t, 3, ll.Size(), "list.Size()=%d", ll.Size())

	gotElements := make([]int, 0)
	ll.ForEach(func(v int) {
		gotElements = append(gotElements, v)
	})

	assert.Equal(t, []int{3, 2, 1}, gotElements, "list=[%v]", gotElements)
}

func TestInsertAtAndAppend(t *testing.T) {
	ll := NewLinkedList()
	ll.InsertAt(0, 1)
	ll.Append(2)
	ll.InsertAt(2, 3)
	ll.Append(4)

	assert.Equal(t, 4, ll.Size(), "list.Size()=%d", ll.Size())

	gotElements := make([]int, 0)
	ll.ForEach(func(v int) {
		gotElements = append(gotElements, v)
	})

	assert.Equal(t, []int{1, 2, 3, 4}, gotElements, "list=[%v]", gotElements)
}
