package combinationSum

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var res [][]int
	var dfs func(p int, t []int, sum int)
	dfs = func(p int, t []int, sum int) {
		if p >= len(candidates) || sum >= target {
			if sum == target {
				tmp := make([]int, len(t))
				copy(tmp, t)
				res = append(res, tmp)
			}
			return
		}
		for i := p; i < len(candidates); i++ {
			t = append(t, candidates[i])
			dfs(i, t, sum+candidates[i])
			t = t[:len(t)-1]
		}
	}
	dfs(0, []int{}, 0)
	return res
}
