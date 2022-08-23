package uglyNumberII

func nthUglyNumber(n int) int {
	dp := make([]int, n)
	a, b, c := 0, 0, 0
	dp[0] = 1
	for i := 1; i < n; i++ {
		c1, c2, c3 := dp[a]*2, dp[b]*3, dp[c]*5
		dp[i] = min(c1, min(c2, c3))
		if dp[i] == c1 {
			a++
		}
		if dp[i] == c2 {
			b++
		}
		if dp[i] == c3 {
			c++
		}
	}
	return dp[n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
