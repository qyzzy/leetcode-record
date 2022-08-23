package perfectSquares

func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = i
		for j := 1; product(j) <= i; j++ {
			dp[i] = min(dp[i], dp[i-product(j)]+1)
		}
	}
	return dp[n]
}

func product(a int) int {
	return a * a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
