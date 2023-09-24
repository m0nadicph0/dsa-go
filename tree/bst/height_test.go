package bst

import (
	"testing"
)

func TestCalculateHeight(t *testing.T) {
	testCases := []struct {
		name     string
		root     *Node
		expected int
	}{
		{
			name:     "Empty Tree",
			root:     nil,
			expected: 0,
		},
		{
			name: "Single Node",
			root: &Node{
				Value: 42,
			},
			expected: 0,
		},
		{
			name: "Balanced Tree",
			root: &Node{
				Value: 10,
				Left: &Node{
					Value: 5,
					Left: &Node{
						Value: 3,
					},
					Right: &Node{
						Value: 7,
					},
				},
				Right: &Node{
					Value: 15,
					Right: &Node{
						Value: 20,
					},
				},
			},
			expected: 2,
		},
		{
			name: "Unbalanced Tree",
			root: &Node{
				Value: 10,
				Left: &Node{
					Value: 5,
					Left: &Node{
						Value: 3,
					},
				},
			},
			expected: 2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := CalculateHeight(tc.root)
			if result != tc.expected {
				t.Errorf("Expected height %d, but got %d", tc.expected, result)
			}
		})
	}
}
