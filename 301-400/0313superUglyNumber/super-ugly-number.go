package superUglyNumber

import "math"

func nthSuperUglyNumber(n int, primes []int) int {
	dp := make([]int, n)
	ct := make([]int, len(primes))
	dp[0] = 1
	for i := 1; i < n; i++ {
		tmp := make([]int, len(primes))
		for j := 0; j < len(primes); j++ {
			tmp[j] = dp[ct[j]] * primes[j]
		}
		dp[i] = math.MaxInt32
		for j := 0; j < len(primes); j++ {
			dp[i] = min(dp[i], tmp[j])
		}
		for j := 0; j < len(primes); j++ {
			if dp[i] == tmp[j] {
				ct[j]++
			}
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
