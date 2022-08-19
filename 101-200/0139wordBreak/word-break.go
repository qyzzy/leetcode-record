package wordBreak

func wordBreak(s string, wordDict []string) bool {
	mp := map[string]bool{}
	for _, w := range wordDict {
		mp[w] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && mp[s[j:i]] {
				dp[i] = true
			}
		}
	}
	return dp[len(s)]
}
