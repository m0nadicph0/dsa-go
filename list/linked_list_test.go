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

func TestLinkedListDeleteAt(t *testing.T) {
	testsCases := []struct {
		name         string
		initial      []int
		position     int
		wantSize     int
		wantElements []int
		wantErr      bool
	}{
		{
			name:         "Deleting from an empty list",
			initial:      []int{},
			position:     0,
			wantSize:     0,
			wantElements: []int{},
			wantErr:      true,
		},
		{
			name:         "Deleting from a list with one element",
			initial:      []int{1},
			position:     0,
			wantSize:     0,
			wantElements: []int{},
			wantErr:      false,
		},
		{
			name:         "Deleting the first element from a list with multiple elements",
			initial:      []int{1, 2, 3, 4},
			position:     0,
			wantSize:     3,
			wantElements: []int{2, 3, 4},
			wantErr:      false,
		},
		{
			name:         "Deleting the last element from a list with multiple elements",
			initial:      []int{1, 2, 3, 4},
			position:     3,
			wantSize:     3,
			wantElements: []int{1, 2, 3},
			wantErr:      false,
		},
		{
			name:         "Deleting an element from the middle of a list with multiple elements",
			initial:      []int{1, 2, 3, 4, 5},
			position:     2,
			wantSize:     4,
			wantElements: []int{1, 2, 4, 5},
			wantErr:      false,
		},
		{
			name:         "Deleting from an index greater than list size",
			initial:      []int{15, 25, 35},
			position:     3,
			wantSize:     3,
			wantElements: []int{15, 25, 35},
			wantErr:      true,
		},
		{
			name:         "Deleting from a negative index",
			initial:      []int{5, 10, 15},
			position:     -1,
			wantSize:     3,
			wantElements: []int{5, 10, 15},
			wantErr:      true,
		},
		{
			name:         "Deleting from a list with only one element with negative index",
			initial:      []int{7},
			position:     -1,
			wantSize:     1,
			wantElements: []int{7},
			wantErr:      true,
		},
		{
			name:         "Deleting the last element from a list with multiple elements",
			initial:      []int{5, 10, 15},
			position:     2,
			wantSize:     2,
			wantElements: []int{5, 10},
			wantErr:      false,
		},
	}

	for _, tc := range testsCases {
		t.Run(tc.name, func(t *testing.T) {
			list := NewLinkedList()
			for _, item := range tc.initial {
				list.Append(item)
			}
			err := list.DeleteAt(tc.position)

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

func TestLinkedListLength(t *testing.T) {
	testsCases := []struct {
		name    string
		initial []int
		want    int
	}{
		{
			name:    "Counting nodes in an empty list",
			initial: []int{},
			want:    0,
		},
		{
			name:    "Counting nodes in a list with one element",
			initial: []int{5},
			want:    1,
		},
		{
			name:    "Counting nodes in a list with multiple elements",
			initial: []int{10, 20, 30, 40},
			want:    4,
		},
	}

	for _, tc := range testsCases {
		t.Run(tc.name, func(t *testing.T) {
			list := NewLinkedList()
			for _, item := range tc.initial {
				list.Append(item)
			}
			got := list.Length()
			assert.Equal(t, tc.want, got, "list.Length()=%d", got)
		})
	}

	t.Run("Counting nodes in a list after appending elements", func(t *testing.T) {
		ll := NewLinkedList()
		ll.Append(15)
		ll.Append(25)
		ll.Append(35)
		assert.Equal(t, 3, ll.Length(), "list.Length()=%d", ll.Length())
	})

	t.Run("Counting nodes in a list after inserting elements", func(t *testing.T) {
		ll := NewLinkedList()
		ll.InsertAt(0, 5)
		ll.InsertAt(1, 10)
		ll.InsertAt(1, 15)
		assert.Equal(t, 3, ll.Length(), "list.Length()=%d", ll.Length())
	})

	t.Run("Counting nodes in a list after deleting elements", func(t *testing.T) {
		ll := NewLinkedList()
		for i := 0; i < 10; i++ {
			ll.Append(i)
		}

		for i := 0; i < 5; i++ {
			ll.DeleteAt(i)
		}
		assert.Equal(t, 5, ll.Length(), "list.Length()=%d", ll.Length())
	})

	t.Run("Counting nodes after performing various operations", func(t *testing.T) {
		ll := NewLinkedList()
		ll.Append(1)
		ll.Append(2)
		ll.Prepend(0)
		ll.Prepend(-1)
		ll.InsertAt(0, -2)
		ll.DeleteAt(1)
		ll.DeleteAt(1)
		assert.Equal(t, 3, ll.Length(), "list.Length()=%d", ll.Length())
	})

	t.Run("Counting nodes in a list with large number of elements", func(t *testing.T) {
		ll := NewLinkedList()
		for i := 0; i < 10000; i++ {
			ll.Append(i)
		}
		assert.Equal(t, 10000, ll.Length(), "list.Length()=%d", ll.Length())

	})

	t.Run("Counting nodes in a list with negative values", func(t *testing.T) {
		ll := NewLinkedList()
		ll.Append(-5)
		ll.Append(-10)
		ll.Append(-15)
		assert.Equal(t, 3, ll.Length(), "list.Length()=%d", ll.Length())
	})

}

func TestLinkedListIsPresent(t *testing.T) {
	testsCases := []struct {
		name    string
		initial []int
		value   int
		want    bool
	}{
		{
			name:    "Checking for presence in an empty list",
			initial: []int{},
			value:   5,
			want:    false,
		},
		{
			name:    "Checking for presence in a list with one element (element present)",
			initial: []int{10},
			value:   10,
			want:    true,
		},
		{
			name:    "Checking for presence in a list with one element (element not present)",
			initial: []int{20},
			value:   5,
			want:    false,
		},
		{
			name:    "Checking for presence in a list with multiple elements (element present)",
			initial: []int{5, 10, 15, 20},
			value:   15,
			want:    true,
		},
		{
			name:    "Checking for presence in a list with multiple elements (element not present)",
			initial: []int{5, 10, 15, 20},
			value:   25,
			want:    false,
		},
	}

	for _, tc := range testsCases {
		t.Run(tc.name, func(t *testing.T) {
			list := NewLinkedList()
			for _, item := range tc.initial {
				list.Append(item)
			}
			got := list.IsPresent(tc.value)
			assert.Equal(t, tc.want, got, "list.IsPresent(%d)=%t", tc.value, got)
		})
	}

	t.Run("Checking for presence after appending an element (element present)", func(t *testing.T) {
		list := NewLinkedList()
		list.Append(7)
		list.Append(14)
		list.Append(21)
		valueToCheck := 21
		got := list.IsPresent(valueToCheck)
		assert.Equal(t, true, got, "list.IsPresent(%d)=%t", valueToCheck, got)
	})

	t.Run("Checking for presence after inserting an element (element present)", func(t *testing.T) {
		list := NewLinkedList()
		for _, i := range []int{3, 6, 12} {
			list.Append(i)
		}
		_ = list.InsertAt(2, 9)

		valueToCheck := 9
		got := list.IsPresent(valueToCheck)
		assert.Equal(t, true, got, "list.IsPresent(%d)=%t", valueToCheck, got)
	})

	t.Run("Checking for presence after deleting an element (element not present)", func(t *testing.T) {
		list := NewLinkedList()
		for _, i := range []int{5, 10, 15} {
			list.Append(i)
		}
		_ = list.DeleteAt(1)

		valueToCheck := 10
		got := list.IsPresent(valueToCheck)
		assert.Equal(t, false, got, "list.IsPresent(%d)=%t", valueToCheck, got)
	})

	t.Run("Checking for presence in a list with large number of elements", func(t *testing.T) {
		list := NewLinkedList()
		for i := 0; i < 10000; i++ {
			list.Append(i)
		}
		valueToCheck := 5000
		got := list.IsPresent(valueToCheck)
		assert.Equal(t, true, got, "list.IsPresent(%d)=%t", valueToCheck, got)
	})

	t.Run("Checking for presence of a negative element", func(t *testing.T) {
		list := NewLinkedList()
		for _, i := range []int{-5, -10, -15} {
			list.Append(i)
		}
		valueToCheck := -10
		got := list.IsPresent(valueToCheck)
		assert.Equal(t, true, got, "list.IsPresent(%d)=%t", valueToCheck, got)
	})
}

func TestLinkedListIsPresentRec(t *testing.T) {
	testsCases := []struct {
		name    string
		initial []int
		value   int
		want    bool
	}{
		{
			name:    "Checking for presence in an empty list",
			initial: []int{},
			value:   5,
			want:    false,
		},
		{
			name:    "Checking for presence in a list with one element (element present)",
			initial: []int{10},
			value:   10,
			want:    true,
		},
		{
			name:    "Checking for presence in a list with one element (element not present)",
			initial: []int{20},
			value:   5,
			want:    false,
		},
		{
			name:    "Checking for presence in a list with multiple elements (element present)",
			initial: []int{5, 10, 15, 20},
			value:   15,
			want:    true,
		},
		{
			name:    "Checking for presence in a list with multiple elements (element not present)",
			initial: []int{5, 10, 15, 20},
			value:   25,
			want:    false,
		},
	}

	for _, tc := range testsCases {
		t.Run(tc.name, func(t *testing.T) {
			list := NewLinkedList()
			for _, item := range tc.initial {
				list.Append(item)
			}
			got := list.IsPresentRec(tc.value)
			assert.Equal(t, tc.want, got, "list.IsPresentRec(%d)=%t", tc.value, got)
		})
	}

	t.Run("Checking for presence after appending an element (element present)", func(t *testing.T) {
		list := NewLinkedList()
		list.Append(7)
		list.Append(14)
		list.Append(21)
		valueToCheck := 21
		got := list.IsPresentRec(valueToCheck)
		assert.Equal(t, true, got, "list.IsPresentRec(%d)=%t", valueToCheck, got)
	})

	t.Run("Checking for presence after inserting an element (element present)", func(t *testing.T) {
		list := NewLinkedList()
		for _, i := range []int{3, 6, 12} {
			list.Append(i)
		}
		_ = list.InsertAt(2, 9)

		valueToCheck := 9
		got := list.IsPresentRec(valueToCheck)
		assert.Equal(t, true, got, "list.IsPresentRec(%d)=%t", valueToCheck, got)
	})

	t.Run("Checking for presence after deleting an element (element not present)", func(t *testing.T) {
		list := NewLinkedList()
		for _, i := range []int{5, 10, 15} {
			list.Append(i)
		}
		_ = list.DeleteAt(1)

		valueToCheck := 10
		got := list.IsPresentRec(valueToCheck)
		assert.Equal(t, false, got, "list.IsPresentRec(%d)=%t", valueToCheck, got)
	})

	t.Run("Checking for presence in a list with large number of elements", func(t *testing.T) {
		list := NewLinkedList()
		for i := 0; i < 10000; i++ {
			list.Append(i)
		}
		valueToCheck := 5000
		got := list.IsPresentRec(valueToCheck)
		assert.Equal(t, true, got, "list.IsPresentRec(%d)=%t", valueToCheck, got)
	})

	t.Run("Checking for presence of a negative element", func(t *testing.T) {
		list := NewLinkedList()
		for _, i := range []int{-5, -10, -15} {
			list.Append(i)
		}
		valueToCheck := -10
		got := list.IsPresentRec(valueToCheck)
		assert.Equal(t, true, got, "list.IsPresentRec(%d)=%t", valueToCheck, got)
	})

}
