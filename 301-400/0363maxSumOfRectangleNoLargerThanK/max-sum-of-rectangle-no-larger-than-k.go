package maxSumOfRectangleNoLargerThanK

import "math"

func maxSumSubmatrix(matrix [][]int, k int) int {
	m, n := len(matrix), len(matrix[0])
	sum := make([][]int, m+1)
	sum[0] = make([]int, n+1)
	for i := 0; i < m; i++ {
		sum[i+1] = make([]int, n+1)
		for j := 0; j < n; j++ {
			sum[i+1][j+1] = sum[i][j+1] + sum[i+1][j] - sum[i][j] + matrix[i][j]
		}
	}
	res := math.MinInt32
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			for ii := 0; ii < i; ii++ {
				for jj := 0; jj < j; jj++ {
					sm := sum[i][j] - sum[ii][j] - sum[i][jj] + sum[ii][jj]
					if sm <= k {
						res = max(res, sm)
					}
				}
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
