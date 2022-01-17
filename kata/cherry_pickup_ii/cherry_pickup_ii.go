package cherry_pickup_ii

func InitMoves(grid[][]int) [][][]int {
	rows := len(grid)
	columns := len(grid[0])
	moves := make([][][]int, rows)
	// both robots share a row at a time
	// but each can be in a different column, so triple loop to instantiate the table
	for i := 0; i < rows; i++ {
		moves[i] = make([][]int, columns)
		for j := 0; j < columns; j++ {
			moves[i][j] = make([]int, columns)
			for k := 0; k < columns; k++ {
				moves[i][j][k] = -1
			}
		}
	}
	return moves
}

// Things to store in state:
// The result of each path from one column to another
func CalculateCherries(row int, col1 int, col2 int, grid[][]int, moves [][][]int) int {
	// Avoid going off he edges
	if col1 < 0 || col1 >= len(grid[0]) || col2 < 0 || col2 >= len(grid[0]){
			return 0
	}
	// First check to see if we've computed this state already
	if moves[row][col1][col2] != -1 {
		return moves[row][col1][col2]
	}

	// Track the total of the path
	result := grid[row][col1]
	// If both robots are not in the same cell, add robot 2 as well
	// If they are both in the same cell, then only count that number of cherries once
	if (col1 != col2) {
		result += grid[row][col2]
	}
	// if we are on the last row, cache the total and return
	if row != len(grid) - 1 {
		max := 0
		// Next calculate each possible transition and get the max
		for newCol1 := col1 - 1; newCol1 <= col1 + 1; newCol1 ++ {
			for newCol2 := col2 - 1; newCol2 <= col2 + 1; newCol2 ++ {
				transitionResult := CalculateCherries(row + 1, newCol1, newCol2, grid, moves)
				if transitionResult > max {
					max = transitionResult
				}
			}
		}

		result += max
	}

	moves[row][col1][col2] = result
	return result
}

func CherryPickup(grid [][]int) int {
  return CalculateCherries(0, 0, len(grid[0]) - 1, grid, InitMoves(grid))
}
