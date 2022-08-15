package spiralMatrixII

func generateMatrix(n int) [][]int {
	dir := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	i, j, k := 0, 0, 1
	ind := 0
	for k <= n*n {
		res[i][j] = k
		if i+dir[ind][0] >= n || i+dir[ind][0] < 0 || j+dir[ind][1] >= n || j+dir[ind][1] < 0 || res[i+dir[ind][0]][j+dir[ind][1]] != 0 {
			ind = (ind + 1) % 4
		}
		i += dir[ind][0]
		j += dir[ind][1]
		k++
	}
	return res
}
