package subsets

func subsets(nums []int) [][]int {
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
			t = append(t, nums[i])
			dfs(i+1, t)
			t = t[:len(t)-1]
		}
	}
	dfs(0, []int{})
	return res
}
