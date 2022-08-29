package battleshipsInABoard

func countBattleships(board [][]byte) int {
	m, n := len(board), len(board[0])
	res := 0
	for i, row := range board {
		for j, ch := range row {
			if ch == 'X' {
				row[j] = '.'
				for k := j + 1; k < n && row[k] == 'X'; k++ {
					row[k] = '.'
				}
				for k := i + 1; k < m && board[k][j] == 'X'; k++ {
					board[k][j] = '.'
				}
				res++
			}
		}
	}
	return res
}
