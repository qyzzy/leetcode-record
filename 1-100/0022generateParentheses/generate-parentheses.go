package generateParentheses

func generateParenthesis(n int) []string {
	res := []string{}
	var dfs func(l, r int, t string)
	dfs = func(l, r int, t string) {
		if l == 0 && r == 0 {
			res = append(res, t)
			return
		}
		if l > 0 {
			dfs(l-1, r, t+"(")
		}
		if l < r {
			dfs(l, r-1, t+")")
		}
	}
	dfs(n, n, "")
	return res
}
