package search_2d_matrix

// Note ints in each row are sorted left to right
// first int of each row is always greater than last int of the previous row
func Search2dMatrix(matrix [][]int, target int) bool {
	if matrix[0][0] > target {
		return false
	}
	rowIndex := len(matrix) - 1
	for rowIndex >= 0 {
		row := matrix[rowIndex]
		// Check the first item in the row, if it is higher than the target, skip this row entirely
		if row[0] <= target {
			for _, cell := range row {
				// At this point if a single item is greater than the target, we know it doesn't exist
				if cell > target {
					return false
				} else if cell == target {
					return true
				}
			}
		}
		rowIndex -= 1
	}
	return false
}
