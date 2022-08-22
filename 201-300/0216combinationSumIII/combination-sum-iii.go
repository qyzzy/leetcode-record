package combinationSumIII

func combinationSum3(k int, n int) [][]int {
	res := [][]int{}
	var dfs func(p int, t []int, sum int)
	dfs = func(p int, t []int, sum int) {
		if sum >= n {
			if sum == n && len(t) == k {
				tmp := make([]int, len(t))
				copy(tmp, t)
				res = append(res, tmp)
			}
			return
		}
		for i := p; i <= 9; i++ {
			t = append(t, i)
			dfs(i+1, t, sum+i)
			t = t[:len(t)-1]
		}
	}
	dfs(1, []int{}, 0)
	return res
}
