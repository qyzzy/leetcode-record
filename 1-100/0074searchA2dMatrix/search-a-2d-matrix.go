package searchA2dMatrix

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])
	for i := 0; i < m; i++ {
		if matrix[i][0] > target {
			return false
		}
		if matrix[i][n-1] >= target {
			for j := 0; j < n; j++ {
				if target == matrix[i][j] {
					return true
				}
			}
		}
		if matrix[i][n-1] < target {
			continue
		}
	}
	return false
}
