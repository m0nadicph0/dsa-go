package coin_change

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCoinChange(t *testing.T) {

	tests := []struct {
		name   string
		coins  []int
		amount int
		want   int
	}{
		{
			name:   "Case 0",
			coins:  []int{1, 2, 5},
			amount: 5,
			want:   4,
		},
		{
			name:   "Case 1",
			coins:  []int{1, 2, 5},
			amount: 3,
			want:   2,
		},
		{
			name:   "Case 3",
			coins:  []int{},
			amount: 3,
			want:   0,
		},
		{
			name:   "Case 4",
			coins:  []int{4, 5},
			amount: 3,
			want:   0,
		},
		{
			name:   "Case 5",
			coins:  []int{1, 2, 5},
			amount: 5,
			want:   4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CoinChange(tt.coins, tt.amount)
			assert.Equal(t, tt.want, got, "CoinChange(%v, %d)=%d", tt.coins, tt.amount, got)
		})
	}
}
