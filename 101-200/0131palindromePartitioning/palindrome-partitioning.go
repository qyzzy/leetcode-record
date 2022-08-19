package palindromePartitioning

func partition(s string) [][]string {
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
	var res [][]string
	var dfs func(p int, t []string)
	dfs = func(p int, t []string) {
		if p >= len(s) {
			tmp := make([]string, len(t))
			copy(tmp, t)
			res = append(res, tmp)
			return
		}
		for i := p; i < len(s); i++ {
			if dp[p][i] {
				t = append(t, s[p:i+1])
				dfs(i+1, t)
				t = t[:len(t)-1]
			}
		}
	}
	dfs(0, []string{})
	return res
}
