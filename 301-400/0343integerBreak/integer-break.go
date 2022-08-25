package integerBreak

func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	dp[2] = 1
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1]
		for j := 2; j < i; j++ {
			dp[i] = max(max(dp[i], j*(i-j)), dp[j]*(i-j))
		}
	}
	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
