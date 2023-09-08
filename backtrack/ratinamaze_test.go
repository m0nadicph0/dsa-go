package backtrack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolveRatInAMaze(t *testing.T) {

	tests := []struct {
		name   string
		maze   Maze
		want   Maze
		exists bool
	}{
		{
			name: "Case 1",
			maze: NewMaze([][]int{
				{1, 0, 0, 0},
				{1, 1, 0, 1},
				{0, 1, 0, 0},
				{1, 1, 1, 1},
			}),
			want: NewMaze([][]int{
				{1, 0, 0, 0},
				{1, 1, 0, 0},
				{0, 1, 0, 0},
				{0, 1, 1, 1},
			}),
			exists: true,
		},
		{
			name: "Case 1",
			maze: NewMaze([][]int{
				{1, 0, 0, 0},
				{1, 1, 0, 1},
				{0, 0, 0, 0},
				{1, 0, 1, 1},
			}),
			want: NewMaze([][]int{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			}),
			exists: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotSolution := SolveRatInAMaze(tt.maze)
			assert.Equalf(t, tt.want, got, "SolveRatInAMaze(%v)", tt.maze)
			assert.Equalf(t, tt.exists, gotSolution, "SolveRatInAMaze(%v)", tt.maze)
		})
	}
}
