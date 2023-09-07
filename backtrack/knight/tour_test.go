package knight

import "testing"

func TestKnightsTour(t *testing.T) {

	tests := []struct {
		name   string
		board  Board
		startX int
		startY int
		want   bool
	}{
		{
			name:   "Case 1",
			board:  NewBoard(8),
			startX: 0,
			startY: 0,
			want:   true,
		},
		{
			name:   "Case 1",
			board:  NewBoard(8),
			startX: 7,
			startY: 0,
			want:   true,
		},
		{
			name:   "Case 2",
			board:  NewBoard(3),
			startX: 0,
			startY: 0,
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := KnightsTour(tt.board, tt.startX, tt.startY); got != tt.want {
				t.Errorf("KnightsTour() = %v, want %v", got, tt.want)
			}
		})
	}
}
