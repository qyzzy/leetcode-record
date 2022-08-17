package combinations

func combine(n int, k int) [][]int {
	var res [][]int
	var dfs func(p int, t []int)
	dfs = func(p int, t []int) {
		if len(t) == k {
			tmp := make([]int, len(t))
			copy(tmp, t)
			res = append(res, tmp)
			return
		}
		for i := p; i <= n; i++ {
			t = append(t, i)
			dfs(i+1, t)
			t = t[:len(t)-1]
		}
	}
	dfs(1, []int{})
	return res
}
