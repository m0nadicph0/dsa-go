package minimum_coin_change

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinCoinChange(t *testing.T) {

	tests := []struct {
		name   string
		coins  []int
		amount int
		want   int
	}{
		{
			name:   "Case 0",
			coins:  []int{1, 5, 10},
			amount: 12,
			want:   3,
		},
		{
			name:   "Case 1",
			coins:  []int{1},
			amount: 12,
			want:   12,
		},
		{
			name:   "Case 2",
			coins:  []int{1, 2},
			amount: 3,
			want:   2,
		},
		{
			name:   "Case 2",
			coins:  []int{1, 10},
			amount: 33,
			want:   6,
		},
		{
			name:   "",
			coins:  []int{1, 3, 4},
			amount: 10,
			want:   3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MinCoinChange(tt.coins, tt.amount)
			assert.Equal(t, tt.want, got, "MinCoinChange(%v, %d) = %v", tt.coins, tt.amount, got)
		})
	}
}
