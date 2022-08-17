package nQueensII

func totalNQueens(n int) (ans int) {
	columns := make([]bool, n)
	diagonals1 := make([]bool, 2*n-1)
	diagonals2 := make([]bool, 2*n-1)
	var dfs func(int)
	dfs = func(row int) {
		if row == n {
			ans++
			return
		}
		for col, hasQueen := range columns {
			d1, d2 := row+n-1-col, row+col
			if hasQueen || diagonals1[d1] || diagonals2[d2] {
				continue
			}
			columns[col] = true
			diagonals1[d1] = true
			diagonals2[d2] = true
			dfs(row + 1)
			columns[col] = false
			diagonals1[d1] = false
			diagonals2[d2] = false
		}
	}
	dfs(0)
	return
}
