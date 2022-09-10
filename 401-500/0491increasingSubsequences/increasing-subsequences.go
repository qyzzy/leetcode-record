package increasingSubsequences

func findSubsequences(nums []int) [][]int {
	res := [][]int{}
	var dfs func(i, last int, t []int)
	dfs = func(i, last int, t []int) {
		if i >= len(nums) {
			if len(t) >= 2 {
				tmp := make([]int, len(t))
				copy(tmp, t)
				res = append(res, tmp)
			}
			return
		}
		if nums[i] >= last {
			t = append(t, nums[i])
			dfs(i+1, nums[i], t)
			t = t[:len(t)-1]
		}
		if nums[i] != last {
			dfs(i+1, last, t)
		}
	}
	dfs(0, -1000, []int{})
	return res
}
