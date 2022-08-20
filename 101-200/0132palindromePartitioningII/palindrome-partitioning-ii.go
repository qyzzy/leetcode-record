package palindromePartitioningII

func minCut(s string) int {
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}
	for i := 0; i < len(s); i++ {
		for j := 0; j <= i; j++ {
			if i == j {
				dp[j][i] = true
			}
			if i-j <= 2 && s[i] == s[j] {
				dp[j][i] = true
			}
			if i-j > 2 && s[i] == s[j] && dp[j+1][i-1] {
				dp[j][i] = true
			}
		}
	}
	cnt := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		if dp[0][i] {
			continue
		}
		cnt[i] = len(s)
		for j := 0; j < i; j++ {
			if dp[j+1][i] && cnt[j]+1 < cnt[i] {
				cnt[i] = cnt[j] + 1
			}
		}
	}
	//fmt.Println(cnt)
	return cnt[len(s)-1]
}
