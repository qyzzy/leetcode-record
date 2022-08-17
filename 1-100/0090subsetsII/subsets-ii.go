package subsetsII

import "sort"

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	var dfs func(p int, t []int)
	dfs = func(p int, t []int) {
		tmp := make([]int, len(t))
		copy(tmp, t)
		res = append(res, tmp)
		if p >= len(nums) {
			return
		}
		for i := p; i < len(nums); i++ {
			if i > p && nums[i] == nums[i-1] {
				continue
			}
			t = append(t, nums[i])
			dfs(i+1, t)
			t = t[:len(t)-1]
		}
	}
	dfs(0, []int{})
	return res
}
