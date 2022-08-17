package wordSearch

func exist(board [][]byte, word string) bool {
	n, m := len(board), len(board[0])
	var dfs func(i, j, k int) bool
	dfs = func(i, j, k int) bool {
		if k == len(word) {
			return true
		}
		if i < 0 || j < 0 || i >= n || j >= m || board[i][j] != word[k] {
			return false
		}
		board[i][j] = '0'
		res := dfs(i+1, j, k+1) || dfs(i-1, j, k+1) || dfs(i, j+1, k+1) || dfs(i, j-1, k+1)
		board[i][j] = word[k]
		return res
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}
