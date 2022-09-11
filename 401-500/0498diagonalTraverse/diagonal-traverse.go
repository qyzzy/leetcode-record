package diagonalTraverse

func findDiagonalOrder(mat [][]int) []int {
	var res []int
	for m, n, k := len(mat), len(mat[0]), 0; k < m+n-1; k++ {
		if k&1 == 1 {
			for x := max(0, k-n+1); x < min(k+1, m); x++ {
				res = append(res, mat[x][k-x])
			}
		} else {
			for x := min(k, m-1); x >= max(0, k-n+1); x-- {
				res = append(res, mat[x][k-x])
			}
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
