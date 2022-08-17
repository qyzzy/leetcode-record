package combinationSumII

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	var res [][]int
	sort.Ints(candidates)
	var dfs func(p int, sum int, t []int)
	dfs = func(p int, sum int, t []int) {
		if sum >= target || p == len(candidates) {
			if sum == target {
				tmp := make([]int, len(t))
				copy(tmp, t)
				res = append(res, tmp)
			}
			return
		}
		for i := p; i < len(candidates); i++ {
			if i > p && candidates[i] == candidates[i-1] {
				continue
			}
			t = append(t, candidates[i])
			dfs(i+1, sum+candidates[i], t)
			t = t[:len(t)-1]
		}
	}
	dfs(0, 0, []int{})
	return res
}
