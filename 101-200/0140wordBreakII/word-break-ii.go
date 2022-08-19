package wordBreakII

func wordBreak(s string, wordDict []string) []string {
	mp := map[string]bool{}
	for _, w := range wordDict {
		mp[w] = true
	}
	var res []string
	var dfs func(p int, t string)
	dfs = func(p int, t string) {
		if p >= len(s) {
			res = append(res, t)
			return
		}
		for i := p; i < len(s); i++ {
			if mp[s[p:i+1]] {
				if p == 0 {
					dfs(i+1, t+s[p:i+1])
				} else {
					t += " "
					dfs(i+1, t+s[p:i+1])
					t = t[:len(t)-1]
				}
			}
		}
	}
	dfs(0, "")
	return res
}
