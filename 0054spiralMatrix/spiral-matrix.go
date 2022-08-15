package spiralMatrix

func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	res := make([]int, m*n)
	dir := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	i, j, k := 0, 0, 0
	ind := 0
	t := make([][]int, m)
	for i := 0; i < m; i++ {
		t[i] = make([]int, n)
	}
	for k < len(res) {
		res[k] = matrix[i][j]
		t[i][j] = 1
		if i+dir[ind][0] >= m || j+dir[ind][1] >= n || i+dir[ind][0] < 0 || j+dir[ind][1] < 0 || t[i+dir[ind][0]][j+dir[ind][1]] == 1 {
			ind = (ind + 1) % 4
		}
		i += dir[ind][0]
		j += dir[ind][1]
		k++
	}
	return res
}
