package backtrack

import "fmt"

type Maze [][]int

func (m Maze) Rows() int {
	return len(m)
}

func (m Maze) Cols() int {
	return len(m[0])
}

func (m Maze) Print() {
	for i := 0; i < m.Rows(); i++ {
		for j := 0; j < m.Cols(); j++ {
			fmt.Printf("%d ", m[i][j])
		}
		fmt.Println()
	}
}

func NewSolutionMaze(maze Maze) Maze {
	sol := make([][]int, maze.Rows())
	for i := range sol {
		sol[i] = make([]int, maze.Cols())
	}
	return sol
}

func NewMaze(config [][]int) Maze {
	maze := make([][]int, len(config))
	for i := range maze {
		maze[i] = make([]int, len(config[0]))
		copy(maze[i], config[i])
	}
	return maze
}
func (m Maze) IsSafe(x, y int) bool {
	return x >= 0 && x < m.Cols() && y >= 0 && y < m.Rows() && m[x][y] == 1
}

func ratInAMazeHelper(maze Maze, solution Maze, x, y int) bool {
	if x == maze.Cols()-1 && y == maze.Rows()-1 {
		// Reached the destination
		solution[x][y] = 1
		return true
	}

	if maze.IsSafe(x, y) {
		solution[x][y] = 1

		// Move right
		if ratInAMazeHelper(maze, solution, x+1, y) {
			return true
		}

		// Move down
		if ratInAMazeHelper(maze, solution, x, y+1) {
			return true
		}

		// If neither right nor down leads to a solution, backtrack
		solution[x][y] = 0
	}

	return false
}

func SolveRatInAMaze(maze Maze) (Maze, bool) {
	solution := NewSolutionMaze(maze)
	solved := ratInAMazeHelper(maze, solution, 0, 0)
	return solution, solved
}
