package permutations

func permute(nums []int) [][]int {
	var res [][]int
	m := make([]int, len(nums))
	var dfs func(t []int, m []int)
	dfs = func(t []int, m []int) {
		if len(t) == len(nums) {
			tmp := make([]int, len(nums))
			copy(tmp, t)
			res = append(res, tmp)
			return
		}
		for i := 0; i < len(nums); i++ {
			if m[i] == 1 {
				continue
			}
			m[i] = 1
			t = append(t, nums[i])
			dfs(t, m)
			m[i] = 0
			t = t[:len(t)-1]
		}
	}
	dfs([]int{}, m)
	return res
}
