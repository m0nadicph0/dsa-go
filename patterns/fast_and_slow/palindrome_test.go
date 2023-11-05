package fast_and_slow

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPalindrome(t *testing.T) {

	tests := []struct {
		name string
		head *Node
		want bool
	}{
		{
			name: "Case 1",
			head: palindromeEven(),
			want: true,
		},
		{
			name: "Case 2",
			head: palindromeOdd(),
			want: true,
		},
		{
			name: "Case 3",
			head: nonPalindromeEven(),
			want: false,
		},
		{
			name: "Case 4",
			head: nonPalindromeOdd(),
			want: false,
		},
		{
			name: "Case 5",
			head: nil,
			want: true,
		},
		{
			name: "Case 6",
			head: &Node{
				value: 1,
				next:  nil,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, IsPalindrome(tt.head), "IsPalindrome(%v)", tt.head)
		})
	}
}

func nonPalindromeOdd() *Node {
	return &Node{
		value: 1,
		next: &Node{
			value: 2,
			next: &Node{
				value: 3,
				next: &Node{
					value: 4,
					next: &Node{
						value: 5,
						next:  nil,
					},
				},
			},
		},
	}
}

func nonPalindromeEven() *Node {
	return &Node{
		value: 1,
		next: &Node{
			value: 2,
			next: &Node{
				value: 3,
				next: &Node{
					value: 4,
					next:  nil,
				},
			},
		},
	}
}

func palindromeOdd() *Node {
	return &Node{
		value: 1,
		next: &Node{
			value: 2,
			next: &Node{
				value: 3,
				next: &Node{
					value: 2,
					next: &Node{
						value: 1,
						next:  nil,
					},
				},
			},
		},
	}
}

func palindromeEven() *Node {
	return &Node{
		value: 1,
		next: &Node{
			value: 2,
			next: &Node{
				value: 2,
				next: &Node{
					value: 1,
					next:  nil,
				},
			},
		},
	}
}
